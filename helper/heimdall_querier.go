package helper

import (
	"encoding/json"
	"fmt"
	"log"

	cliContext "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/maticnetwork/heimdall/app"
	authTypes "github.com/maticnetwork/heimdall/auth/types"
	"github.com/maticnetwork/heimdall/bridge/setu/util"
	chTypes "github.com/maticnetwork/heimdall/checkpoint/types"
	"github.com/maticnetwork/heimdall/helper"
	"github.com/maticnetwork/heimdall/types"
	hmTypes "github.com/maticnetwork/heimdall/types"
	"github.com/maticnetwork/heimdall/types/rest"
)

// HeimdallQuerier heimdall endpoint querier
type HeimdallQuerier struct {
	BasePath string
	CliCtx   cliContext.CLIContext
}

// NewHeimdallQuerier - initialze heimdall querier
func NewHeimdallQuerier(url string) (HeimdallQuerier HeimdallQuerier, err error) {
	HeimdallQuerier.BasePath = url
	HeimdallQuerier.CliCtx = cliContext.NewCLIContext().WithCodec(app.MakeCodec())
	return
}

// GetValidatorInfo returns validatorInfo
func (h *HeimdallQuerier) GetValidatorInfo(valID int64) (val *hmTypes.Validator, err error) {
	response, err := helper.FetchFromAPI(h.CliCtx, h.BasePath+fmt.Sprintf(util.ValidatorURL, valID))
	if err != nil {
		log.Println("Error Fetching validatorInfo from HeimdallServer ", "error", err)
		return
	}
	if err = json.Unmarshal(response.Result, &val); err != nil {
		log.Println("Error unmarshalling validatorInfo received from Heimdall Server", "error", err)
		return
	}
	return
}

// GetLGetLatestCheckpoint gets the latest checkpoint object
func (h *HeimdallQuerier) GetLatestCheckpoint() (c hmTypes.Checkpoint, err error) {
	resp, err := h.Query(util.LatestCheckpointURL, nil)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(resp.Result, &c)
	return
}

// GetLGetLatestCheckpoint gets the latest checkpoint object
func (h *HeimdallQuerier) GetCheckpointByID(id int64) (c hmTypes.Checkpoint, err error) {
	resp, err := h.Query(fmt.Sprintf("/checkpoints/%v", id), nil)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(resp.Result, &c)
	return
}

// GetLGetLatestCheckpoint gets the latest checkpoint object
func (h *HeimdallQuerier) GetCheckpointCount() (i int64, err error) {
	resp, err := h.Query("/checkpoints/count", nil)
	if err != nil {
		return i, err
	}
	c := struct {
		R int64 `json:"result"`
	}{}
	err = json.Unmarshal(resp.Result, &c)
	i = c.R
	return
}

// GetAccount returns heimdall auth account
func (h *HeimdallQuerier) GetAccount(address types.HeimdallAddress) (account authTypes.Account, err error) {
	// call account rest api
	response, err := helper.FetchFromAPI(h.CliCtx, h.BasePath+fmt.Sprintf(util.AccountDetailsURL, address))
	if err != nil {
		return
	}
	if err = h.CliCtx.Codec.UnmarshalJSON(response.Result, &account); err != nil {
		log.Println("Error unmarshalling account details", "error", err)
		return
	}
	return
}

// GetDividendAccount - fetch dividend account of user
func (h *HeimdallQuerier) GetDividendAccount(address hmTypes.HeimdallAddress) (dividendAccount hmTypes.DividendAccount, err error) {
	const DividendAccountURL = "/topup/dividend-account/%v"
	// call account rest api
	response, err := helper.FetchFromAPI(h.CliCtx, h.BasePath+fmt.Sprintf(DividendAccountURL, address))
	if err != nil {
		return
	}
	if err = h.CliCtx.Codec.UnmarshalJSON(response.Result, &dividendAccount); err != nil {
		log.Println("Error unmarshalling account details", "error", err)
		return
	}
	return
}

// GetAccountProof - Account proof of topup withdraw amount
func (h *HeimdallQuerier) GetAccountProof(address types.HeimdallAddress) (proof hmTypes.DividendAccountProof, err error) {
	const AccountProofURL = "/topup/account-proof/%v"
	// call account rest api
	response, err := helper.FetchFromAPI(h.CliCtx, h.BasePath+fmt.Sprintf(AccountProofURL, address))
	if err != nil {
		return
	}

	if err = json.Unmarshal(response.Result, &proof); err != nil {
		log.Println("Error unmarshalling account proof", "error", err)
		return
	}

	return
}

func (h *HeimdallQuerier) GetSpanSeed() (spanSeed string, err error) {
	resp, err := h.Query(util.NextSpanSeedURL, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Result, &spanSeed)
	if err != nil {
		return
	}
	return
}

func (h *HeimdallQuerier) GetLatestSpan() (s hmTypes.Span, err error) {
	resp, err := h.Query(util.LatestSpanURL, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Result, &s)
	return
}

func (h *HeimdallQuerier) PrepareNextSpan(chainID, proposer, spanID, startBlock string) (s hmTypes.Span, err error) {
	resp, err := h.Query(util.NextSpanInfoURL, []string{"chain_id=" + chainID, "proposer=" + proposer, "span_id=" + spanID, "start_block=" + startBlock})
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Result, &s)
	return
}

func (h *HeimdallQuerier) Query(url string, params []string) (resp rest.ResponseWithHeight, err error) {
	if l := len(params); l > 0 {
		url += "?"
		for i := 0; i < l; i++ {
			url += params[i]
			if i != l-1 {
				url += "&"
			}
		}
	}
	resp, err = helper.FetchFromAPI(h.CliCtx, h.BasePath+url)
	if err != nil {
		log.Println("Error Fetching validatorInfo from HeimdallServer ", "error", err)
	}
	return
}

// GetCheckpointParams returns the paramaters/config for a checkpoint being
// submitted
func (h *HeimdallQuerier) GetCheckpointParams() (chPar chTypes.Params, err error) {
	resp, err := h.Query(util.CheckpointParamsURL, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp.Result, &chPar)
	return
}
