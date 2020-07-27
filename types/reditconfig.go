package types

// RedditConfig initialization
type RedditConfig struct {
	InitialSupply        uint64
	InitialKarma         uint64
	SubscriptionPrice    uint64
	Duration             uint64
	RenewBefore          uint64
	SubredditName        string
	SubredditTokenSymbol string
	SubredditTokenName   string
}
