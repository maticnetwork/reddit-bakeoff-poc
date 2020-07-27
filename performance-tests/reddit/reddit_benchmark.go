package reddit

import (
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/maticnetwork/bor/accounts/abi"
	"github.com/maticnetwork/bor/accounts/abi/bind"
	"github.com/maticnetwork/bor/common"
	"github.com/maticnetwork/bor/crypto"
	hmTypes "github.com/maticnetwork/heimdall/types"
	"github.com/maticnetwork/reddit-bakeoff-poc/helper"
	"github.com/maticnetwork/reddit-bakeoff-poc/types"
)

// BenchmarkReddit - add all Redit benchmark tests here
var BenchmarkReddit = func(b *testing.B) {
	b.Run("Initialize", initializeReddit())
	b.Run("Claim", claimBenchmark())
	b.Run("Transfer", transferBenchmark())
	b.Run("Subscribe", subscribeBenchmark())
}

var initializeReddit = func() func(*testing.B) {
	return func(b *testing.B) {
		fmt.Println("initializeRedditContracts")
		rc := helper.GetRedditConf()
		conf := helper.GetConf()

		distributionAddress := hmTypes.HexToHeimdallAddress(rc.RedditContracts.Distributions_v0).EthAddress()
		distributionsInstance, _ := rc.RedditCaller.GetDistributionsInstance(distributionAddress)

		subscriptionsAddress := hmTypes.HexToHeimdallAddress(rc.RedditContracts.Subscriptions_v0).EthAddress()
		subscriptionsInstance, _ := rc.RedditCaller.GetSubscriptionsInstance(subscriptionsAddress)

		subredditPointsAddress := hmTypes.HexToHeimdallAddress(rc.RedditContracts.SubredditPoints_v0).EthAddress()
		subredditPointsInstance, _ := rc.RedditCaller.GetSubredditpointsInstance(subredditPointsAddress)

		subredditPointsParentAddress := hmTypes.HexToHeimdallAddress(rc.RedditContracts.SubredditPointsParent).EthAddress()
		childChainAddress := hmTypes.HexToHeimdallAddress(conf.ContractAddresses.ChildChain.ChildChain).EthAddress()

		initialSupply := big.NewInt(0).SetUint64(rc.RedditConfig.InitialSupply)
		initialKarma := big.NewInt(0).SetUint64(rc.RedditConfig.InitialKarma)
		price := big.NewInt(0).SetUint64(rc.RedditConfig.SubscriptionPrice)
		duration := big.NewInt(0).SetUint64(rc.RedditConfig.Duration)
		renewBefore := big.NewInt(0).SetUint64(rc.RedditConfig.RenewBefore)

		ownerAddress := hmTypes.HexToHeimdallAddress(rc.RedditAccounts[0].Address).EthAddress()

		// initialize contracts
		if err := rc.RedditCaller.InitializeSubredditpoints(subredditPointsInstance, subredditPointsAddress, distributionAddress, ownerAddress, rc.RedditConfig.SubredditName, rc.RedditConfig.SubredditTokenName, rc.RedditConfig.SubredditTokenSymbol, subredditPointsParentAddress, childChainAddress, rc.RedditAccounts[0]); err != nil {
			b.Error(err)
			b.FailNow()
		}
		if err := rc.RedditCaller.InitializeDistributions(distributionsInstance, distributionAddress, subredditPointsAddress, initialSupply, initialKarma, ownerAddress, ownerAddress, rc.RedditAccounts[0]); err != nil {
			b.Error(err)
			b.FailNow()
		}
		if err := rc.RedditCaller.InitializeSubscriptions(subscriptionsInstance, subscriptionsAddress, subredditPointsAddress, price, duration, renewBefore, ownerAddress, rc.RedditAccounts[0]); err != nil {
			b.Error(err)
			b.FailNow()
		}
		if err := rc.RedditCaller.AddDefaultOperator(subredditPointsInstance, subredditPointsAddress, subscriptionsAddress, rc.RedditAccounts[0]); err != nil {
			b.Error(err)
			b.FailNow()
		}
		fmt.Println("Initialization Successful")
	}
}

