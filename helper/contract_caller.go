package helper

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	ethereum "github.com/maticnetwork/bor"
	"github.com/maticnetwork/bor/accounts/abi"
	"github.com/maticnetwork/bor/accounts/abi/bind"
	"github.com/maticnetwork/bor/common"
	"github.com/maticnetwork/bor/crypto"
	"github.com/maticnetwork/bor/ethclient"
	"github.com/maticnetwork/bor/rpc"
	"golang.org/x/sync/errgroup"

	hmTypes "github.com/maticnetwork/heimdall/types"
	"github.com/maticnetwork/reddit-bakeoff-poc/artifacts/childerc20"
	"github.com/maticnetwork/reddit-bakeoff-poc/artifacts/depositmanager"
	"github.com/maticnetwork/reddit-bakeoff-poc/artifacts/erc20"
	"github.com/maticnetwork/reddit-bakeoff-poc/artifacts/erc20predicate"
	"github.com/maticnetwork/reddit-bakeoff-poc/artifacts/mrc20"
	"github.com/maticnetwork/reddit-bakeoff-poc/artifacts/rootchain"
	"github.com/maticnetwork/reddit-bakeoff-poc/artifacts/slashmanager"
	"github.com/maticnetwork/reddit-bakeoff-poc/artifacts/stakemanager"
	"github.com/maticnetwork/reddit-bakeoff-poc/artifacts/stakinginfo"
	"github.com/maticnetwork/reddit-bakeoff-poc/artifacts/statereceiver"
	"github.com/maticnetwork/reddit-bakeoff-poc/artifacts/statesender"
	"github.com/maticnetwork/reddit-bakeoff-poc/artifacts/validatorset"
	"github.com/maticnetwork/reddit-bakeoff-poc/artifacts/validatorshare"
	"github.com/maticnetwork/reddit-bakeoff-poc/artifacts/withdrawmanager"

	ethTypes "github.com/maticnetwork/bor/core/types"
	hmHelper "github.com/maticnetwork/heimdall/helper"

	"github.com/maticnetwork/reddit-bakeoff-poc/types"
)

// ContractCaller contract caller
type ContractCaller struct {
	MainChainClient    *ethclient.Client
	MaticChainClient   *ethclient.Client
	MainRPCClient      *rpc.Client
	MaticRPCClient     *rpc.Client
	RootChainABI       abi.ABI
	StakingInfoABI     abi.ABI
	DepositManagerABI  abi.ABI
	WithdrawManagerABI abi.ABI
	ValidatorSetABI    abi.ABI
	ValidatorShareABI  abi.ABI
	StateReceiverABI   abi.ABI
	StateSenderABI     abi.ABI
	StakeManagerABI    abi.ABI
	SlashManagerABI    abi.ABI
	MaticTokenABI      abi.ABI
	ChildERC20ABI      abi.ABI
	MRC20ABI           abi.ABI

	ERC20PredicateABI abi.ABI

	ContractInstanceCache map[common.Address]interface{}
}

const WITHDRAW_EVENT_SIG = "0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f"

// NewContractCaller contract caller
func NewContractCaller(ethRPCUrl string, borRPCUrl string) (contractCallerObj ContractCaller, err error) {
	contractCallerObj.MainChainClient, contractCallerObj.MainRPCClient = GetMainClient(ethRPCUrl)
	fmt.Println("borRPCUrl", borRPCUrl)
	contractCallerObj.MaticChainClient, contractCallerObj.MaticRPCClient = GetMaticClient(borRPCUrl)

	//
	// ABIs
	//

	if contractCallerObj.RootChainABI, err = getABI(string(rootchain.RootchainABI)); err != nil {
		return
	}

	if contractCallerObj.DepositManagerABI, err = getABI(string(depositmanager.DepositmanagerABI)); err != nil {
		return
	}

	if contractCallerObj.WithdrawManagerABI, err = getABI(string(withdrawmanager.WithdrawmanagerABI)); err != nil {
		return
	}

	if contractCallerObj.StakingInfoABI, err = getABI(string(stakinginfo.StakinginfoABI)); err != nil {
		return
	}

	if contractCallerObj.ValidatorSetABI, err = getABI(string(validatorset.ValidatorsetABI)); err != nil {
		return
	}

	if contractCallerObj.ValidatorShareABI, err = getABI(string(validatorshare.ValidatorshareABI)); err != nil {
		return
	}

	if contractCallerObj.StateReceiverABI, err = getABI(string(statereceiver.StatereceiverABI)); err != nil {
		return
	}

	if contractCallerObj.StateSenderABI, err = getABI(string(statesender.StatesenderABI)); err != nil {
		return
	}

	if contractCallerObj.StakeManagerABI, err = getABI(string(stakemanager.StakemanagerABI)); err != nil {
		return
	}

	if contractCallerObj.SlashManagerABI, err = getABI(string(slashmanager.SlashmanagerABI)); err != nil {
		return
	}

	if contractCallerObj.MaticTokenABI, err = getABI(string(erc20.Erc20ABI)); err != nil {
		return
	}

	if contractCallerObj.ChildERC20ABI, err = getABI(string(childerc20.Childerc20ABI)); err != nil {
		return
	}

	if contractCallerObj.MRC20ABI, err = getABI(string(mrc20.Mrc20ABI)); err != nil {
		return
	}

	if contractCallerObj.ERC20PredicateABI, err = getABI(string(erc20predicate.Erc20predicateABI)); err != nil {
		return
	}

	contractCallerObj.ContractInstanceCache = make(map[common.Address]interface{})

	return
}

