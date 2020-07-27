package helper

import (
	"fmt"
	"math/big"

	"github.com/maticnetwork/bor/accounts/abi"
	"github.com/maticnetwork/bor/accounts/abi/bind"
	"github.com/maticnetwork/bor/common"
	"github.com/maticnetwork/bor/ethclient"
	"github.com/maticnetwork/bor/rpc"

	"github.com/maticnetwork/reddit-bakeoff-poc/artifacts/distributions"
	"github.com/maticnetwork/reddit-bakeoff-poc/artifacts/subredditpoints"
	"github.com/maticnetwork/reddit-bakeoff-poc/artifacts/subscriptions"
	"github.com/maticnetwork/reddit-bakeoff-poc/types"
)

// RedditCaller - api's to interact with Reddit contracts
type RedditCaller struct {
	MainChainClient  *ethclient.Client
	MaticChainClient *ethclient.Client
	MainRPCClient    *rpc.Client
	MaticRPCClient   *rpc.Client

	DistributionsABI   abi.ABI
	SubscriptionsABI   abi.ABI
	SubredditpointsABI abi.ABI

	ContractInstanceCache map[common.Address]interface{}
}

// NewRedditCaller contract caller
func NewRedditCaller(ethRPCUrl string, borRPCUrl string) (redditCallerObj RedditCaller, err error) {
	redditCallerObj.MainChainClient, redditCallerObj.MainRPCClient = GetMainClient(ethRPCUrl)
	redditCallerObj.MaticChainClient, redditCallerObj.MaticRPCClient = GetMaticClient(borRPCUrl)

	if redditCallerObj.DistributionsABI, err = getABI(string(distributions.DistributionsABI)); err != nil {
		return
	}

	if redditCallerObj.SubscriptionsABI, err = getABI(string(subscriptions.SubscriptionsABI)); err != nil {
		return
	}

	if redditCallerObj.SubredditpointsABI, err = getABI(string(subredditpoints.SubredditpointsABI)); err != nil {
		return
	}

	redditCallerObj.ContractInstanceCache = make(map[common.Address]interface{})
	return
}

// GetDistributionsInstance returns distributions contract instance for selected base chain
func (rc *RedditCaller) GetDistributionsInstance(distributionsAddress common.Address) (*distributions.Distributions, error) {
	contractInstance, ok := rc.ContractInstanceCache[distributionsAddress]
	if !ok {
		ci, err := distributions.NewDistributions(distributionsAddress, rc.MaticChainClient)
		rc.ContractInstanceCache[distributionsAddress] = ci
		return ci, err
	}
	return contractInstance.(*distributions.Distributions), nil
}

// GetSubscriptionsInstance returns subscriptions contract instance for selected base chain
func (rc *RedditCaller) GetSubscriptionsInstance(subscriptionsAddress common.Address) (*subscriptions.Subscriptions, error) {
	contractInstance, ok := rc.ContractInstanceCache[subscriptionsAddress]
	if !ok {
		ci, err := subscriptions.NewSubscriptions(subscriptionsAddress, rc.MaticChainClient)
		rc.ContractInstanceCache[subscriptionsAddress] = ci
		return ci, err
	}
	return contractInstance.(*subscriptions.Subscriptions), nil
}

// GetSubredditpointsInstance returns subredditpoints contract instance for selected base chain
func (rc *RedditCaller) GetSubredditpointsInstance(subredditpointsAddress common.Address) (*subredditpoints.Subredditpoints, error) {
	contractInstance, ok := rc.ContractInstanceCache[subredditpointsAddress]
	if !ok {
		ci, err := subredditpoints.NewSubredditpoints(subredditpointsAddress, rc.MaticChainClient)
		rc.ContractInstanceCache[subredditpointsAddress] = ci
		return ci, err
	}
	return contractInstance.(*subredditpoints.Subredditpoints), nil
}

// InitializeDistributions - initialize distributions contract
func (rc *RedditCaller) InitializeDistributions(distributionsInstance *distributions.Distributions, distributionsAddress common.Address, subredditpoints common.Address, initialSupply *big.Int, initialKarma *big.Int, ownerAddress common.Address, karmaSource common.Address, signer types.Account) error {
	// pack data based on method definition
	data, err := rc.DistributionsABI.Pack("initializecontract", ownerAddress, subredditpoints, initialSupply, initialKarma, karmaSource)
	if err != nil {
		return err
	}

	auth, err := GenerateAuthObj(rc.MaticChainClient, distributionsAddress, data, signer, nil)
	if err != nil {
		return err
	}

	tx, err := distributionsInstance.Initializecontract(
		auth,
		ownerAddress,
		subredditpoints,
		initialSupply,
		initialKarma,
		karmaSource,
	)

	if err != nil {
		return err
	}
	fmt.Println("initialize distributions sucessfully", "txHash", tx.Hash().String())
	return nil
}

