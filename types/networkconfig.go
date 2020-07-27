package types

// NetworkConfig for tests
type NetworkConfig struct {
	EthRPCUrl               string `json:"eth_rpc_url"`
	BorRPCUrl               string `json:"bor_rpc_url"`
	HeimdallChainID         string `json:"heimdall-chain-id"`
	TendermintRPCUrl        string `json:"tendermint_rpc_url"`
	HeimdallBaseURL         string `json:"heimdallBaseUrl"`
	BorChainID              string `json:"bor-chain-id"`
	SpanDuration            int    `json:"spanDuration"`
	UpdateSignerUpdateLimit string `json:"update-signer-update-limit"`
}