// GetRootChainInstance returns RootChain contract instance for selected base chain
func (c *ContractCaller) GetValidatorSetInstance(validatorSetAddress common.Address) (*validatorset.Validatorset, error) {
	contractInstance, ok := c.ContractInstanceCache[validatorSetAddress]
	if !ok {
		ci, err := validatorset.NewValidatorset(validatorSetAddress, c.MaticChainClient)
		c.ContractInstanceCache[validatorSetAddress] = ci
		return ci, err
	}
	return contractInstance.(*validatorset.Validatorset), nil
}

func (c *ContractCaller) GetSpanDetails(id *big.Int, validatorSetInstance *validatorset.Validatorset) (number, startBlock, endBlock *big.Int, err error) {
	d, err := validatorSetInstance.GetSpan(nil, id)
	return d.Number, d.StartBlock, d.EndBlock, err
}

// GetRootChainInstance returns RootChain contract instance for selected base chain
func (c *ContractCaller) GetRootChainInstance(rootchainAddress common.Address) (*rootchain.Rootchain, error) {
	contractInstance, ok := c.ContractInstanceCache[rootchainAddress]
	if !ok {
		ci, err := rootchain.NewRootchain(rootchainAddress, c.MainChainClient)
		c.ContractInstanceCache[rootchainAddress] = ci
		return ci, err
	}
	return contractInstance.(*rootchain.Rootchain), nil
}

// GetStakeManagerInstance returns stakinginfo contract instance for selected base chain
func (c *ContractCaller) GetStakeManagerInstance(stakingManagerAddress common.Address) (*stakemanager.Stakemanager, error) {
	return stakemanager.NewStakemanager(stakingManagerAddress, c.MainChainClient)
}

// GetDepositManagerInstance returns depositManager contract instance
func (c *ContractCaller) GetDepositManagerInstance(depositManagerAddress common.Address) (*depositmanager.Depositmanager, error) {
	return depositmanager.NewDepositmanager(depositManagerAddress, c.MainChainClient)
}

// GetWithdrawManagerInstance returns depositManager contract instance
func (c *ContractCaller) GetWithdrawManagerInstance(withdrawManagerAddress common.Address) (*withdrawmanager.Withdrawmanager, error) {
	return withdrawmanager.NewWithdrawmanager(withdrawManagerAddress, c.MainChainClient)
}

// GetStakingInfoInstance returns stakinginfo contract instance for selected base chain
func (c *ContractCaller) GetStakingInfoInstance(stakingInfoAddress common.Address) (*stakinginfo.Stakinginfo, error) {
	contractInstance, ok := c.ContractInstanceCache[stakingInfoAddress]
	if !ok {
		ci, err := stakinginfo.NewStakinginfo(stakingInfoAddress, c.MainChainClient)
		c.ContractInstanceCache[stakingInfoAddress] = ci
		return ci, err
	}
	return contractInstance.(*stakinginfo.Stakinginfo), nil
}

// GetERC20TokenInstance returns stakinginfo contract instance for selected base chain
func (c *ContractCaller) GetERC20TokenInstance(maticTokenAddress common.Address, parent bool) (*erc20.Erc20, error) {
	if parent {
		return erc20.NewErc20(maticTokenAddress, c.MainChainClient)
	}
	return erc20.NewErc20(maticTokenAddress, c.MaticChainClient)
}

// GetChildERC20Instance returns ChildERC20 contract instance
func (c *ContractCaller) GetChildERC20Instance(childErc20Address common.Address) (*childerc20.Childerc20, error) {
	return childerc20.NewChilderc20(childErc20Address, c.MaticChainClient)
}

// GetMRC20Instance returns MRC20 contract instance
func (c *ContractCaller) GetMRC20Instance(mrc20Address common.Address) (*mrc20.Mrc20, error) {
	return mrc20.NewMrc20(mrc20Address, c.MaticChainClient)
}

// GetMRC20Instance returns MRC20 contract instance
func (c *ContractCaller) GetERC20PredicateInstance(erc20PredicateAddress common.Address) (*erc20predicate.Erc20predicate, error) {
	return erc20predicate.NewErc20predicate(erc20PredicateAddress, c.MainChainClient)
}

