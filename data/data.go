package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"sync"

	"github.com/DonggyuLim/Alliance-Rank/account"
	"github.com/DonggyuLim/Alliance-Rank/db"
	"github.com/DonggyuLim/Alliance-Rank/utils"
)

const (
	sCOR = "ibc/D7AA592A1C1C00FE7C9E15F4BB7ADB4B779627DD3FBB3C877CD4DB27F56E35B4"
	sORD = "ibc/3FA98D26F2D6CCB58D8E4D1B332C6EB8EE4AC7E3F0AD5B5B05201155CEB1AD1D"
	uatr = "uatr"
	uhar = "uhar"
	ucor = "ucor"
	uord = "uord"
)
const (
	ATREIDES = iota
	Harkonnen
	CORRINO
	ORDOS
)

func GetEndopoint(a int) string {
	switch a {
	case 0:
		// return "https://atreides.terra.dev:1317"
		return "http://localhost:1317"
	case 1:
		return "http://localhost:2317"
	case 2:
		return "http://localhost:3317"
	case 3:
		return "http://localhost:4317"
	}
	return ""
}

func Main(wg *sync.WaitGroup) {
	defer wg.Done()
	height := 37274

	for {
		latestBlockHeight := []int{GetLastBlock(ATREIDES), GetLastBlock(CORRINO), GetLastBlock(Harkonnen), GetLastBlock(ORDOS)}
		sort.Ints(latestBlockHeight)
		if height > latestBlockHeight[0] {
			height = latestBlockHeight[len(latestBlockHeight)-1]
		}

		MakeData(height, ATREIDES)
		MakeData(height, Harkonnen)
		MakeData(height, CORRINO)
		MakeData(height, ORDOS)
		height += 1
	}
}

func MakeData(height, chainCode int) {

	latestBlockHeight := GetLastBlock(chainCode)

	if height > latestBlockHeight {
		height = latestBlockHeight
		fmt.Println(height)
	} else {
		fmt.Println(height)
	}

	delegations, err := GetDelegations(height, chainCode)
	if delegations.Deligations == nil || err != nil {
		return
	}
	for _, el := range delegations.Deligations {

		resReward, err := GetRewards(
			chainCode,
			height,
			el.Delegation.DelegatorAddress,
			el.Delegation.ValidatorAddress,
			el.Delegation.Denom,
		)
		if err != nil || len(resReward) == 0 {
			continue
		}
		reward := account.Reward{
			LastHeight: height,
		}

		for _, re := range resReward {
			switch re.Denom {
			case sCOR:
				reward.SCOR = utils.DecimalAddString(reward.SCOR, re.Amount)

			case sORD:
				reward.SORD = utils.DecimalAddString(reward.SORD, re.Amount)
			case uatr:
				reward.UAtr = utils.DecimalAddString(reward.UAtr, re.Amount)
			case uhar:
				reward.UHar = utils.DecimalAddString(reward.UHar, re.Amount)
			case ucor:
				reward.UCor = utils.DecimalAddString(reward.UCor, re.Amount)
			case uord:
				reward.UOrd = utils.DecimalAddString(reward.UOrd, re.Amount)
			}
		}

		bytes, ok := db.Get(utils.MakeAddress(el.Delegation.DelegatorAddress))
		account := account.Account{}
		switch ok {
		//이미 있는 경우
		case true:
			account.FromBytes(bytes)
			//없는경우
		case false:
			fmt.Println("New Account")
			account.SetAccount(el.Delegation.DelegatorAddress)
		}

		account.UpdateUndelegate(chainCode, height)
		account.UpdateClaimAndReward(
			chainCode,
			el.Delegation.DelegatorAddress,
			el.Delegation.ValidatorAddress,
			reward)

		account.CalculateTotal(chainCode)
		if utils.MakeAddress(account.Address) == "atreides1dnwrgk5rj7zyyp35ad5nee62uxvalvf57p3qg0" {
			file, _ := json.MarshalIndent(account, "", " ")
			ioutil.WriteFile("./account.json", file, 0644)

			// fmt.Println(err)
		}
		db.Add(utils.MakeAddress(el.Delegation.DelegatorAddress), account.EncodeByte())

	}

}
