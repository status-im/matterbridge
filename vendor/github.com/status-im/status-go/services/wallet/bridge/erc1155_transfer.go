package bridge

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/status-im/status-go/account"
	"github.com/status-im/status-go/contracts/ierc1155"
	"github.com/status-im/status-go/eth-node/types"
	"github.com/status-im/status-go/params"
	"github.com/status-im/status-go/rpc"
	"github.com/status-im/status-go/services/wallet/token"
	"github.com/status-im/status-go/transactions"
)

type ERC1155TransferTxArgs struct {
	transactions.SendTxArgs
	TokenID   *hexutil.Big   `json:"tokenId"`
	Recipient common.Address `json:"recipient"`
	Amount    *hexutil.Big   `json:"amount"`
}

type ERC1155TransferBridge struct {
	rpcClient  *rpc.Client
	transactor *transactions.Transactor
}

func NewERC1155TransferBridge(rpcClient *rpc.Client, transactor *transactions.Transactor) *ERC1155TransferBridge {
	return &ERC1155TransferBridge{rpcClient: rpcClient, transactor: transactor}
}

func (s *ERC1155TransferBridge) Name() string {
	return "ERC1155Transfer"
}

func (s *ERC1155TransferBridge) AvailableFor(from, to *params.Network, token *token.Token, toToken *token.Token) (bool, error) {
	return from.ChainID == to.ChainID && toToken == nil, nil
}

func (s *ERC1155TransferBridge) CalculateFees(from, to *params.Network, token *token.Token, amountIn *big.Int) (*big.Int, *big.Int, error) {
	return big.NewInt(0), big.NewInt(0), nil
}

func (s *ERC1155TransferBridge) EstimateGas(fromNetwork *params.Network, toNetwork *params.Network, from common.Address, to common.Address, token *token.Token, toToken *token.Token, amountIn *big.Int) (uint64, error) {
	ethClient, err := s.rpcClient.EthClient(fromNetwork.ChainID)
	if err != nil {
		return 0, err
	}

	var input []byte
	value := new(big.Int)

	abi, err := abi.JSON(strings.NewReader(ierc1155.Ierc1155ABI))
	if err != nil {
		return 0, err
	}
	id, success := big.NewInt(0).SetString(token.Symbol, 0)
	if !success {
		return 0, fmt.Errorf("failed to convert %s to big.Int", token.Symbol)
	}
	input, err = abi.Pack("safeTransferFrom",
		from,
		to,
		id,
		amountIn,
		[]byte{},
	)

	if err != nil {
		return 0, err
	}

	ctx := context.Background()

	if code, err := ethClient.PendingCodeAt(ctx, token.Address); err != nil {
		return 0, err
	} else if len(code) == 0 {
		return 0, bind.ErrNoCode
	}

	msg := ethereum.CallMsg{
		From:  from,
		To:    &token.Address,
		Value: value,
		Data:  input,
	}

	estimation, err := ethClient.EstimateGas(ctx, msg)
	if err != nil {
		return 0, err
	}
	increasedEstimation := float64(estimation) * IncreaseEstimatedGasFactor
	return uint64(increasedEstimation), nil
}

func (s *ERC1155TransferBridge) BuildTx(network, _ *params.Network, fromAddress common.Address, toAddress common.Address, token *token.Token, amountIn *big.Int, _ *big.Int) (*ethTypes.Transaction, error) {
	contractAddress := types.Address(token.Address)

	// We store ERC1155 Token ID using big.Int.String() in token.Symbol
	tokenID, success := new(big.Int).SetString(token.Symbol, 10)
	if !success {
		return nil, fmt.Errorf("failed to convert ERC1155's Symbol %s to big.Int", token.Symbol)
	}

	sendArgs := &TransactionBridge{
		ERC1155TransferTx: &ERC1155TransferTxArgs{
			SendTxArgs: transactions.SendTxArgs{
				From:  types.Address(fromAddress),
				To:    &contractAddress,
				Value: (*hexutil.Big)(amountIn),
				Data:  types.HexBytes("0x0"),
			},
			TokenID:   (*hexutil.Big)(tokenID),
			Recipient: toAddress,
			Amount:    (*hexutil.Big)(amountIn),
		},
		ChainID: network.ChainID,
	}

	return s.BuildTransaction(sendArgs)
}

func (s *ERC1155TransferBridge) sendOrBuild(sendArgs *TransactionBridge, signerFn bind.SignerFn) (tx *ethTypes.Transaction, err error) {
	ethClient, err := s.rpcClient.EthClient(sendArgs.ChainID)
	if err != nil {
		return tx, err
	}

	contract, err := ierc1155.NewIerc1155(common.Address(*sendArgs.ERC1155TransferTx.To), ethClient)
	if err != nil {
		return tx, err
	}

	nonce, err := s.transactor.NextNonce(s.rpcClient, sendArgs.ChainID, sendArgs.ERC1155TransferTx.From)
	if err != nil {
		return tx, err
	}

	argNonce := hexutil.Uint64(nonce)
	sendArgs.ERC1155TransferTx.Nonce = &argNonce
	txOpts := sendArgs.ERC1155TransferTx.ToTransactOpts(signerFn)
	tx, err = contract.SafeTransferFrom(
		txOpts, common.Address(sendArgs.ERC1155TransferTx.From),
		sendArgs.ERC1155TransferTx.Recipient,
		sendArgs.ERC1155TransferTx.TokenID.ToInt(),
		sendArgs.ERC1155TransferTx.Amount.ToInt(),
		[]byte{},
	)
	return tx, err
}

func (s *ERC1155TransferBridge) Send(sendArgs *TransactionBridge, verifiedAccount *account.SelectedExtKey) (hash types.Hash, err error) {
	tx, err := s.sendOrBuild(sendArgs, getSigner(sendArgs.ChainID, sendArgs.ERC1155TransferTx.From, verifiedAccount))
	if err != nil {
		return hash, err
	}
	return types.Hash(tx.Hash()), nil
}

func (s *ERC1155TransferBridge) BuildTransaction(sendArgs *TransactionBridge) (*ethTypes.Transaction, error) {
	return s.sendOrBuild(sendArgs, nil)
}

func (s *ERC1155TransferBridge) CalculateAmountOut(from, to *params.Network, amountIn *big.Int, symbol string) (*big.Int, error) {
	return amountIn, nil
}

func (s *ERC1155TransferBridge) GetContractAddress(network *params.Network, token *token.Token) *common.Address {
	return &token.Address
}