func (c *ContractCaller) GetValidatorShareInstance(validatorAddress common.Address) (*validatorshare.Validatorshare, error) {
	return validatorshare.NewValidatorshare(validatorAddress, c.MainChainClient)
}

func (c *ContractCaller) GetValidatorReward(contractAddress, validatorAddress common.Address) (*big.Int, error) {
	vShare, err := c.GetValidatorShareInstance(contractAddress)
	if err != nil {
		return nil, err
	}

	bobj := bind.CallOpts{From: validatorAddress}
	return vShare.Rewards(&bobj)
}

func (c *ContractCaller) StakeFor(val common.Address, stakeAmount *big.Int, feeAmount *big.Int, acceptDelegation bool, stakeManagerAddress common.Address, stakeManagerInstance *stakemanager.Stakemanager, signer types.Account) error {
	// get priv key
	key := signer.PrivKey[2:]

	// create ecdsa private key
	ecdsaPrivateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return err
	}

	publicKey := ecdsaPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	signerPubkeyBytes := crypto.FromECDSAPub(publicKeyECDSA)[1:]

	// pack data based on method definition
	data, err := c.StakeManagerABI.Pack("stakeFor", val, stakeAmount, feeAmount, acceptDelegation, signerPubkeyBytes)
	if err != nil {
		return err
	}

	auth, err := GenerateAuthObj(c.MainChainClient, stakeManagerAddress, data, signer, nil)
	if err != nil {
		return err
	}

	// stake for stake manager
	tx, err := stakeManagerInstance.StakeFor(
		auth,
		val,
		stakeAmount,
		feeAmount,
		acceptDelegation,
		signerPubkeyBytes,
	)

	if err != nil {
		return err
	}

	fmt.Println("Submitted stake sucessfully", "txHash", tx.Hash().String())
	return nil
}

// Topup - Add topup to user
func (c *ContractCaller) Topup(userAddr common.Address, feeAmount *big.Int, stakeManagerAddress common.Address, stakeManagerInstance *stakemanager.Stakemanager, signer types.Account) error {

	// pack data based on method definition
	data, err := c.StakeManagerABI.Pack("topUpForFee", userAddr, feeAmount)
	if err != nil {
		return err
	}

	auth, err := GenerateAuthObj(c.MainChainClient, stakeManagerAddress, data, signer, nil)
	if err != nil {
		return err
	}

	// stake for stake manager
	tx, err := stakeManagerInstance.TopUpForFee(
		auth,
		userAddr,
		feeAmount,
	)

	if err != nil {
		fmt.Println("error", err)
		return err
	}

	fmt.Println("Submitted topup sucessfully", "txHash", tx.Hash().String())
	return nil
}

// ClaimFee - submit heimdall proof and claim topup
func (c *ContractCaller) ClaimFee(fee *big.Int, index *big.Int, proof []byte, stakeManagerAddress common.Address, stakeManagerInstance *stakemanager.Stakemanager, signer types.Account) error {
	// pack data based on method definition
	data, err := c.StakeManagerABI.Pack("claimFee", fee, index, proof)
	if err != nil {
		return err
	}

	auth, err := GenerateAuthObj(c.MainChainClient, stakeManagerAddress, data, signer, nil)
	if err != nil {
		return err
	}

	tx, err := stakeManagerInstance.ClaimFee(auth, fee, index, proof)
	if err != nil {
		fmt.Println("error", err)
		return err
	}
	fmt.Println("Claimed topup sucessfully", "txHash", tx.Hash().String())
	return nil
}

func (c *ContractCaller) Restake(valId *big.Int, stakeAmount *big.Int, stakeRewards bool, stakeManagerAddress common.Address, stakeManagerInstance *stakemanager.Stakemanager, signer types.Account) error {
	// pack data based on method definition
	data, err := c.StakeManagerABI.Pack("restake", valId, stakeAmount, stakeRewards)
	if err != nil {
		return err
	}

	auth, err := GenerateAuthObj(c.MainChainClient, stakeManagerAddress, data, signer, nil)
	if err != nil {
		return err
	}

	// stake for stake manager
	tx, err := stakeManagerInstance.Restake(
		auth,
		valId,
		stakeAmount,
		stakeRewards,
	)

	if err != nil {
		return err
	}

	fmt.Println("Submitted restake sucessfully", "txHash", tx.Hash().String())
	return nil
}

func (c *ContractCaller) Unstake(valId *big.Int, stakeManagerAddress common.Address, stakeManagerInstance *stakemanager.Stakemanager, signer types.Account) error {
	// pack data based on method definition
	data, err := c.StakeManagerABI.Pack("unstake", valId)
	if err != nil {
		return err
	}

	auth, err := GenerateAuthObj(c.MainChainClient, stakeManagerAddress, data, signer, nil)
	if err != nil {
		return err
	}

	// stake for stake manager
	tx, err := stakeManagerInstance.Unstake(
		auth,
		valId,
	)

	if err != nil {
		return err
	}

	fmt.Println("Submitted unstake sucessfully", "txHash", tx.Hash().String())
	return nil
}