// InitializeSubscriptions - initialize subscriptions contract
func (rc *RedditCaller) InitializeSubscriptions(subscriptionsInstance *subscriptions.Subscriptions, subscriptionsAddress common.Address, subredditpoints common.Address, price *big.Int, duration *big.Int, renewBefore *big.Int, ownerAddress common.Address, signer types.Account) error {
	// pack data based on method definition
	data, err := rc.SubscriptionsABI.Pack("initializecontract", ownerAddress, subredditpoints, price, duration, renewBefore)
	if err != nil {
		return err
	}

	auth, err := GenerateAuthObj(rc.MaticChainClient, subscriptionsAddress, data, signer, nil)
	if err != nil {
		return err
	}

	tx, err := subscriptionsInstance.Initializecontract(
		auth,
		ownerAddress,
		subredditpoints,
		price,
		duration,
		renewBefore,
	)

	if err != nil {
		return err
	}

	fmt.Println("initialize subscriptions sucessfully", "txHash", tx.Hash().String())
	return nil
}

// InitializeSubredditpoints - initialize Subredditpoints contract
func (rc *RedditCaller) InitializeSubredditpoints(subredditpointsInstance *subredditpoints.Subredditpoints, subredditpointsAddress common.Address, distributionAddress common.Address, ownerAddress common.Address, subredditName string, subredditTokenName string, subredditTokenSymbol string, rootToken common.Address, childChain common.Address, signer types.Account) error {
	// pack data based on method definition
	data, err := rc.SubredditpointsABI.Pack("initializecontract", ownerAddress, distributionAddress, subredditName, subredditTokenName, subredditTokenSymbol, []common.Address{}, rootToken, childChain)
	if err != nil {
		return err
	}

	auth, err := GenerateAuthObj(rc.MaticChainClient, subredditpointsAddress, data, signer, nil)
	if err != nil {
		return err
	}

	tx, err := subredditpointsInstance.Initializecontract(
		auth,
		ownerAddress,
		distributionAddress,
		subredditName,
		subredditTokenName,
		subredditTokenSymbol,
		[]common.Address{},
		rootToken,
		childChain,
	)

	if err != nil {
		return err
	}

	fmt.Println("initialize subredditpoints sucessfully", "txHash", tx.Hash().String())
	return nil
}

// AddDefaultOperator - Adds subscriptionsAddress as DefaultOperator
func (rc *RedditCaller) AddDefaultOperator(subredditpointsInstance *subredditpoints.Subredditpoints, subredditpointsAddress common.Address, subscriptionsAddress common.Address, signer types.Account) error {
	// pack data based on method definition
	data, err := rc.SubredditpointsABI.Pack("addDefaultOperator", subscriptionsAddress)
	if err != nil {
		return err
	}

	auth, err := GenerateAuthObj(rc.MaticChainClient, subredditpointsAddress, data, signer, nil)
	if err != nil {
		return err
	}

	tx, err := subredditpointsInstance.AddDefaultOperator(
		auth,
		subscriptionsAddress,
	)

	if err != nil {
		return err
	}

	fmt.Println("Added DefaultOperator sucessfully", "txHash", tx.Hash().String())
	return nil
}

// CreateClaim - creates auth
func (rc *RedditCaller) CreateClaim(distributionsAddress common.Address, round *big.Int, account common.Address, karma *big.Int, signature []byte, signer types.Account) (auth *bind.TransactOpts, err error) {

	// pack data based on method definition
	data, err := rc.DistributionsABI.Pack("claim", round, account, karma, signature)
	if err != nil {
		return
	}

	auth, err = GenerateSimpleAuthObj(rc.MaticChainClient, distributionsAddress, data, signer)
	if err != nil {
		fmt.Println("Error generating obj", "error", err)
	}
	return
}

