module github.com/maticnetwork/reddit-bakeoff-poc

go 1.14

require (
	github.com/cbergoon/merkletree v0.2.0
	github.com/cosmos/cosmos-sdk v0.37.4
	github.com/ethereum/go-ethereum v1.9.15
	github.com/maticnetwork/bor v0.1.7-0.20200507151553-e03cd94ed12b
	github.com/maticnetwork/heimdall v0.2.0
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/tendermint/tendermint v0.33.3
	golang.org/x/crypto v0.0.0-20200429183012-4b2356b1ed79
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
)

replace github.com/tendermint/tendermint => github.com/maticnetwork/tendermint v0.26.0-dev0.0.20200429080413-edc079e7d4c9

replace github.com/cosmos/cosmos-sdk => github.com/maticnetwork/cosmos-sdk v0.37.5-0.20200503092858-55131f25dd9d
