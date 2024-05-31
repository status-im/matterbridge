package transfer

import (
	"context"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	ethTypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	"github.com/status-im/status-go/account"
	"github.com/status-im/status-go/eth-node/crypto"
	"github.com/status-im/status-go/eth-node/types"
	"github.com/status-im/status-go/services/wallet/bridge"
	wallet_common "github.com/status-im/status-go/services/wallet/common"
	"github.com/status-im/status-go/signal"
)

const multiTransactionColumns = "id, from_network_id, from_tx_hash, from_address, from_asset, from_amount, to_network_id, to_tx_hash, to_address, to_asset, to_amount, type, cross_tx_id, timestamp"
const selectMultiTransactionColumns = "id, COALESCE(from_network_id, 0), from_tx_hash, from_address, from_asset, from_amount, COALESCE(to_network_id, 0), to_tx_hash, to_address, to_asset, to_amount, type, cross_tx_id, timestamp"

func rowsToMultiTransactions(rows *sql.Rows) ([]*MultiTransaction, error) {
	var multiTransactions []*MultiTransaction
	for rows.Next() {
		multiTransaction := &MultiTransaction{}
		var fromAmountDB, toAmountDB sql.NullString
		var fromTxHash, toTxHash sql.RawBytes
		err := rows.Scan(
			&multiTransaction.ID,
			&multiTransaction.FromNetworkID,
			&fromTxHash,
			&multiTransaction.FromAddress,
			&multiTransaction.FromAsset,
			&fromAmountDB,
			&multiTransaction.ToNetworkID,
			&toTxHash,
			&multiTransaction.ToAddress,
			&multiTransaction.ToAsset,
			&toAmountDB,
			&multiTransaction.Type,
			&multiTransaction.CrossTxID,
			&multiTransaction.Timestamp,
		)
		if len(fromTxHash) > 0 {
			multiTransaction.FromTxHash = common.BytesToHash(fromTxHash)
		}
		if len(toTxHash) > 0 {
			multiTransaction.ToTxHash = common.BytesToHash(toTxHash)
		}
		if err != nil {
			return nil, err
		}

		if fromAmountDB.Valid {
			multiTransaction.FromAmount = new(hexutil.Big)
			if _, ok := (*big.Int)(multiTransaction.FromAmount).SetString(fromAmountDB.String, 0); !ok {
				return nil, errors.New("failed to convert fromAmountDB.String to big.Int: " + fromAmountDB.String)
			}
		}

		if toAmountDB.Valid {
			multiTransaction.ToAmount = new(hexutil.Big)
			if _, ok := (*big.Int)(multiTransaction.ToAmount).SetString(toAmountDB.String, 0); !ok {
				return nil, errors.New("failed to convert fromAmountDB.String to big.Int: " + toAmountDB.String)
			}
		}

		multiTransactions = append(multiTransactions, multiTransaction)
	}

	return multiTransactions, nil
}

// insertMultiTransaction inserts a multi transaction into the database and updates timestamp
func insertMultiTransaction(db *sql.DB, multiTransaction *MultiTransaction) error {
	insert, err := db.Prepare(fmt.Sprintf(`INSERT INTO multi_transactions (%s)
											VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, multiTransactionColumns))
	if err != nil {
		return err
	}
	_, err = insert.Exec(
		multiTransaction.ID,
		multiTransaction.FromNetworkID,
		multiTransaction.FromTxHash,
		multiTransaction.FromAddress,
		multiTransaction.FromAsset,
		multiTransaction.FromAmount.String(),
		multiTransaction.ToNetworkID,
		multiTransaction.ToTxHash,
		multiTransaction.ToAddress,
		multiTransaction.ToAsset,
		multiTransaction.ToAmount.String(),
		multiTransaction.Type,
		multiTransaction.CrossTxID,
		multiTransaction.Timestamp,
	)
	if err != nil {
		return err
	}
	defer insert.Close()

	return err
}

func (tm *TransactionManager) InsertMultiTransaction(multiTransaction *MultiTransaction) (wallet_common.MultiTransactionIDType, error) {
	return multiTransaction.ID, insertMultiTransaction(tm.db, multiTransaction)
}

func updateMultiTransaction(db *sql.DB, multiTransaction *MultiTransaction) error {
	if multiTransaction.ID == wallet_common.NoMultiTransactionID {
		return fmt.Errorf("no multitransaction ID")
	}

	update, err := db.Prepare(fmt.Sprintf(`REPLACE INTO multi_transactions (%s)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, multiTransactionColumns))

	if err != nil {
		return err
	}
	_, err = update.Exec(
		multiTransaction.ID,
		multiTransaction.FromNetworkID,
		multiTransaction.FromTxHash,
		multiTransaction.FromAddress,
		multiTransaction.FromAsset,
		multiTransaction.FromAmount.String(),
		multiTransaction.ToNetworkID,
		multiTransaction.ToTxHash,
		multiTransaction.ToAddress,
		multiTransaction.ToAsset,
		multiTransaction.ToAmount.String(),
		multiTransaction.Type,
		multiTransaction.CrossTxID,
		multiTransaction.Timestamp,
	)
	if err != nil {
		return err
	}
	return update.Close()
}

