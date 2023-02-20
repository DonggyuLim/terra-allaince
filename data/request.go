package data

import (
	"fmt"
	"strconv"

	"github.com/DonggyuLim/Alliance-Rank/request"
	"github.com/DonggyuLim/Alliance-Rank/utils"
	"github.com/imroc/req/v3"
)

func GetRewards(chainCode, height int, delegator, validator, denom string) ([]request.Reward, error) {

	client := req.R().
		SetHeader("x-cosmos-block-height", fmt.Sprintf("%v", height)).SetHeader("Content-Type", "application/json")
	var req request.RewardRequest
	endpoint := fmt.Sprintf("%s/terra/alliances/rewards/%s/%s/%s",
		GetEndopoint(chainCode),
		delegator,
		validator,
		denom,
	)
	//{delegator_addr}/{validator_addr}/{denom}
	// endpoint := fmt.Sprintf("%s/terra/alliances/rewards/%s/{validator_addr}/{denom}", chain, el.deligator, validator, denom)
	_, err := client.SetSuccessResult(&req).Get(endpoint)

	return req.Rewards, err
}

func GetAddress(chainCode int, address string) string {
	switch chainCode {
	case 0:
		return utils.MakeAddress2(address, "atreides")
	case 1:
		return utils.MakeAddress2(address, "harkonnen")
	case 2:
		return utils.MakeAddress2(address, "corrino")
	case 3:
		return utils.MakeAddress2(address, "ordos")
	}
	return ""
}

func GetDelegations(height, chainCode int) (request.DelegationRequest, error) {

	value := fmt.Sprintf("%v", height)
	client := req.R().
		SetHeader("x-cosmos-block-height", value).SetHeader("Content-Type", "application/json")
	endpoint := fmt.Sprintf("%s/terra/alliances/delegations",
		GetEndopoint(chainCode),
		// GetAddress(chainCode, address),
	)

	var req request.DelegationRequest
	_, err := client.SetSuccessResult(&req).Get(endpoint)

	return req, err
}

func GetLastBlock(chainCode int) int {
	client := req.R()

	endpoint := fmt.Sprintf("%s/cosmos/base/tendermint/v1beta1/blocks/latest",
		GetEndopoint(chainCode),
	)
	var lastBlock request.LastBlock
	_, err := client.SetSuccessResult(&lastBlock).Get(endpoint)
	utils.PanicError(err)
	latestHeight, err := strconv.Atoi(lastBlock.Block.Header.Height)
	utils.PanicError(err)
	return latestHeight

}
