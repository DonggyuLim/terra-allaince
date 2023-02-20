package account

import (
	"bytes"
	"encoding/gob"
	"encoding/json"

	"github.com/DonggyuLim/Alliance-Rank/utils"
	"github.com/shopspring/decimal"
)

type Account struct {
	Address   string `json:"address"`
	Atrides   Chain  `json:"atrides"`
	Harkonnen Chain  `json:"harkonnen"`
	Corrino   Chain  `json:"corrino"`
	Ordos     Chain  `json:"ordos"`
	Total     Total  `json:"reward_total"`
}
type Chain struct {
	Address string            `json:"address"`
	Rewards map[string]Reward `json:"rewards"` //key = validator Address
	Claim   Cliam             `json:"claim"`
	Total   Total             `json:"total"`
}

type Reward struct {
	LastHeight int             `json:"last_height"`
	UHar       decimal.Decimal `json:"uhar"`
	UOrd       decimal.Decimal `json:"uord"`
	UCor       decimal.Decimal `json:"ucor"`
	UAtr       decimal.Decimal `json:"uatr"`
	SCOR       decimal.Decimal `json:"scor"`
	SORD       decimal.Decimal `json:"sord"`
}

type Total struct {
	UAtr  decimal.Decimal `json:"uatr"`
	UCor  decimal.Decimal `json:"ucor"`
	UHar  decimal.Decimal `json:"uhar"`
	UOrd  decimal.Decimal `json:"uord"`
	SCOR  decimal.Decimal `json:"scor"`
	SORD  decimal.Decimal `json:"sord"`
	Total decimal.Decimal `json:"total"`
}
type Cliam struct {
	UAtr decimal.Decimal `json:"uatr"`
	UCor decimal.Decimal `json:"ucor"`
	UHar decimal.Decimal `json:"uhar"`
	UOrd decimal.Decimal `json:"uord"`
	SCOR decimal.Decimal `json:"scor"`
	SORD decimal.Decimal `json:"sord"`
}

func (a *Account) SetAccount(address string) {
	m1 := make(map[string]Reward)
	m2 := make(map[string]Reward)
	m3 := make(map[string]Reward)
	m4 := make(map[string]Reward)
	a.Address = utils.MakeAddress(address)
	a.Atrides.Rewards = m1
	a.Harkonnen.Rewards = m2
	a.Corrino.Rewards = m3
	a.Ordos.Rewards = m4

}

func (a Account) EncodeByte() []byte {
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	utils.PanicError(encoder.Encode(a))
	return aBuffer.Bytes()
}

func (a *Account) FromBytes(data []byte) {
	encoder := gob.NewDecoder(bytes.NewReader(data))
	utils.PanicError(encoder.Decode(&a))
}

func (a *Account) UpdateClaimAndReward(
	chainCode int,
	delegator,
	validator string,
	reward Reward) {
	switch chainCode {
	case 0:
		a.Atrides.Address = delegator
		origin := a.Atrides.Rewards[validator]
		if origin.UAtr.GreaterThan(reward.UAtr) {
			claim := origin.UAtr.Sub(reward.UAtr)
			a.Atrides.Claim.UAtr =
				a.Atrides.Claim.UAtr.Add(claim)
		}
		a.Atrides.Rewards[validator] = reward
	case 1:
		a.Harkonnen.Address = delegator
		origin := a.Harkonnen.Rewards[validator]
		if origin.UHar.GreaterThan(reward.UHar) {
			claim := origin.UHar.Sub(reward.UHar)
			a.Harkonnen.Claim.UHar =
				a.Harkonnen.Claim.UHar.Add(claim)
		}
		a.Harkonnen.Rewards[validator] = reward
	case 2:
		a.Corrino.Address = delegator
		origin := a.Corrino.Rewards[validator]
		if origin.UCor.GreaterThan(reward.UCor) {
			claim := origin.UCor.Sub(reward.UCor)
			a.Corrino.Claim.UCor =
				a.Corrino.Claim.UCor.Add(claim)
		}
		a.Corrino.Rewards[validator] = reward
	case 3:
		a.Ordos.Address = delegator
		origin := a.Ordos.Rewards[validator]
		if origin.UOrd.GreaterThan(reward.UOrd) {
			claim := origin.UOrd.Sub(reward.UOrd)
			a.Ordos.Claim.UOrd =
				a.Ordos.Claim.UOrd.Add(claim)
		}
		a.Ordos.Rewards[validator] = reward
	}
}

