package helper

import (
	"github.com/maticnetwork/reddit-bakeoff-poc/types"
)

// Conf config for selected network
type Conf struct {
	MainchainFaucet    types.Account
	MaticFaucet        types.Account
	GenesisValAccounts []types.Account
	NetworkConfig      types.NetworkConfig
	ContractAddresses  types.ContractAddresses
	ContractCaller     ContractCaller
}

var c Conf

// InitAndSetConf - initializes conf for given network
func InitAndSetConf(network *string) {
	c = NewConf(network)
}

// NewConf - create new conf
func NewConf(network *string) (c Conf) {
	// base path for all configuration // path is relative to main_test.go
	basePath := "config/" + *network
	c.GenesisValAccounts = ParseGenesisValidators(basePath)
	c.MainchainFaucet = ParseMainchainFaucet(basePath)
	c.MaticFaucet = ParseMaticFaucet(basePath)
	c.NetworkConfig = ParseNetworkConfig(basePath)
	c.ContractAddresses = ParseContractAddresses(basePath)
	c.ContractCaller, _ = NewContractCaller(c.NetworkConfig.EthRPCUrl, c.NetworkConfig.BorRPCUrl)
	return
}

// GetConf - returns conf
func GetConf() Conf {
	return c
}