func (c *ContractCaller) UpdateSigner(valId *big.Int, newValidator types.Account, stakeManagerAddress common.Address, stakeManagerInstance *stakemanager.Stakemanager, signer types.Account) error {
	key := newValidator.PrivKey[2:]

	// create ecdsa private key
	ecdsaPrivateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return err
	}

	publicKey := ecdsaPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	newPubkeyBytes := crypto.FromECDSAPub(publicKeyECDSA)[1:]

	// pack data based on method definition
	data, err := c.StakeManagerABI.Pack("updateSigner", valId, newPubkeyBytes)
	if err != nil {
		return err
	}

	auth, err := GenerateAuthObj(c.MainChainClient, stakeManagerAddress, data, signer, nil)
	if err != nil {
		return err
	}

	// stake for stake manager
	tx, err := stakeManagerInstance.UpdateSigner(
		auth,
		valId,
		newPubkeyBytes,
	)

	if err != nil {
		return err
	}

	fmt.Println("Submitted update signer sucessfully", "txHash", tx.Hash().String())
	return nil
}

// ApproveTokens approves matic token for stake
func (c *ContractCaller) ApproveTokens(amount *big.Int, approveFor common.Address, tokenAddress common.Address, maticTokenInstance *erc20.Erc20, signer types.Account) error {
	data, err := c.MaticTokenABI.Pack("approve", approveFor, amount)
	if err != nil {
		return err
	}

	auth, err := GenerateAuthObj(c.MainChainClient, tokenAddress, data, signer, nil)
	if err != nil {
		return err
	}

	tx, err := maticTokenInstance.Approve(auth, approveFor, amount)
	if err != nil {
		return err
	}

	fmt.Println("Sent approve tx sucessfully", "txHash", tx.Hash().String())
	return nil
}

// TransferToken transfers tokens from sender to receiver
func (c *ContractCaller) TransferToken(amount *big.Int, receiver common.Address, tokenAddress common.Address, maticTokenInstance *erc20.Erc20, signer types.Account, parent bool) error {
	data, err := c.MaticTokenABI.Pack("transfer", receiver, amount)
	if err != nil {
		return err
	}
	var client *ethclient.Client
	if parent == true {
		client = c.MainChainClient
	} else {
		client = c.MaticChainClient
	}
	auth, err := GenerateAuthObj(client, tokenAddress, data, signer, nil)
	if err != nil {
		return err
	}

	tx, err := maticTokenInstance.Transfer(auth, receiver, amount)
	if err != nil {
		return err
	}

	fmt.Println("Sent transfer token tx sucessfully", "txHash", tx.Hash().String())
	return nil
}

// ProcessExits Deposits tokens to matic chain
func (c *ContractCaller) ProcessExits(tokenAddress common.Address, withdrawManagerAddress common.Address, withdrawManagerInstance *withdrawmanager.Withdrawmanager, signer types.Account) error {
	data, err := c.WithdrawManagerABI.Pack("processExits", tokenAddress)
	if err != nil {
		fmt.Println("error packing", "error", err)
		return err
	}

	auth, err := GenerateAuthObj(c.MainChainClient, withdrawManagerAddress, data, signer, nil)
	if err != nil {
		fmt.Println("error generating auth", "error", err)
		return err
	}

	tx, err := withdrawManagerInstance.ProcessExits(auth, tokenAddress)
	if err != nil {
		return err
	}

	fmt.Println("ProcessExit tx sucessfully", "txHash", tx.Hash().String())
	return nil
}

func (c *ContractCaller) BurnERC20Tokens(amount *big.Int, tokenAddress common.Address, childerc20Instance *childerc20.Childerc20, signer types.Account) (string, error) {
	data, err := c.ChildERC20ABI.Pack("withdraw", amount)
	if err != nil {
		return "", err
	}

	auth, err := GenerateAuthObj(c.MaticChainClient, tokenAddress, data, signer, nil)
	if err != nil {
		return "", err
	}

	tx, err := childerc20Instance.Withdraw(auth, amount)
	if err != nil {
		return "", err
	}

	fmt.Println("Withdraw initiated sucessfully", "txHash", tx.Hash().String())
	return tx.Hash().String(), nil
}

func (c *ContractCaller) BurnMaticTokens(amount *big.Int, tokenAddress common.Address, mrc20Instance *mrc20.Mrc20, signer types.Account) (string, error) {
	data, err := c.MRC20ABI.Pack("withdraw", amount)
	if err != nil {
		return "", err
	}

	auth, err := GenerateAuthObj(c.MaticChainClient, tokenAddress, data, signer, amount)
	if err != nil {
		return "", err
	}

	tx, err := mrc20Instance.Withdraw(auth, amount)
	if err != nil {
		return "", err
	}

	fmt.Println("Withdraw initiated sucessfully", "txHash", tx.Hash().String())
	return tx.Hash().String(), nil
}