func (a *Account) UpdateUndelegate(chainCode, height int) {
	deleteKey := []string{}
	switch chainCode {
	case 0:
		for k, v := range a.Atrides.Rewards {
			if v.LastHeight < height {
				a.Atrides.Claim.UAtr =
					a.Atrides.Claim.UAtr.Add(v.UAtr)
				a.Atrides.Claim.SCOR =
					a.Atrides.Claim.SCOR.Add(v.SCOR)
				a.Atrides.Claim.SORD =
					a.Atrides.Claim.SORD.Add(v.SORD)
				deleteKey = append(deleteKey, k)
			}
		}

		//delete key
		for _, key := range deleteKey {
			delete(a.Atrides.Rewards, key)
		}
	case 1:
		for k, v := range a.Harkonnen.Rewards {
			if v.LastHeight < height {
				a.Harkonnen.Claim.UHar =
					a.Harkonnen.Claim.UHar.Add(v.UHar)
				a.Harkonnen.Claim.SCOR =
					a.Harkonnen.Claim.SCOR.Add(v.SCOR)
				a.Harkonnen.Claim.SORD =
					a.Harkonnen.Claim.SORD.Add(v.SORD)
				deleteKey = append(deleteKey, k)
			}
		}

		//delete key
		for _, key := range deleteKey {
			delete(a.Harkonnen.Rewards, key)
		}
	case 2:
		for k, v := range a.Corrino.Rewards {
			if v.LastHeight < height {
				a.Corrino.Claim.UCor =
					a.Corrino.Claim.UCor.Add(v.UCor)
				a.Corrino.Claim.SCOR =
					a.Corrino.Claim.SCOR.Add(v.SCOR)
				a.Corrino.Claim.SORD =
					a.Corrino.Claim.SORD.Add(v.SORD)
				deleteKey = append(deleteKey, k)
			}
		}

		//delete key
		for _, key := range deleteKey {
			delete(a.Corrino.Rewards, key)
		}
	case 3:
		for k, v := range a.Ordos.Rewards {
			if v.LastHeight < height {
				a.Ordos.Claim.UOrd =
					a.Ordos.Claim.UOrd.Add(v.UOrd)
				a.Ordos.Claim.SCOR =
					a.Ordos.Claim.SCOR.Add(v.SCOR)
				a.Ordos.Claim.SORD =
					a.Ordos.Claim.SORD.Add(v.SORD)
				deleteKey = append(deleteKey, k)
			}
		}

		//delete key
		for _, key := range deleteKey {
			delete(a.Ordos.Rewards, key)
		}
	}

}

