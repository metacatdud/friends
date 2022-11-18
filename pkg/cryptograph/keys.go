package cryptograph

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
)

func GenPrivateKey() string {
	pk, _ := btcec.NewPrivateKey()
	return pk.Key.String()
}

func GetPublicKey(pk string) (string, error) {
	b, err := hex.DecodeString(pk)
	if err != nil {
		return "", err
	}

	_, pubK := btcec.PrivKeyFromBytes(b)
	return hex.EncodeToString(schnorr.SerializePubKey(pubK)), nil
}
