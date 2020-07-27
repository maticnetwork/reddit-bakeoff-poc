package helper

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/maticnetwork/bor/common/hexutil"
	"github.com/maticnetwork/bor/crypto"
	"github.com/maticnetwork/reddit-bakeoff-poc/types"
)

func GenerateAccount() types.Account {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)
	account := types.Account{
		PrivKey: hexutil.Encode(privateKeyBytes)[2:],
		PubKey:  hexutil.Encode(publicKeyBytes),
		Address: address,
	}
	return account
}

func GetAmountFromPower(amount *big.Int) (*big.Int, error) {
	decimals18 := big.NewInt(10).Exp(big.NewInt(10), big.NewInt(18), nil)
	return amount.Mul(amount, decimals18), nil
}
