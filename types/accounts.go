package types

import (
	"crypto/ecdsa"
	"log"

	"github.com/maticnetwork/bor/common/hexutil"
	"github.com/maticnetwork/bor/crypto"
)

// Account used tests
type Account struct {
	Address string `json:"address"`
	PubKey  string `json:"pub_key"`
	PrivKey string `json:"priv_key"`
}

// RandomAccounts generates n random accounts
func RandomAccounts(n int) []Account {
	accs := make([]Account, n)
	for i := 0; i < n; i++ {
		// private key
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}
		privateKeyBytes := crypto.FromECDSA(privateKey)
		accs[i].PrivKey = hexutil.Encode(privateKeyBytes)

		// public key
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

		if !ok {
			log.Fatal("error casting public key to ECDSA")
		}
		publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
		accs[i].PubKey = hexutil.Encode(publicKeyBytes)

		// address
		accs[i].Address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	}

	return accs
}