// CreateClaimWithEstimate - creates auth
func (rc *RedditCaller) CreateClaimWithEstimate(distributionsAddress common.Address, round *big.Int, account common.Address, karma *big.Int, signature []byte, signer types.Account) (auth *bind.TransactOpts, err error) {

	// pack data based on method definition
	data, err := rc.DistributionsABI.Pack("claim", round, account, karma, signature)
	if err != nil {
		return
	}

	auth, err = GenerateAuthObj(rc.MaticChainClient, distributionsAddress, data, signer, nil)
	if err != nil {
		fmt.Println("Error generating obj", "error", err)
	}
	return
}

// SendClaim - calls claim method on distributions
func (rc *RedditCaller) SendClaim(distributionsInstance *distributions.Distributions, auth *bind.TransactOpts, round *big.Int, account common.Address, karma *big.Int, signature []byte) (txHash string, err error) {
	tx, err := distributionsInstance.Claim(
		auth,
		round,
		account,
		karma,
		signature,
	)

	if err != nil {
		return
	}

	txHash = tx.Hash().String()
	return
}

// CreateTransfer - creates auth
func (rc *RedditCaller) CreateTransfer(subredditpointsInstance *subredditpoints.Subredditpoints, subredditpointsAddress common.Address, sender common.Address, recipient common.Address, amount *big.Int, userData []byte, operatorData []byte, signer types.Account) (auth *bind.TransactOpts, err error) {

	// pack data based on method definition
	data, err := rc.SubredditpointsABI.Pack("operatorSend", sender, recipient, amount, userData, operatorData)
	if err != nil {
		return
	}

	auth, err = GenerateSimpleAuthObj(rc.MaticChainClient, subredditpointsAddress, data, signer)

	return
}

// CreateTransferWithEstimate - creates auth
func (rc *RedditCaller) CreateTransferWithEstimate(subredditpointsInstance *subredditpoints.Subredditpoints, subredditpointsAddress common.Address, sender common.Address, recipient common.Address, amount *big.Int, userData []byte, operatorData []byte, signer types.Account) (auth *bind.TransactOpts, err error) {

	// pack data based on method definition
	data, err := rc.SubredditpointsABI.Pack("operatorSend", sender, recipient, amount, userData, operatorData)
	if err != nil {
		return
	}

	auth, err = GenerateAuthObj(rc.MaticChainClient, subredditpointsAddress, data, signer, nil)

	return
}

// Transfer - calls operatorSend method on subredditpoints
func (rc *RedditCaller) Transfer(subredditpointsInstance *subredditpoints.Subredditpoints, auth *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int, userData []byte, operatorData []byte) (txHash string, err error) {

	tx, err := subredditpointsInstance.OperatorSend(
		auth,
		sender,
		recipient,
		amount,
		userData,
		operatorData,
	)

	if err != nil {
		return
	}

	txHash = tx.Hash().String()
	return
}

// CreateSubscribe - creates auth
func (rc *RedditCaller) CreateSubscribe(subscriptionsInstance *subscriptions.Subscriptions, subscriptionsAddress common.Address, recipient common.Address, renewable bool, signer types.Account) (auth *bind.TransactOpts, err error) {
	// pack data based on method definition
	data, err := rc.SubscriptionsABI.Pack("subscribe", recipient, renewable)
	if err != nil {
		return
	}

	auth, err = GenerateSimpleAuthObj(rc.MaticChainClient, subscriptionsAddress, data, signer)
	return
}

// CreateSubscribeWithEstimate - creates auth
func (rc *RedditCaller) CreateSubscribeWithEstimate(subscriptionsInstance *subscriptions.Subscriptions, subscriptionsAddress common.Address, recipient common.Address, renewable bool, signer types.Account) (auth *bind.TransactOpts, err error) {
	// pack data based on method definition
	data, err := rc.SubscriptionsABI.Pack("subscribe", recipient, renewable)
	if err != nil {
		return
	}

	auth, err = GenerateAuthObj(rc.MaticChainClient, subscriptionsAddress, data, signer, nil)
	return
}

// SendSubscribe - calls subscribe on subscriptions
func (rc *RedditCaller) SendSubscribe(subscriptionsInstance *subscriptions.Subscriptions, auth *bind.TransactOpts, recipient common.Address, renewable bool) (txHash string, err error) {

	tx, err := subscriptionsInstance.Subscribe(
		auth,
		recipient,
		renewable,
	)

	if err != nil {
		return
	}
	txHash = tx.Hash().String()
	return
}
