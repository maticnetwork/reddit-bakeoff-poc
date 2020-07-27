package helper

import (
	"log"

	"github.com/maticnetwork/bor/rpc"

	"github.com/maticnetwork/bor/ethclient"
)

//
// Get main/matic clients
//

// GetMainClient returns main chain's eth client
func GetMainClient(ethRPCUrl string) (*ethclient.Client, *rpc.Client) {
	mainRPCClient, err := rpc.Dial(ethRPCUrl)
	if err != nil {
		log.Fatalln("Unable to dial via ethClient", "URL=", ethRPCUrl, "Error", err)
	}

	mainChainClient := ethclient.NewClient(mainRPCClient)
	return mainChainClient, mainRPCClient
}

// GetMaticClient returns matic's eth client
func GetMaticClient(borRPCUrl string) (*ethclient.Client, *rpc.Client) {
	maticRPCClient, err := rpc.Dial(borRPCUrl)
	if err != nil {
		log.Fatal(err)
	}

	maticClient := ethclient.NewClient(maticRPCClient)
	return maticClient, maticRPCClient
}
