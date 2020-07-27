package helper

import (
	"fmt"
	"math/big"
	"time"

	"github.com/maticnetwork/bor/common"
	hmTypes "github.com/maticnetwork/heimdall/types"
)

// GetERC20FromFaucet provides matic tokens for testing purpose
func GetERC20FromFaucet(amount uint64, receiverAddr string, contractAddr common.Address) error {
	fmt.Println("Transfering tokens")
	// get matic Erc20 instance
	erc20TokenInstance, err := c.ContractCaller.GetERC20TokenInstance(contractAddr, true)
	if err != nil {
		return err
	}

	// send tx to transfer tokens from faucet account to user address
	bigAmount := big.NewInt(0).SetUint64(amount)
	receiver := hmTypes.HexToHeimdallAddress(receiverAddr).EthAddress()

	c.ContractCaller.TransferToken(bigAmount, receiver, contractAddr, erc20TokenInstance, c.MainchainFaucet, true)
	time.Sleep(5 * time.Second)
	return nil
}

// GetTestEthMainchain provides eth for testing purpose
func GetTestEthMainchain(amount uint64, address string) error {
	// send tx to transfer eth from faucet account to user address
	bigAmount := big.NewInt(0).SetUint64(amount)
	receiver := hmTypes.HexToHeimdallAddress(address).EthAddress()

	c.ContractCaller.TransferEth(bigAmount, receiver, c.MainchainFaucet, c.ContractCaller.MainChainClient)
	time.Sleep(5 * time.Second)
	return nil
}

// GetTestEthMatic provides eth for testing purpose
func GetTestEthMatic(amount uint64, address string) error {
	// send tx to transfer eth from faucet account to user address
	bigAmount := big.NewInt(0).SetUint64(amount)
	receiver := hmTypes.HexToHeimdallAddress(address).EthAddress()

	c.ContractCaller.TransferEth(bigAmount, receiver, c.MaticFaucet, c.ContractCaller.MaticChainClient)
	time.Sleep(5 * time.Second)
	return nil
}