func (c *ContractCaller) StartExitWithBurntTokens(payload []byte, address common.Address, erc20PredicateInstance *erc20predicate.Erc20predicate, signer types.Account) (string, error) {
	data, err := c.ERC20PredicateABI.Pack("startExitWithBurntTokens", payload)
	if err != nil {
		return "", err
	}

	auth, err := GenerateAuthObj(c.MainChainClient, address, data, signer, nil)
	if err != nil {
		return "", err
	}

	tx, err := erc20PredicateInstance.StartExitWithBurntTokens(auth, payload)
	if err != nil {
		return "", err
	}

	fmt.Println("Confirm withdraw submitted sucessfully", "txHash", tx.Hash().String())
	return tx.Hash().String(), nil
}

// DepositERC20ForUser Deposits tokens to matic chain
func (c *ContractCaller) DepositERC20ForUser(amount *big.Int, user common.Address, tokenAddress common.Address, depositManagerInstance *depositmanager.Depositmanager, signer types.Account) error {
	data, err := c.DepositManagerABI.Pack("depositERC20ForUser", tokenAddress, user, amount)
	if err != nil {
		fmt.Println("error packing", "error", err)
		return err
	}

	auth, err := GenerateAuthObj(c.MainChainClient, tokenAddress, data, signer, nil)
	if err != nil {
		fmt.Println("error generating auth", "error", err)
		return err
	}

	tx, err := depositManagerInstance.DepositERC20ForUser(auth, tokenAddress, user, amount)
	if err != nil {
		fmt.Println("error deposition", "error", err)
		return err
	}

	fmt.Println("Deposit token tx sucessfully", "txHash", tx.Hash().String())
	return nil
}

func (c *ContractCaller) UpdateSignerUpdateLimit(limit *big.Int, stakeManagerAddress common.Address, stakeManagerInstance *stakemanager.Stakemanager, signer types.Account) error {
	data, err := c.StakeManagerABI.Pack("updateSignerUpdateLimit", limit)
	if err != nil {
		return err
	}

	auth, err := GenerateAuthObj(c.MainChainClient, stakeManagerAddress, data, signer, nil)
	if err != nil {
		return err
	}

	tx, err := stakeManagerInstance.UpdateSignerUpdateLimit(auth, limit)
	if err != nil {
		return err
	}

	fmt.Println("Signer limit updated sucessfully", "txHash", tx.Hash().String())
	return nil
}