func (tm *TransactionManager) UpdateMultiTransaction(multiTransaction *MultiTransaction) error {
	return updateMultiTransaction(tm.db, multiTransaction)
}

func (tm *TransactionManager) CreateMultiTransactionFromCommand(ctx context.Context, command *MultiTransactionCommand,
	data []*bridge.TransactionBridge) (*MultiTransaction, error) {

	multiTransaction := multiTransactionFromCommand(command)

	if multiTransaction.Type == MultiTransactionSend && multiTransaction.FromNetworkID == 0 && len(data) == 1 {
		multiTransaction.FromNetworkID = data[0].ChainID
	}

	return multiTransaction, nil
}

func (tm *TransactionManager) SendTransactionForSigningToKeycard(ctx context.Context, multiTransaction *MultiTransaction, data []*bridge.TransactionBridge, bridges map[string]bridge.Bridge) error {
	acc, err := tm.accountsDB.GetAccountByAddress(types.Address(multiTransaction.FromAddress))
	if err != nil {
		return err
	}

	kp, err := tm.accountsDB.GetKeypairByKeyUID(acc.KeyUID)
	if err != nil {
		return err
	}

	if !kp.MigratedToKeycard() {
		return fmt.Errorf("account being used is not migrated to a keycard, password is required")
	}

	tm.multiTransactionForKeycardSigning = multiTransaction
	tm.transactionsBridgeData = data
	hashes, err := tm.buildTransactions(bridges)
	if err != nil {
		return err
	}

	signal.SendTransactionsForSigningEvent(hashes)

	return nil
}

func (tm *TransactionManager) SendTransactions(ctx context.Context, multiTransaction *MultiTransaction, data []*bridge.TransactionBridge, bridges map[string]bridge.Bridge, account *account.SelectedExtKey) (*MultiTransactionCommandResult, error) {
	hashes, err := tm.sendTransactions(multiTransaction, data, bridges, account)
	if err != nil {
		return nil, err
	}

	return &MultiTransactionCommandResult{
		ID:     int64(multiTransaction.ID),
		Hashes: hashes,
	}, nil
}

func (tm *TransactionManager) ProceedWithTransactionsSignatures(ctx context.Context, signatures map[string]SignatureDetails) (*MultiTransactionCommandResult, error) {
	if tm.multiTransactionForKeycardSigning == nil {
		return nil, errors.New("no multi transaction to proceed with")
	}
	if len(tm.transactionsBridgeData) == 0 {
		return nil, errors.New("no transactions bridge data to proceed with")
	}
	if len(tm.transactionsForKeycardSingning) == 0 {
		return nil, errors.New("no transactions to proceed with")
	}
	if len(signatures) != len(tm.transactionsForKeycardSingning) {
		return nil, errors.New("not all transactions have been signed")
	}

	// check if all transactions have been signed
	for hash, desc := range tm.transactionsForKeycardSingning {
		sigDetails, ok := signatures[hash.String()]
		if !ok {
			return nil, fmt.Errorf("missing signature for transaction %s", hash)
		}

		rBytes, _ := hex.DecodeString(sigDetails.R)
		sBytes, _ := hex.DecodeString(sigDetails.S)
		vByte := byte(0)
		if sigDetails.V == "01" {
			vByte = 1
		}

		desc.signature = make([]byte, crypto.SignatureLength)
		copy(desc.signature[32-len(rBytes):32], rBytes)
		copy(desc.signature[64-len(rBytes):64], sBytes)
		desc.signature[64] = vByte
	}

	// send transactions
	hashes := make(map[uint64][]types.Hash)
	for _, desc := range tm.transactionsForKeycardSingning {
		hash, err := tm.transactor.AddSignatureToTransactionAndSend(
			desc.chainID,
			desc.from,
			tm.multiTransactionForKeycardSigning.FromAsset,
			tm.multiTransactionForKeycardSigning.ID,
			desc.builtTx,
			desc.signature,
		)
		if err != nil {
			return nil, err
		}
		hashes[desc.chainID] = append(hashes[desc.chainID], hash)
	}

	_, err := tm.InsertMultiTransaction(tm.multiTransactionForKeycardSigning)
	if err != nil {
		log.Error("failed to insert multi transaction", "err", err)
	}

	return &MultiTransactionCommandResult{
		ID:     int64(tm.multiTransactionForKeycardSigning.ID),
		Hashes: hashes,
	}, nil
}