func (a *Account) CalculateTotal(chainCode int) {

	ct := Total{}

	switch chainCode {
	case 0:

		for _, el := range a.Atrides.Rewards {
			ct.UAtr =
				ct.UAtr.Add(el.UAtr)
			ct.SCOR =
				ct.SCOR.Add(el.SCOR)
			ct.SORD =
				ct.SORD.Add(el.SORD)
		}
		//claim reward +
		ct.UAtr =
			ct.UAtr.Add(a.Atrides.Claim.UAtr)
		ct.SCOR =
			ct.SCOR.Add(a.Atrides.Claim.SCOR)
		ct.SORD =
			ct.SORD.Add(a.Atrides.Claim.SORD)
		ct.Total =
			ct.UAtr.Add(ct.SCOR).Add(ct.SORD)
		a.Atrides.Total = ct

		//harkonnen
	case 1:
		for _, el := range a.Harkonnen.Rewards {
			ct.UHar =
				ct.UHar.Add(el.UHar)
			ct.SCOR =
				ct.SCOR.Add(el.SCOR)
			ct.SORD =
				ct.SORD.Add(el.SORD)
		}
		//claim reward +
		ct.UHar =
			ct.UHar.Add(a.Harkonnen.Claim.UHar)
		ct.SCOR =
			ct.SCOR.Add(a.Harkonnen.Claim.SCOR)
		ct.SORD =
			ct.SORD.Add(a.Harkonnen.Claim.SORD)
		ct.Total =
			ct.UHar.Add(ct.SCOR).Add(ct.SORD)
		a.Harkonnen.Total = ct
		// a.Total = a.Total.Add(a.Harkonnen.Total.NativeTotal).Add(a.Harkonnen.Total.SCOR).Add(a.Harkonnen.Total.SORD)
	case 2:
		for _, el := range a.Corrino.Rewards {
			ct.UCor =
				ct.UCor.Add(el.UCor)
			ct.SCOR =
				ct.SCOR.Add(el.SCOR)
			ct.SORD =
				ct.SORD.Add(el.SORD)
		}
		//claim reward +
		ct.UCor =
			ct.UCor.Add(a.Corrino.Claim.UCor)
		ct.SCOR =
			ct.SCOR.Add(a.Corrino.Claim.SCOR)
		ct.SORD =
			ct.SORD.Add(a.Corrino.Claim.SORD)
		ct.Total =
			ct.UCor.Add(ct.SCOR).Add(ct.SORD)
		a.Corrino.Total = ct
		// a.Total = a.Total.Add(a.Corrino.Total.NativeTotal).Add(a.Corrino.Total.SCOR).Add(a.Corrino.Total.SORD)
	case 3:
		for _, el := range a.Ordos.Rewards {
			ct.UOrd =
				ct.UOrd.Add(el.UOrd)
			ct.SCOR =
				ct.SCOR.Add(el.SCOR)
			ct.SORD =
				ct.SORD.Add(el.SORD)
		}
		//claim reward +
		ct.UOrd =
			ct.UOrd.Add(a.Ordos.Claim.UOrd)
		ct.SCOR =
			ct.SCOR.Add(a.Ordos.Claim.SCOR)
		ct.SORD =
			ct.SORD.Add(a.Ordos.Claim.SORD)
		ct.Total =
			ct.UOrd.Add(ct.SCOR).Add(ct.SORD)
		a.Ordos.Total = ct

		// a.Total = a.Total.Add(a.Ordos.Total.NativeTotal).Add(a.Ordos.Total.SCOR).Add(a.Ordos.Total.SORD)
	}
	a.Total = Total{}
	//calculate NativeTotal
	a.Total.UAtr = a.Total.UAtr.Add(a.Atrides.Total.UAtr)
	a.Total.UHar = a.Total.UHar.Add(a.Harkonnen.Total.UHar)
	a.Total.UCor = a.Total.UCor.Add(a.Corrino.Total.UCor)
	a.Total.UOrd = a.Total.UOrd.Add(a.Ordos.Total.UOrd)

	//calculate SCOR Total
	a.Total.SCOR = a.Total.SCOR.
		Add(a.Atrides.Total.SCOR).
		Add(a.Harkonnen.Total.SCOR).
		Add(a.Corrino.Total.SCOR).
		Add(a.Ordos.Total.SCOR)

	///calculate SORD Total
	a.Total.SORD = a.Total.SORD.
		Add(a.Atrides.Total.SORD).
		Add(a.Harkonnen.Total.SORD).
		Add(a.Corrino.Total.SORD).
		Add(a.Ordos.Total.SORD)

	a.Total.Total =
		a.Total.UAtr.Add(a.Total.UHar).Add(a.Total.UCor).Add(a.Total.UOrd).Add(a.Total.SCOR).Add(a.Total.SORD)
}

func (r Reward) EncodeJson() string {
	bytes, err := json.MarshalIndent(r, "", "   ")
	utils.PanicError(err)
	return string(bytes)
}

// func (r Reward) GetReward(endpint string, chainCode int) {
// 	client := req.C.R()

// }
