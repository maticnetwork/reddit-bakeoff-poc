package types

// ContractAddresses of all contracts
type ContractAddresses struct {
	RootChain  RootChain  `json:"root"`
	ChildChain ChildChain `json:"child"`
}

// RootChain contract addresses
type RootChain struct {
	Predicates           Predicates
	Registry             string
	RootChain            string
	Governance           string
	GovernanceProxy      string
	RootChainProxy       string
	DepositManager       string
	DepositManagerProxy  string
	WithdrawManager      string
	WithdrawManagerProxy string
	StakeManager         string
	StakeManagerProxy    string
	ValidatorShare       string
	SlashingManager      string
	StakingInfo          string
	ExitNFT              string
	StateSender          string
	Tokens               Tokens
}

// Child contract addresses
type ChildChain struct {
	ChildChain string
	Tokens     Tokens
}

// Tokens addresses used by selected network
type Tokens struct {
	MaticToken string
	TestToken  string
	RootERC721 string
	MaticWeth  string
}

// Predicates Addresses
type Predicates struct {
	ERC20Predicate           string
	ERC721Predicate          string
	MarketplacePredicate     string
	TransferWithSigPredicate string
}