var claimBenchmark = func() func(*testing.B) {
	return func(b *testing.B) {
		rc := helper.GetRedditConf()
		b.Log("BenchmarkClaims - ", "no of claims", rc.RedditBenchmark.NumClaims)

		b.N = rc.RedditBenchmark.NumClaims
		signerAcc, err := testDataClaim(b.N)
		if err != nil {
			b.Error(err)
			b.FailNow()
		}

		// contract instances
		distributionAddress := hmTypes.HexToHeimdallAddress(rc.RedditContracts.Distributions_v0).EthAddress()
		distributionsInstance, _ := rc.RedditCaller.GetDistributionsInstance(distributionAddress)

		// test data
		currentRound := big.NewInt(0)
		userKarma := big.NewInt(10000)

		fmt.Println("Setting up test data to run benchmark for claims", "currentTime", time.Now())
		var auths = make([]*bind.TransactOpts, b.N)
		var users = types.RandomAccounts(b.N)
		var signatures = make([][]byte, b.N)
		nonce := 0
		for j := 0; j < b.N; j++ {
			userAddr := hmTypes.HexToHeimdallAddress(users[j].Address).EthAddress()
			signature, err := claimSig(currentRound, userAddr, userKarma)
			if err != nil {
				b.Error(err)
				b.FailNow()
			}
			// create claim tx
			auth, err := rc.RedditCaller.CreateClaim(distributionAddress, currentRound, userAddr, userKarma, signature, signerAcc)
			if err != nil {
				b.Error(err)
				b.FailNow()
			}
			auth.Nonce = big.NewInt(int64(nonce))
			auths[j] = auth
			signatures[j] = signature
			nonce++
		}
		time.Sleep(20 * time.Second)

		// run benchmark
		b.ResetTimer()
		var index int32 = -1 // Count up to N but atomically
		fmt.Println("Starting benchmark for claims", "currentTime", time.Now())
		claims := func(pb *testing.PB) {
			for pb.Next() {
				i := atomic.AddInt32(&index, 1)
				userAddr := hmTypes.HexToHeimdallAddress(users[i].Address).EthAddress()
				// transfer amount from sender to receiver
				txHash, err := rc.RedditCaller.SendClaim(distributionsInstance, auths[i], currentRound, userAddr, userKarma, signatures[i])
				if err != nil {
					b.Error(err)
					b.FailNow()
				}
				fmt.Println("Claim Successful", "txHash", txHash, "nonce", auths[i].Nonce, "currentRound", currentRound, "userAddr", userAddr.Hex(), "userKarma", userKarma, "signature", hmTypes.BytesToHexBytes(signatures[i]))
			}
		}
		b.RunParallel(claims)
		b.SetParallelism(300)
		printClaimStats(rc.RedditBenchmark.NumClaims)
	}
}

func testDataClaim(numClaims int) (user types.Account, err error) {
	user = types.RandomAccounts(1)[0] // create 1 random account
	feePerTx, _ := big.NewInt(0).SetString("20000000000000000", 10)
	numOfTx := big.NewInt(0).SetInt64(int64(numClaims))
	totalFee := big.NewInt(0).Mul(numOfTx, feePerTx)
	err = helper.GetTestEthMatic(totalFee.Uint64(), user.Address)
	time.Sleep(20 * time.Second)
	return
}

func printClaimStats(numClaims int) {
	user, _ := testDataClaim(1)
	rc := helper.GetRedditConf()
	userAddr := hmTypes.HexToHeimdallAddress(user.Address).EthAddress()
	// contract instances
	distributionAddress := hmTypes.HexToHeimdallAddress(rc.RedditContracts.Distributions_v0).EthAddress()

	signature, _ := claimSig(big.NewInt(0), userAddr, big.NewInt(100))
	auth, _ := rc.RedditCaller.CreateClaimWithEstimate(distributionAddress, big.NewInt(0), userAddr, big.NewInt(100), signature, user)
	fmt.Println("Statistics", "gasPrice", auth.GasPrice.Uint64(), "gasLimit", auth.GasLimit, "costPerClaim", auth.GasPrice.Uint64()*auth.GasLimit, "totalClaims", numClaims)
}

