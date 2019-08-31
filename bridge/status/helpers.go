package status

import (
	"encoding/hex"
	"crypto/ecdsa"

	crypto "github.com/ethereum/go-ethereum/crypto"
)

func publicKeyToHex(pubkey *ecdsa.PublicKey) string {
	return "0x" + hex.EncodeToString(crypto.FromECDSAPub(pubkey))
}

// isPubKeyEqual checks that two public keys are equal
func isPubKeyEqual(a, b *ecdsa.PublicKey) bool {
	// the curve is always the same, just compare the points
	return a.X.Cmp(b.X) == 0 && a.Y.Cmp(b.Y) == 0
}