func multiTransactionFromCommand(command *MultiTransactionCommand) *MultiTransaction {
	multiTransaction := NewMultiTransaction(
		/* Timestamp:     */ uint64(time.Now().Unix()),
		/* FromNetworkID: */ 0,
		/* ToNetworkID:	  */ 0,
		/* FromTxHash:    */ common.Hash{},
		/* ToTxHash:      */ common.Hash{},
		/* FromAddress:   */ command.FromAddress,
		/* ToAddress:     */ command.ToAddress,
		/* FromAsset:     */ command.FromAsset,
		/* ToAsset:       */ command.ToAsset,
		/* FromAmount:    */ command.FromAmount,
		/* ToAmount:      */ new(hexutil.Big),
		/* Type:		  */ command.Type,
		/* CrossTxID:	  */ "",
	)

	return multiTransaction
}

func (tm *TransactionManager) buildTransactions(bridges map[string]bridge.Bridge) ([]string, error) {
	tm.transactionsForKeycardSingning = make(map[common.Hash]*TransactionDescription)
	var hashes []string
	for _, bridgeTx := range tm.transactionsBridgeData {
		builtTx, err := bridges[bridgeTx.BridgeName].BuildTransaction(bridgeTx)
		if err != nil {
			return hashes, err
		}

		signer := ethTypes.NewLondonSigner(big.NewInt(int64(bridgeTx.ChainID)))
		txHash := signer.Hash(builtTx)

		tm.transactionsForKeycardSingning[txHash] = &TransactionDescription{
			from:    common.Address(bridgeTx.From()),
			chainID: bridgeTx.ChainID,
			builtTx: builtTx,
		}

		hashes = append(hashes, txHash.String())
	}

	return hashes, nil
}

func (tm *TransactionManager) sendTransactions(multiTransaction *MultiTransaction,
	data []*bridge.TransactionBridge, bridges map[string]bridge.Bridge, account *account.SelectedExtKey) (
	map[uint64][]types.Hash, error) {

	hashes := make(map[uint64][]types.Hash)
	for _, tx := range data {
		if tx.TransferTx != nil {
			tx.TransferTx.MultiTransactionID = multiTransaction.ID
			tx.TransferTx.Symbol = multiTransaction.FromAsset
		}
		if tx.HopTx != nil {
			tx.HopTx.MultiTransactionID = multiTransaction.ID
			tx.HopTx.Symbol = multiTransaction.FromAsset
		}
		if tx.CbridgeTx != nil {
			tx.CbridgeTx.MultiTransactionID = multiTransaction.ID
			tx.CbridgeTx.Symbol = multiTransaction.FromAsset
		}
		if tx.ERC721TransferTx != nil {
			tx.ERC721TransferTx.MultiTransactionID = multiTransaction.ID
			tx.ERC721TransferTx.Symbol = multiTransaction.FromAsset
		}
		if tx.ERC1155TransferTx != nil {
			tx.ERC1155TransferTx.MultiTransactionID = multiTransaction.ID
			tx.ERC1155TransferTx.Symbol = multiTransaction.FromAsset
		}
		if tx.SwapTx != nil {
			tx.SwapTx.MultiTransactionID = multiTransaction.ID
			tx.SwapTx.Symbol = multiTransaction.FromAsset
		}

		hash, err := bridges[tx.BridgeName].Send(tx, account)
		if err != nil {
			return nil, err // TODO: One of transfers within transaction could have been sent. Need to notify user about it
		}
		hashes[tx.ChainID] = append(hashes[tx.ChainID], hash)
	}
	return hashes, nil
}