// BalanceOf returns balance of given user
func (c *ContractCaller) BalanceOf(user common.Address, tokenAddress common.Address, parent bool) (*big.Int, error) {
	erc20TokenInstance, _ := c.GetERC20TokenInstance(tokenAddress, parent)

	balance, err := erc20TokenInstance.BalanceOf(nil, user)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

// BalanceOf returns balance of given user
func (c *ContractCaller) BalanceOfChildERC20(user common.Address, childERC20Instance *childerc20.Childerc20) (*big.Int, error) {

	balance, err := childERC20Instance.BalanceOf(nil, user)
	if err != nil {
		return nil, err
	}

	fmt.Println("balance", balance.String())
	return balance, nil
}

// FindHeaderBlockNumber fetches current header block
func (c *ContractCaller) FindHeaderBlockNumber(childBlockNumber *big.Int, rootChainInstance *rootchain.Rootchain) *big.Int {
	var ans uint64
	start := big.NewInt(1).Uint64()
	CHECKPOINT_ID_INTERVAL := big.NewInt(10000)

	currentHeaderBlock, _ := c.CurrentHeaderBlock(rootChainInstance)
	end := currentHeaderBlock.Div(currentHeaderBlock, CHECKPOINT_ID_INTERVAL).Uint64()
	// binary search on all the checkpoints to find the checkpoint that contains the childBlockNumber
	for start <= end {
		if start == end {
			ans = start
			break
		}
		mid := (start + end) / uint64(2)
		bigMid := big.NewInt(0).SetUint64(mid)

		headerBlock, _ := c.HeaderBlock(bigMid.Mul(bigMid, CHECKPOINT_ID_INTERVAL), rootChainInstance)
		headerStart := headerBlock.StartBlock
		headerEnd := headerBlock.EndBlock

		if headerStart <= childBlockNumber.Uint64() && childBlockNumber.Uint64() < headerEnd {
			ans = mid
			break
		} else if headerStart > childBlockNumber.Uint64() {
			end = mid - uint64(1)
		} else if headerEnd < childBlockNumber.Uint64() {
			start = mid + uint64(1)
		}
	}
	return big.NewInt(0).SetUint64(ans * CHECKPOINT_ID_INTERVAL.Uint64())
}

// CurrentHeaderBlock fetches current header block
func (c *ContractCaller) HeaderBlock(blockNumber *big.Int, rootChainInstance *rootchain.Rootchain) (hmTypes.Checkpoint, error) {
	headerBlock, err := rootChainInstance.HeaderBlocks(nil, blockNumber)
	if err != nil {
		fmt.Println("Could not fetch current header block from rootchain contract", "Error", err)
		return hmTypes.Checkpoint{}, err
	}
	block := hmTypes.CreateBlock(headerBlock.Start.Uint64(), headerBlock.End.Uint64(), hmTypes.ZeroHeimdallHash, hmTypes.BytesToHeimdallAddress(headerBlock.Proposer.Bytes()), "", uint64(0))
	return block, nil
}

// CurrentHeaderBlock fetches current header block
func (c *ContractCaller) UnmutatedHeaderBlock(blockNumber *big.Int, rootChainInstance *rootchain.Rootchain) (hmTypes.Checkpoint, error) {
	headerBlock, err := rootChainInstance.HeaderBlocks(nil, blockNumber)
	if err != nil {
		fmt.Println("Could not fetch current header block from rootchain contract", "Error", err)
		return hmTypes.Checkpoint{}, err
	}
	block := hmTypes.CreateBlock(headerBlock.Start.Uint64(), headerBlock.End.Uint64(), headerBlock.Root, hmTypes.BytesToHeimdallAddress(headerBlock.Proposer.Bytes()), "", uint64(0))
	return block, nil
}

// CurrentHeaderBlock fetches current header block
func (c *ContractCaller) CurrentHeaderBlock(rootChainInstance *rootchain.Rootchain) (*big.Int, error) {
	currentHeaderBlock, err := rootChainInstance.CurrentHeaderBlock(nil)
	if err != nil {
		fmt.Println("Could not fetch current header block from rootchain contract", "Error", err)
		return nil, err
	}
	return currentHeaderBlock, nil
}

// GetLastChildBlock fetches current header block
func (c *ContractCaller) GetLastChildBlock(rootChainInstance *rootchain.Rootchain) (uint64, error) {
	getLastChildBlock, err := rootChainInstance.GetLastChildBlock(nil)
	if err != nil {
		fmt.Println("Could not fetch current header block from rootchain contract", "Error", err)
		return 0, err
	}
	return getLastChildBlock.Uint64(), nil
}

// WaitForCheckpoints - Wait for checkpoints to go through on contract
func (c *ContractCaller) WaitForCheckpoints(numOfCheckPoints uint64, rootChainInstance *rootchain.Rootchain) {
	startingCheckpoint, _ := c.CurrentHeaderBlock(rootChainInstance)
	interval := 60 * time.Second

	ticker := time.NewTicker(interval)
	// stop ticker when everything done
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			currentCheckpoint, _ := c.CurrentHeaderBlock(rootChainInstance)
			if (currentCheckpoint.Uint64() - startingCheckpoint.Uint64()) == numOfCheckPoints*10000 {
				return
			}
		}
	}
}

// GetHalfExitPeriod fetches current header block
func (c *ContractCaller) GetHalfExitPeriod(withdrawManagerInstance *withdrawmanager.Withdrawmanager) (*big.Int, error) {
	halfExitPeriod, err := withdrawManagerInstance.HALFEXITPERIOD(nil)
	if err != nil {
		fmt.Println("Could not fetch HALFEXITPERIOD From contract", "Error", err)
		return nil, err
	}
	return halfExitPeriod, nil
}

// TransferEth transfers eth from sender to receiver
func (c *ContractCaller) TransferEth(amount *big.Int, receiver common.Address, signer types.Account, client *ethclient.Client) error {
	auth, err := GenerateAuthObj(client, receiver, nil, signer, nil)
	if err != nil {
		return err
	}

	// Create the transaction, sign it and schedule it for execution
	rawTx := ethTypes.NewTransaction(auth.Nonce.Uint64(), receiver, amount, auth.GasLimit, auth.GasPrice, nil)

	// signer
	signedTx, err := auth.Signer(ethTypes.HomesteadSigner{}, auth.From, rawTx)
	if err != nil {
		fmt.Println("Error signing the transaction", "error", err)
		return err
	}

	fmt.Println("Sending transaction to transfer eth", "txHash", signedTx.Hash().Hex())

	// broadcast transaction
	if err := client.SendTransaction(context.Background(), signedTx); err != nil {
		fmt.Println("Error while broadcasting the transaction", "error", err)
		return err
	}

	fmt.Println("Sent Eth transfer tx sucessfully", "txHash", signedTx.Hash().String())
	return nil
}

