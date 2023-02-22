package data

import (
	"fmt"
	"sort"
	"strconv"
	"sync"

	"github.com/DonggyuLim/Alliance-Rank/account"
	"github.com/DonggyuLim/Alliance-Rank/db"
	"github.com/DonggyuLim/Alliance-Rank/utils"
	"go.mongodb.org/mongo-driver/bson"
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

func Main(wg *sync.WaitGroup) {

	defer wg.Done()
	// height := 438

	height, _ := strconv.Atoi(utils.LoadENV("HEIGHT", "height.env"))
	fmt.Println(height)
	for {
		latestBlockHeight := []int{GetLastBlock(ATREIDES), GetLastBlock(CORRINO), GetLastBlock(Harkonnen), GetLastBlock(ORDOS)}
		sort.Ints(latestBlockHeight)
		if height > latestBlockHeight[0] {
			height = latestBlockHeight[len(latestBlockHeight)-1]
		}

		MakeData(height, ATREIDES)  //atr
		MakeData(height, Harkonnen) //har
		MakeData(height, CORRINO)   //cor
		MakeData(height, ORDOS)     //ord
		height += 1
		utils.WriteENV("HEIGHT", strconv.Itoa(height), "height.env")
	}
}

func MakeData(height, chainCode int) {

	latestBlockHeight := GetLastBlock(chainCode)

	if height > latestBlockHeight {
		height = latestBlockHeight
		fmt.Println(height)
	}

	delegations, err := GetDelegations(height, chainCode)
	if len(delegations.Deligations) == 0 || err != nil {
		fmt.Printf("%v %v Not Delegate\n", chainCode, height)
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
			fmt.Printf("%v %v Not Reward\n", chainCode, height)
			continue
		}
		reward := account.Reward{
			LastHeight: uint(height),
			UAtr:       0,
			UHar:       0,
			UCor:       0,
			UOrd:       0,
			SCOR:       0,
			SORD:       0,
		}

		//GET REWARD
		for _, re := range resReward {
			switch re.Denom {
			case sCOR:
				amount, err := strconv.Atoi(re.Amount)
				utils.PanicError(err)
				reward.SCOR = uint(amount)

			case sORD:
				amount, err := strconv.Atoi(re.Amount)
				utils.PanicError(err)
				reward.SORD = uint(amount)
			case uatr:
				amount, err := strconv.Atoi(re.Amount)
				utils.PanicError(err)
				reward.UAtr = uint(amount)
			case uhar:
				amount, err := strconv.Atoi(re.Amount)
				utils.PanicError(err)
				reward.UHar = uint(amount)
			case ucor:
				amount, err := strconv.Atoi(re.Amount)
				utils.PanicError(err)
				reward.UCor = uint(amount)
			case uord:
				amount, err := strconv.Atoi(re.Amount)
				utils.PanicError(err)
				reward.UOrd = uint(amount)
			}
		}

		account := account.Account{}

		filter := bson.D{{Key: "address", Value: utils.MakeAddress(el.Delegation.DelegatorAddress)}}
		ok := db.FindOne(filter, &account)

		if ok != nil {
			account.SetAccount(
				el.Delegation.DelegatorAddress,
				el.Delegation.ValidatorAddress,
				reward,
				chainCode,
			)
			db.Insert(account)
			continue
		}

		account.UpdateClaimAndReward(
			chainCode,
			el.Delegation.DelegatorAddress,
			el.Delegation.ValidatorAddress,
			reward)

		// account.UpdateUndelegate(chainCode, height)
		account.CalculateTotal(chainCode)

		db.ReplaceOne(bson.D{{Key: "address", Value: account.Address}}, account)

	}

}
