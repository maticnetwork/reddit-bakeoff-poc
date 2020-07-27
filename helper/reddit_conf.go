package helper

import (
	"github.com/maticnetwork/reddit-bakeoff-poc/types"
)

// RedditConf config for selected network
type RedditConf struct {
	RedditContracts types.RedditContracts
	RedditConfig    types.RedditConfig
	RedditBenchmark types.RedditBenchmark
	RedditAccounts  []types.Account
	RedditCaller    RedditCaller
}

var redditConf RedditConf

// InitAndSetRedditConf - initializes RedditConf for given network
func InitAndSetRedditConf(network *string) {
	redditConf = NewRedditConf(network)
}

// NewRedditConf - create new RedditConf
func NewRedditConf(network *string) (redditConf RedditConf) {
	// base path for all configuration // path is relative to main_test.go
	c := GetConf()
	basePath := "config/" + *network
	redditConf.RedditAccounts = ParseRedditAccounts(basePath)
	redditConf.RedditConfig = ParseRedditConfig(basePath)
	redditConf.RedditBenchmark = ParseRedditBenchmark(basePath)
	redditConf.RedditContracts = ParseRedditContracts(basePath)
	redditConf.RedditCaller, _ = NewRedditCaller(c.NetworkConfig.EthRPCUrl, c.NetworkConfig.BorRPCUrl)
	return
}

// GetRedditConf - returns RedditConf
func GetRedditConf() RedditConf {
	return redditConf
}