// GetValidatorInfo get validator info
func (c *ContractCaller) GetValidatorInfo(valID *big.Int, stakingInfoInstance *stakinginfo.Stakinginfo) (validator hmTypes.Validator, err error) {
	// amount, startEpoch, endEpoch, signer, status, err := c.StakingInfoInstance.GetStakerDetails(nil, big.NewInt(int64(valID)))
	stakerDetails, err := stakingInfoInstance.GetStakerDetails(nil, valID)
	fmt.Println("val id", stakerDetails)
	fmt.Println("val id", hmTypes.NewValidatorID(valID.Uint64()))
	if err != nil {
		fmt.Println("Error fetching validator information from stake manager", "error", err, "validatorId", valID, "status", stakerDetails.Status)
		return
	}

	newAmount, err := hmHelper.GetPowerFromAmount(stakerDetails.Amount)
	if err != nil {
		return
	}

	// newAmount
	validator = hmTypes.Validator{
		ID:          hmTypes.NewValidatorID(valID.Uint64()),
		VotingPower: newAmount.Int64(),
		StartEpoch:  stakerDetails.ActivationEpoch.Uint64(),
		EndEpoch:    stakerDetails.DeactivationEpoch.Uint64(),
		Signer:      hmTypes.BytesToHeimdallAddress(stakerDetails.Signer.Bytes()),
	}

	return validator, nil
}

// GetValidators get validator info
func (c *ContractCaller) GetValidators(valID *big.Int, stakeManagerInstance *stakemanager.Stakemanager) (validator hmTypes.Validator, err error) {
	// amount, startEpoch, endEpoch, signer, status, err := c.StakingInfoInstance.GetStakerDetails(nil, big.NewInt(int64(valID)))
	stakerDetails, err := stakeManagerInstance.Validators(nil, valID)
	if err != nil {
		fmt.Println("Error fetching validator information from stake manager", "error", err, "validatorId", valID, "status", stakerDetails.Status)
		return
	}
	newAmount, err := hmHelper.GetPowerFromAmount(stakerDetails.Amount)
	if err != nil {
		return
	}

	// newAmount
	validator = hmTypes.Validator{
		ID:          hmTypes.NewValidatorID(valID.Uint64()),
		VotingPower: newAmount.Int64(),
		StartEpoch:  stakerDetails.ActivationEpoch.Uint64(),
		EndEpoch:    stakerDetails.DeactivationEpoch.Uint64(),
		Signer:      hmTypes.BytesToHeimdallAddress(stakerDetails.Signer.Bytes()),
	}

	return validator, nil
}

func (c *ContractCaller) GetValidatorId(val common.Address, stakeManagerInstance *stakemanager.Stakemanager) (*big.Int, error) {
	validator, err := stakeManagerInstance.GetValidatorId(nil, val)

	if err != nil {
		return nil, err
	}

	return validator, nil
}

func (c *ContractCaller) SignerToValidator(val common.Address, stakeManagerInstance *stakemanager.Stakemanager) (*big.Int, error) {
	validator, err := stakeManagerInstance.SignerToValidator(nil, val)

	if err != nil {
		return nil, err
	}

	return validator, nil
}

func (c *ContractCaller) IsValidator(id *big.Int, stakeManagerInstance *stakemanager.Stakemanager) (bool, error) {
	isValidator, err := stakeManagerInstance.IsValidator(nil, id)

	if err != nil {
		return false, err
	}

	return isValidator, nil
}

func (c *ContractCaller) TotalStakedFor(user common.Address, stakeManagerInstance *stakemanager.Stakemanager) (*big.Int, error) {
	stakeAmount, err := stakeManagerInstance.TotalStakedFor(nil, user)

	if err != nil {
		return nil, err
	}

	return stakeAmount, nil
}

func (c *ContractCaller) ValidatorStake(val *big.Int, stakeManagerInstance *stakemanager.Stakemanager) (*big.Int, error) {
	stakeAmount, err := stakeManagerInstance.ValidatorStake(nil, val)

	if err != nil {
		return nil, err
	}

	return stakeAmount, nil
}

// TotalHeimdallFee - total heimdall fee
func (c *ContractCaller) TotalHeimdallFee(stakeManagerInstance *stakemanager.Stakemanager) (*big.Int, error) {
	totalHeimdallFee, err := stakeManagerInstance.TotalHeimdallFee(nil)
	if err != nil {
		return nil, err
	}
	return totalHeimdallFee, nil
}

func (c *ContractCaller) GetTransaction(txHash string) (tx *ethTypes.Transaction) {
	tx, _, err := c.MaticChainClient.TransactionByHash(context.Background(), hmTypes.HexToHeimdallHash(txHash).EthHash())
	if err != nil {
		return
	}
	return tx
}

func (c *ContractCaller) GetTransactionReceipt(txHash string, parent bool) *ethTypes.Receipt {
	var receipt *ethTypes.Receipt
	var err error
	if parent == true {
		receipt, err = c.MainChainClient.TransactionReceipt(context.Background(), hmTypes.HexToHeimdallHash(txHash).EthHash())
		if err != nil {
			fmt.Println("Error", err)
			return nil
		}
	} else {
		receipt, err = c.MaticChainClient.TransactionReceipt(context.Background(), hmTypes.HexToHeimdallHash(txHash).EthHash())
		if err != nil {
			fmt.Println("matic Error", err)

			return nil
		}
	}

	return receipt
}