var claimSig = func(round *big.Int, account common.Address, karma *big.Int) (signature []byte, err error) {
	rc := helper.GetRedditConf()
	karmaSource := rc.RedditAccounts[0]
	// sign data
	name := strings.ToLower(rc.RedditConfig.SubredditName)

	// get priv key
	key := karmaSource.PrivKey[2:]

	// create ecdsa private key
	ecdsaPrivateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return
	}

	uintType, _ := abi.NewType("uint256", nil)
	addressType, _ := abi.NewType("address", nil)
	stringType, _ := abi.NewType("string", nil)

	arguments := abi.Arguments{
		{
			Type: stringType,
		},
		{
			Type: uintType,
		},
		{
			Type: addressType,
		},
		{
			Type: uintType,
		},
	}

	dataBytes, _ := arguments.Pack(
		name,
		round,
		account,
		karma,
	)

	hash := crypto.Keccak256(dataBytes)
	dataHash := crypto.Keccak256Hash(append([]byte("\x19Ethereum Signed Message:\n32"), hash...))

	signature, err = crypto.Sign(dataHash.Bytes(), ecdsaPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	v := signature[64] + 27
	signature = append(signature[:64], v)
	return
}

var transferBenchmark = func() func(*testing.B) {
	return func(b *testing.B) {
		rc := helper.GetRedditConf()
		b.Log("BenchmarkTransfers -", "no of transfers", rc.RedditBenchmark.NumTransfers)

		// contract instances
		subredditPointsAddress := hmTypes.HexToHeimdallAddress(rc.RedditContracts.SubredditPoints_v0).EthAddress()
		subredditPointsInstance, _ := rc.RedditCaller.GetSubredditpointsInstance(subredditPointsAddress)

		// setup
		b.N = rc.RedditBenchmark.NumTransfers
		sender, err := testDataTransfer(b.N)
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
		recipient := types.RandomAccounts(1)[0]
		amount := big.NewInt(1)

		// test data
		userData := hmTypes.ZeroHeimdallAddress
		operatorData := hmTypes.ZeroHeimdallAddress
		senderAddr := hmTypes.HexToHeimdallAddress(sender.Address).EthAddress()
		recipientAddr := hmTypes.HexToHeimdallAddress(recipient.Address).EthAddress()
		fmt.Println("Setting up test data to run benchmark for transfers", "currentTime", time.Now())
		var auths = make([]*bind.TransactOpts, b.N)
		nonce := 1
		for j := 0; j < b.N; j++ {
			// create transfer tx
			auth, err := rc.RedditCaller.CreateTransfer(subredditPointsInstance, subredditPointsAddress, senderAddr, recipientAddr, amount, userData.Bytes(), operatorData.Bytes(), sender)
			if err != nil {
				b.Error(err)
				b.FailNow()
			}
			auth.Nonce = big.NewInt(int64(nonce))
			auths[j] = auth
			nonce++
		}
		time.Sleep(20 * time.Second)

		// run benchmark
		b.ResetTimer()
		fmt.Println("Starting benchmark for transfer", "currentTime", time.Now())
		var index int32 = -1 // Count up to N but atomically

		transfers := func(pb *testing.PB) {
			for pb.Next() {
				i := atomic.AddInt32(&index, 1)
				// transfer amount from sender to receiver
				txHash, err := rc.RedditCaller.Transfer(subredditPointsInstance, auths[i], senderAddr, recipientAddr, amount, userData.Bytes(), operatorData.Bytes())
				if err != nil {
					b.Error(err)
					b.FailNow()
				}
				fmt.Println("Transfer Successful", "txHash", txHash, "nonce", auths[i].Nonce, "senderAddr", senderAddr.Hex(), "recipientAddr", recipientAddr.Hex(), "amount", amount.Uint64())
			}
		}
		b.RunParallel(transfers)
		b.SetParallelism(300)
		printTransferStats(rc.RedditBenchmark.NumTransfers)
	}
}

