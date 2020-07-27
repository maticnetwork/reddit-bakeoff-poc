package helper

import (
	"encoding/json"
	"io/ioutil"

	"github.com/maticnetwork/reddit-bakeoff-poc/types"
)

// ParseGenesisValidators - Read signers.json and create accounts
func ParseGenesisValidators(basePath string) []types.Account {
	var accounts []types.Account
	signersFilePath := basePath + "/genesis-validators.json"

	// Read signers
	byteValue, err := ioutil.ReadFile(signersFilePath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(byteValue, &accounts)
	return accounts
}

// ParseMainchainFaucet - Read mainchain-faucet.json and create faucet account
func ParseMainchainFaucet(basePath string) types.Account {
	var faucetAccount types.Account
	faucetAccountFilePath := basePath + "/mainchain-faucet.json"

	// Read signers
	byteValue, err := ioutil.ReadFile(faucetAccountFilePath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(byteValue, &faucetAccount)
	return faucetAccount
}

// ParseMaticFaucet - Read matic-faucet.json and create faucet account
func ParseMaticFaucet(basePath string) types.Account {
	var faucetAccount types.Account
	faucetAccountFilePath := basePath + "/matic-faucet.json"

	// Read signers
	byteValue, err := ioutil.ReadFile(faucetAccountFilePath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(byteValue, &faucetAccount)
	return faucetAccount
}

// ParseNetworkConfig - Read network-config.json and create NetworkConfig
func ParseNetworkConfig(basePath string) types.NetworkConfig {
	var networkConfig types.NetworkConfig
	networkConfigFilePath := basePath + "/network-config.json"

	// Read network-config.json
	byteValue, err := ioutil.ReadFile(networkConfigFilePath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(byteValue, &networkConfig)
	return networkConfig
}

// ParseContractAddresses - Read addresses.json and create ContractAddress
func ParseContractAddresses(basePath string) types.ContractAddresses {
	var contractAddresses types.ContractAddresses
	contractAddressesFilePath := basePath + "/contractAddresses.json"

	// Read contract address
	byteValue, err := ioutil.ReadFile(contractAddressesFilePath)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(byteValue, &contractAddresses)
	return contractAddresses
}

// ParseRedditContracts - Read reditContracts.json and create Redit contracts
func ParseRedditContracts(basePath string) types.RedditContracts {
	var redditContracts types.RedditContracts
	redditContractsFilePath := basePath + "/reddit/redditContracts.json"

	// Read contract address
	byteValue, err := ioutil.ReadFile(redditContractsFilePath)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(byteValue, &redditContracts)
	return redditContracts
}

// ParseRedditConfig - Read reddit-config.json and create RedditConfig
func ParseRedditConfig(basePath string) types.RedditConfig {
	var redditConfig types.RedditConfig
	redditConfigFilePath := basePath + "/reddit/reddit-config.json"

	// Read reddit-config.json
	byteValue, err := ioutil.ReadFile(redditConfigFilePath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(byteValue, &redditConfig)
	return redditConfig
}

// ParseRedditBenchmark - Read reddit-benchmark.json and create RedditBenchmark
func ParseRedditBenchmark(basePath string) types.RedditBenchmark {
	var redditBenchmark types.RedditBenchmark
	redditConfigFilePath := basePath + "/reddit/reddit-benchmark.json"

	// Read reddit-config.json
	byteValue, err := ioutil.ReadFile(redditConfigFilePath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(byteValue, &redditBenchmark)
	return redditBenchmark
}

// ParseRedditAccounts - Read reddit-accounts.json and create accounts
func ParseRedditAccounts(basePath string) []types.Account {
	var accounts []types.Account
	signersFilePath := basePath + "/reddit/reddit-accounts.json"

	// Read signers
	byteValue, err := ioutil.ReadFile(signersFilePath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(byteValue, &accounts)
	return accounts
}
