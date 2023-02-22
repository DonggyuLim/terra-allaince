package account

import (
	"bytes"
	"encoding/gob"
	"encoding/json"

	"github.com/DonggyuLim/Alliance-Rank/utils"
)

type Account struct {
	Address   string `bson:"address" json:"address"`
	Atreides  Chain  `bson:"atreides" json:"atreides"`
	Harkonnen Chain  `bson:"harkonnen" json:"harkonnen"`
	Corrino   Chain  `bson:"corrino" json:"corrino"`
	Ordos     Chain  `bson:"ordos" json:"ordos"`
	Total     Total  `bson:"total" json:"total"`
}
type Chain struct {
	Address string            `bson:"address" json:"address"`
	Rewards map[string]Reward `bson:"reward" json:"rewards"` //key = validator Address
	Claim   Claim             `bson:"claim" json:"claim"`
	Total   ChainTotal        `bson:"total" json:"total"`
}

type Reward struct {
	LastHeight uint `bson:"last_height" json:"last_height"`
	UAtr       uint `bson:"uatr" json:"uatr"`
	UHar       uint `bson:"uhar" json:"uhar"`
	UOrd       uint `bson:"uord" json:"uord"`
	UCor       uint `bson:"ucor" json:"ucor"`
	SCOR       uint `bson:"scor" json:"scor"`
	SORD       uint `bson:"sord" json:"sord"`
}

type Claim struct {
	UAtr uint `json:"uatr"`
	UCor uint `json:"ucor"`
	UHar uint `json:"uhar"`
	UOrd uint `json:"uord"`
	SCOR uint `json:"scor"`
	SORD uint `json:"sord"`
}
type Total struct {
	UAtr  uint `json:"uatr"`
	UCor  uint `json:"ucor"`
	UHar  uint `json:"uhar"`
	UOrd  uint `json:"uord"`
	SCOR  uint `json:"scor"`
	SORD  uint `json:"sord"`
	Total uint `json:"total"`
}

type ChainTotal struct {
	UAtr uint `json:"uatr"`
	UCor uint `json:"ucor"`
	UHar uint `json:"uhar"`
	UOrd uint `json:"uord"`
	SCOR uint `json:"scor"`
	SORD uint `json:"sord"`
	// Total uint `json:"total"`
}

func (a *Account) SetAccount(address, validator string, reward Reward, chainCode int) {
	m1 := make(map[string]Reward)
	m2 := make(map[string]Reward)
	m3 := make(map[string]Reward)
	m4 := make(map[string]Reward)
	a.Address = utils.MakeAddress(address)
	a.Atreides.Rewards = m1
	a.Harkonnen.Rewards = m2
	a.Corrino.Rewards = m3
	a.Ordos.Rewards = m4
	switch chainCode {
	case 0:
		a.Atreides.Rewards[validator] = reward
	case 1:
		a.Harkonnen.Rewards[validator] = reward
	case 2:
		a.Corrino.Rewards[validator] = reward
	case 3:
		a.Ordos.Rewards[validator] = reward
	}
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
		a.Atreides.Address = delegator
		o := a.Atreides.Rewards[validator]
		if o.UAtr > reward.UAtr {
			claim := o.UAtr - reward.UAtr
			a.Atreides.Claim.UAtr += claim
			a.Atreides.Claim.SCOR += o.SCOR
			a.Atreides.Claim.SORD += o.SORD
			a.Atreides.Rewards[validator] = reward
		} else {
			a.Atreides.Rewards[validator] = reward
		}

	case 1:
		a.Harkonnen.Address = delegator
		o := a.Harkonnen.Rewards[validator]
		if o.UHar > reward.UHar {
			claim := o.UHar - reward.UHar
			a.Harkonnen.Claim.UHar += claim
			a.Harkonnen.Claim.SCOR += o.SCOR
			a.Harkonnen.Claim.SORD += o.SORD
			a.Harkonnen.Rewards[validator] = reward
		} else {
			a.Harkonnen.Rewards[validator] = reward
		}

	case 2:
		a.Corrino.Address = delegator
		o := a.Corrino.Rewards[validator]
		if o.UCor > reward.UCor {
			claim := o.UCor - reward.UCor
			a.Corrino.Claim.UCor += claim
			a.Corrino.Claim.SCOR += o.SCOR
			a.Corrino.Claim.SORD += o.SORD
			a.Corrino.Rewards[validator] = reward
		} else {
			a.Corrino.Rewards[validator] = reward
		}
	case 3:
		a.Ordos.Address = delegator
		o := a.Ordos.Rewards[validator]
		if o.UOrd > reward.UOrd {
			claim := o.UOrd - reward.UOrd
			a.Corrino.Claim.UOrd += claim
			a.Corrino.Claim.SCOR += o.SCOR
			a.Corrino.Claim.SORD += o.SORD
			a.Ordos.Rewards[validator] = reward
		} else {
			a.Ordos.Rewards[validator] = reward
		}

	}
}