func testDataTransfer(numTransfer int) (sender types.Account, err error) {
	rc := helper.GetRedditConf()
	sender = types.RandomAccounts(1)[0] // create 1 random account
	feePerTx, _ := big.NewInt(0).SetString("20000000000000000", 10)
	numOfTx := big.NewInt(0).SetInt64(int64(numTransfer) + 1)
	totalFee := big.NewInt(0).Mul(numOfTx, feePerTx)
	err = helper.GetTestEthMatic(totalFee.Uint64(), sender.Address)
	fmt.Println("Sender", sender, "totalFee", totalFee.Uint64(), "error", err)
	// wait for eth to be sent
	time.Sleep(20 * time.Second)
	// contract instances
	distributionAddress := hmTypes.HexToHeimdallAddress(rc.RedditContracts.Distributions_v0).EthAddress()
	distributionsInstance, _ := rc.RedditCaller.GetDistributionsInstance(distributionAddress)

	// test data
	currentRound := big.NewInt(0)
	userKarma := big.NewInt(int64(numTransfer))
	senderAddr := hmTypes.HexToHeimdallAddress(sender.Address).EthAddress()
	signature, err := claimSig(currentRound, senderAddr, userKarma)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	// create claim tx
	auth, err := rc.RedditCaller.CreateClaim(distributionAddress, currentRound, senderAddr, userKarma, signature, sender)
	if err != nil {
		return
	}
	// send claim tx
	txHash, err := rc.RedditCaller.SendClaim(distributionsInstance, auth, currentRound, senderAddr, userKarma, signature)
	if err != nil {
		return
	}
	fmt.Println("Claim Successful", "txHash", txHash, "currentRound", currentRound, "userAddr", senderAddr.Hex(), "userKarma", userKarma, "signature", hmTypes.BytesToHexBytes(signature))
	return
}

func printTransferStats(numTransfers int) {
	sender, _ := testDataTransfer(1)
	recipient := types.RandomAccounts(1)[0]
	rc := helper.GetRedditConf()
	// contract instances
	subredditPointsAddress := hmTypes.HexToHeimdallAddress(rc.RedditContracts.SubredditPoints_v0).EthAddress()
	subredditPointsInstance, _ := rc.RedditCaller.GetSubredditpointsInstance(subredditPointsAddress)

	senderAddr := hmTypes.HexToHeimdallAddress(sender.Address).EthAddress()
	recipientAddr := hmTypes.HexToHeimdallAddress(recipient.Address).EthAddress()

	auth, _ := rc.RedditCaller.CreateTransferWithEstimate(subredditPointsInstance, subredditPointsAddress, senderAddr, recipientAddr, big.NewInt(1), hmTypes.ZeroHeimdallAddress.Bytes(), hmTypes.ZeroHeimdallAddress.Bytes(), sender)
	fmt.Println("Statistics", "gasPrice", auth.GasPrice.Uint64(), "gasLimit", auth.GasLimit, "costPerTransfer", auth.GasPrice.Uint64()*auth.GasLimit, "totalTransfer", numTransfers)
}