func (tm *TransactionManager) GetMultiTransactions(ctx context.Context, ids []wallet_common.MultiTransactionIDType) ([]*MultiTransaction, error) {
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, v := range ids {
		placeholders[i] = "?"
		args[i] = v
	}

	stmt, err := tm.db.Prepare(fmt.Sprintf(`SELECT %s
											FROM multi_transactions
											WHERE id in (%s)`,
		selectMultiTransactionColumns,
		strings.Join(placeholders, ",")))
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return rowsToMultiTransactions(rows)
}

func (tm *TransactionManager) getBridgeMultiTransactions(ctx context.Context, toChainID uint64, crossTxID string) ([]*MultiTransaction, error) {
	stmt, err := tm.db.Prepare(fmt.Sprintf(`SELECT %s
											FROM multi_transactions
											WHERE type=? AND to_network_id=? AND cross_tx_id=?`,
		multiTransactionColumns))
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(MultiTransactionBridge, toChainID, crossTxID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return rowsToMultiTransactions(rows)
}

func (tm *TransactionManager) GetBridgeOriginMultiTransaction(ctx context.Context, toChainID uint64, crossTxID string) (*MultiTransaction, error) {
	multiTxs, err := tm.getBridgeMultiTransactions(ctx, toChainID, crossTxID)
	if err != nil {
		return nil, err
	}

	for _, multiTx := range multiTxs {
		// Origin MultiTxs will have a missing "ToTxHash"
		if multiTx.ToTxHash == emptyHash {
			return multiTx, nil
		}
	}

	return nil, nil
}

func (tm *TransactionManager) GetBridgeDestinationMultiTransaction(ctx context.Context, toChainID uint64, crossTxID string) (*MultiTransaction, error) {
	multiTxs, err := tm.getBridgeMultiTransactions(ctx, toChainID, crossTxID)
	if err != nil {
		return nil, err
	}

	for _, multiTx := range multiTxs {
		// Destination MultiTxs will have a missing "FromTxHash"
		if multiTx.FromTxHash == emptyHash {
			return multiTx, nil
		}
	}

	return nil, nil
}

func idFromTimestamp() wallet_common.MultiTransactionIDType {
	return wallet_common.MultiTransactionIDType(time.Now().UnixMilli())
}

var multiTransactionIDGenerator func() wallet_common.MultiTransactionIDType = idFromTimestamp

func (tm *TransactionManager) removeMultiTransactionByAddress(address common.Address) error {
	// We must not remove those transactions, where from_address and to_address are different and both are stored in accounts DB
	// and one of them is equal to the address, as we want to keep the records for the other address
	// That is why we don't use cascade delete here with references to transfers table, as we might have 2 records in multi_transactions
	// for the same transaction, one for each address

	stmt, err := tm.db.Prepare(`SELECT id, from_address, to_address
								FROM multi_transactions
								WHERE from_address=? OR to_address=?`)
	if err != nil {
		return err
	}

	rows, err := stmt.Query(address, address)
	if err != nil {
		return err
	}
	defer rows.Close()

	ids := make([]int, 0)
	id, fromAddress, toAddress := 0, common.Address{}, common.Address{}
	for rows.Next() {
		err = rows.Scan(&id, &fromAddress, &toAddress)
		if err != nil {
			log.Error("Failed to scan row", "error", err)
			continue
		}

		// Remove self transactions as well, leave only those where we have the counterparty in accounts DB
		if fromAddress != toAddress {
			// If both addresses are stored in accounts DB, we don't remove the record
			var addressToCheck common.Address
			if fromAddress == address {
				addressToCheck = toAddress
			} else {
				addressToCheck = fromAddress
			}
			counterpartyExists, err := tm.accountsDB.AddressExists(types.Address(addressToCheck))
			if err != nil {
				log.Error("Failed to query accounts db for a given address", "address", address, "error", err)
				continue
			}

			// Skip removal if counterparty is in accounts DB and removed address is not sender
			if counterpartyExists && address != fromAddress {
				continue
			}
		}

		ids = append(ids, id)
	}

	if len(ids) > 0 {
		for _, id := range ids {
			_, err = tm.db.Exec(`DELETE FROM multi_transactions WHERE id=?`, id)
			if err != nil {
				log.Error("Failed to remove multitransaction", "id", id, "error", err)
			}
		}
	}

	return err
}