func (a *Account) UpdateUndelegate(chainCode, height int) {

	h := uint(height)
	switch chainCode {
	case 0:
		for k, v := range a.Atreides.Rewards {
			if (v.LastHeight + 1) != h {
				a.Atreides.Claim.UAtr += v.UAtr
				a.Atreides.Claim.SCOR += v.SCOR
				a.Atreides.Claim.SORD += v.SORD
				delete(a.Atreides.Rewards, k)
			}
		}

	case 1:
		for k, v := range a.Harkonnen.Rewards {
			if (v.LastHeight + 1) != h {
				a.Harkonnen.Claim.UHar += v.UHar
				a.Harkonnen.Claim.SCOR += v.SCOR
				a.Harkonnen.Claim.SORD += v.SORD
				delete(a.Harkonnen.Rewards, k)
			}
		}

	case 2:
		for k, v := range a.Corrino.Rewards {
			if (v.LastHeight + 1) != h {
				a.Corrino.Claim.UCor =
					a.Corrino.Claim.UCor + v.UCor
				a.Corrino.Claim.SCOR =
					a.Corrino.Claim.SCOR + v.SCOR
				a.Corrino.Claim.SORD =
					a.Corrino.Claim.SORD + v.SORD
				delete(a.Corrino.Rewards, k)
			}
		}

	case 3:
		for k, v := range a.Ordos.Rewards {
			if (v.LastHeight + 1) != h {
				a.Ordos.Claim.UOrd =
					a.Ordos.Claim.UOrd + v.UOrd
				a.Ordos.Claim.SCOR =
					a.Ordos.Claim.SCOR + v.SCOR
				a.Ordos.Claim.SORD =
					a.Ordos.Claim.SORD + v.SORD
				delete(a.Ordos.Rewards, k)
			}
		}
	}

}

func (a *Account) CalculateTotal(chainCode int) {

	ct := ChainTotal{}

	switch chainCode {
	case 0:

		for _, el := range a.Atreides.Rewards {
			ct.UAtr =
				ct.UAtr + el.UAtr
			ct.SCOR =
				ct.SCOR + el.SCOR
			ct.SORD =
				ct.SORD + el.SORD
		}
		//claim reward +
		ct.UAtr += a.Atreides.Claim.UAtr
		ct.SCOR += a.Atreides.Claim.SCOR
		ct.SORD += a.Atreides.Claim.SORD
		a.Atreides.Total = ct

		//harkonnen
	case 1:
		for _, el := range a.Harkonnen.Rewards {
			ct.UHar += el.UHar
			ct.SCOR += el.SCOR
			ct.SORD += el.SORD
		}
		//claim reward +
		ct.UHar += a.Harkonnen.Claim.UHar
		ct.SCOR += a.Harkonnen.Claim.SCOR
		ct.SORD += a.Harkonnen.Claim.SORD

		a.Harkonnen.Total = ct
		// a.Total = a.Total+ a.Harkonnen.Total.NativeTotal)+ a.Harkonnen.Total.SCOR)+ a.Harkonnen.Total.SORD)
	case 2:
		for _, el := range a.Corrino.Rewards {
			ct.UCor += el.UCor
			ct.SCOR += el.SCOR
			ct.SORD += el.SORD
		}
		//claim reward +
		ct.UCor += a.Corrino.Claim.UCor
		ct.SCOR += a.Corrino.Claim.SCOR
		ct.SORD += a.Corrino.Claim.SORD

		a.Corrino.Total = ct

	case 3:
		for _, el := range a.Ordos.Rewards {
			ct.UOrd += el.UOrd

			ct.SCOR += el.SCOR

			ct.SORD += el.SORD

		}
		//claim reward +
		ct.UOrd += a.Ordos.Claim.UOrd
		ct.SCOR += a.Ordos.Claim.SCOR
		ct.SORD += a.Ordos.Claim.SORD

		a.Ordos.Total = ct

		// a.Total = a.Total+ a.Ordos.Total.NativeTotal)+ a.Ordos.Total.SCOR)+ a.Ordos.Total.SORD)
	}
	a.Total = Total{}
	//calculate NativeTotal

	a.Total.UAtr = a.Atreides.Total.UAtr
	a.Total.UHar = a.Harkonnen.Total.UHar
	a.Total.UCor = a.Corrino.Total.UCor
	a.Total.UOrd = a.Ordos.Total.UOrd
	//calculate SCOR Total
	a.Total.SCOR = a.Atreides.Total.SCOR + a.Harkonnen.Total.SCOR + a.Corrino.Total.SCOR + a.Ordos.Total.SCOR
	///calculate SORD Total
	a.Total.SORD = a.Atreides.Total.SORD + a.Harkonnen.Total.SORD + a.Corrino.Total.SORD + a.Ordos.Total.SORD
	a.Total.Total = a.Total.UAtr + a.Total.UHar + a.Total.UCor + a.Total.UOrd + a.Total.SCOR + a.Total.SORD
}

func (r Reward) EncodeJson() string {
	bytes, err := json.MarshalIndent(r, "", "   ")
	utils.PanicError(err)
	return string(bytes)
}