var subscribeBenchmark = func() func(*testing.B) {
	return func(b *testing.B) {
		rc := helper.GetRedditConf()
		b.Log("BenchmarkSubscribe -", "no of subscribe", rc.RedditBenchmark.NumSubscribe)

		b.N = rc.RedditBenchmark.NumSubscribe
		payer, err := testDataSubscribe(int64(b.N))
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
		recipient := types.RandomAccounts(1)[0]

		// contract instances
		subscriptionsAddress := hmTypes.HexToHeimdallAddress(rc.RedditContracts.Subscriptions_v0).EthAddress()
		subscriptionsInstance, _ := rc.RedditCaller.GetSubscriptionsInstance(subscriptionsAddress)

		// test data
		recipientAddr := hmTypes.HexToHeimdallAddress(recipient.Address).EthAddress()
		renewable := false

		fmt.Println("Setting up test data to run benchmark for subscriptions", "currentTime", time.Now())
		var auths = make([]*bind.TransactOpts, b.N)
		nonce := 1
		for j := 0; j < b.N; j++ {
			// create subscribe tx
			auth, err := rc.RedditCaller.CreateSubscribe(subscriptionsInstance, subscriptionsAddress, recipientAddr, renewable, payer)
			if err != nil {
				b.Error(err)
				b.FailNow()
			}
			auth.Nonce = big.NewInt(int64(nonce))
			auths[j] = auth
			nonce++
		}

		time.Sleep(20 * time.Second)
		// run benchmark
		b.ResetTimer()
		fmt.Println("Starting benchmark for subscriptions", "currentTime", time.Now())
		var index int32 = -1 // Count up to N but atomically

		subscriptions := func(pb *testing.PB) {
			for pb.Next() {
				i := atomic.AddInt32(&index, 1)
				// transfer amount from sender to receiver
				txHash, err := rc.RedditCaller.SendSubscribe(subscriptionsInstance, auths[i], recipientAddr, renewable)
				if err != nil {
					b.Error(err)
					b.FailNow()
				}
				fmt.Println("Subscribe Successful", "txHash", txHash, "nonce", auths[i].Nonce, "recipientAddr", recipientAddr.Hex())
			}
		}
		b.RunParallel(subscriptions)
		b.SetParallelism(300)
		fmt.Println("Benchmark completed", "currentTime", time.Now())
		printSubscribeStats(rc.RedditBenchmark.NumSubscribe)
	}
}

func testDataSubscribe(numSubscribe int64) (payer types.Account, err error) {
	rc := helper.GetRedditConf()
	payer = types.RandomAccounts(1)[0] // create 1 random account
	feePerTx, _ := big.NewInt(0).SetString("20000000000000000", 10)
	numOfTx := big.NewInt(0).SetInt64(numSubscribe)
	totalFee := big.NewInt(0).Mul(numOfTx, feePerTx)
	helper.GetTestEthMatic(totalFee.Uint64(), payer.Address)
	fmt.Println("Payer", payer, "totalfee", totalFee.Uint64())
	time.Sleep(20 * time.Second)
	// contract instances
	distributionAddress := hmTypes.HexToHeimdallAddress(rc.RedditContracts.Distributions_v0).EthAddress()
	distributionsInstance, _ := rc.RedditCaller.GetDistributionsInstance(distributionAddress)

	// test data
	currentRound := big.NewInt(0)
	userKarma := big.NewInt(numSubscribe * int64(rc.RedditConfig.SubscriptionPrice))
	payerAddr := hmTypes.HexToHeimdallAddress(payer.Address).EthAddress()
	signature, err := claimSig(currentRound, payerAddr, userKarma)
	if err != nil {
		return
	}
	// create claim tx
	auth, _ := rc.RedditCaller.CreateClaim(distributionAddress, currentRound, payerAddr, userKarma, signature, payer)

	// send claim tx
	rc.RedditCaller.SendClaim(distributionsInstance, auth, currentRound, payerAddr, userKarma, signature)
	return
}

func printSubscribeStats(numSubscribe int) {
	payer, _ := testDataSubscribe(1)
	recipient := types.RandomAccounts(1)[0]
	rc := helper.GetRedditConf()
	// contract instances
	subscriptionsAddress := hmTypes.HexToHeimdallAddress(rc.RedditContracts.Subscriptions_v0).EthAddress()
	subscriptionsInstance, _ := rc.RedditCaller.GetSubscriptionsInstance(subscriptionsAddress)

	recipientAddr := hmTypes.HexToHeimdallAddress(recipient.Address).EthAddress()

	auth, _ := rc.RedditCaller.CreateSubscribeWithEstimate(subscriptionsInstance, subscriptionsAddress, recipientAddr, false, payer)
	fmt.Println("Statistics", "gasPrice", auth.GasPrice.Uint64(), "gasLimit", auth.GasLimit, "costPerSubscribe", auth.GasPrice.Uint64()*auth.GasLimit, "totalSubscribes", numSubscribe)
}