func (c *ContractCaller) GetBlockByNumber(number *big.Int, parent bool) *ethTypes.Block {
	var block *ethTypes.Block
	var err error
	if parent == true {
		block, err = c.MainChainClient.BlockByNumber(context.Background(), number)
		if err != nil {
			return nil
		}
	} else {
		block, err = c.MaticChainClient.BlockByNumber(context.Background(), number)
		if err != nil {
			return nil
		}
	}

	return block
}

func (c *ContractCaller) GetHeaderByNumber(number *big.Int) *ethTypes.Header {
	header, err := c.MainChainClient.HeaderByNumber(context.Background(), number)
	if err != nil {
		return nil
	}
	return header
}

func (c *ContractCaller) GetRootHash(start uint64, end uint64) string {
	rootHash, err := c.MaticChainClient.GetRootHash(context.Background(), start, end)
	if err != nil {
		return ""
	}
	return rootHash
}

// GenerateAuthObj - creates tx for user
func GenerateAuthObj(client *ethclient.Client, address common.Address, data []byte, signer types.Account, value *big.Int) (auth *bind.TransactOpts, err error) {
	// generate call msg
	callMsg := ethereum.CallMsg{
		To:   &address,
		Data: data,
	}

	// get priv key
	key := signer.PrivKey[2:]
	addr := signer.Address

	// create ecdsa private key
	ecdsaPrivateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return
	}

	// from address
	fromAddress := common.HexToAddress(addr)

	// fetch gas price
	gasprice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	// fetch nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return
	}

	// fetch gas limit
	callMsg.From = fromAddress
	// gasLimit, err := client.EstimateGas(context.Background(), callMsg)
	gasLimit := 800000
	// create auth
	auth = bind.NewKeyedTransactor(ecdsaPrivateKey)
	auth.GasPrice = gasprice
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(gasLimit) // uint64(gasLimit)
	auth.Value = value
	return
}

// GenerateSimpleAuthObj - creates tx for user
func GenerateSimpleAuthObj(client *ethclient.Client, address common.Address, data []byte, signer types.Account) (auth *bind.TransactOpts, err error) {
	// generate call msg
	callMsg := ethereum.CallMsg{
		To:   &address,
		Data: data,
	}

	// get priv key
	key := signer.PrivKey[2:]
	addr := signer.Address

	// create ecdsa private key
	ecdsaPrivateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return
	}

	// from address
	fromAddress := common.HexToAddress(addr)

	// fetch gas price
	gasprice, _ := big.NewInt(0).SetString("1000000000", 10)

	// fetch nonce
	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	// 	return
	// }
	nonce := 0

	// fetch gas limit
	callMsg.From = fromAddress
	// gasLimit, err := client.EstimateGas(context.Background(), callMsg)
	gasLimit := 97083

	// create auth
	auth = bind.NewKeyedTransactor(ecdsaPrivateKey)
	auth.GasPrice = gasprice
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(gasLimit) // uint64(gasLimit)

	return
}

//
// private abi methods
//

func getABI(data string) (abi.ABI, error) {
	return abi.JSON(strings.NewReader(data))
}

func nextPowerOfTwo(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	// http://graphics.stanford.edu/~seander/bithacks.html#RoundUpPowerOf2
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	n++
	return n
}

// spins go-routines to fetch batch elements to allow creation of large merkle trees
func fetchBatchElements(rpcClient *rpc.Client, elements []rpc.BatchElem) (err error) {
	var batchLength = int(len(elements))
	// group
	var g errgroup.Group

	for i := 0; i < len(elements); i += batchLength {
		var newBatch []rpc.BatchElem
		if len(elements) < i+batchLength {
			newBatch = elements[i:]
		} else {
			newBatch = elements[i : i+batchLength]
		}

		// spawn go-routine
		g.Go(func() error {
			// Batch call
			err := rpcClient.BatchCall(newBatch)
			return err
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	// common.CheckpointLogger.Info("Fetched all headers", "len", len(elements))
	return nil
}

func appendBytes32(data ...[]byte) []byte {
	var result []byte
	for _, v := range data {
		paddedV, err := convertTo32(v)
		if err == nil {
			result = append(result, paddedV[:]...)
		}
	}
	return result
}

func convertTo32(input []byte) (output [32]byte, err error) {
	l := len(input)
	if l > 32 || l == 0 {
		return
	}
	copy(output[32-l:], input[:])
	return
}

func convert(input []([32]byte)) [][]byte {
	var output [][]byte
	for _, in := range input {
		newInput := make([]byte, len(in[:]))
		copy(newInput, in[:])
		output = append(output, newInput)

	}
	return output
}
