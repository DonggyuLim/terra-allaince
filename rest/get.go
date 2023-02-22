package rest

import (
	"fmt"

	"github.com/DonggyuLim/Alliance-Rank/account"
	"github.com/DonggyuLim/Alliance-Rank/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

type ToTalResponse struct {
	Address string `json:"address"`
	UAtr    string `json:"uatr"`
	UCor    string `json:"ucor"`
	UHar    string `json:"uhar"`
	UOrd    string `json:"uord"`
	SCOR    string `json:"scor"`
	SORD    string `json:"sord"`
	Total   string `json:"total"`
}

func Root(c *gin.Context) {
	fmt.Println("ROOT")
	list, err := db.Find("", "", "total.total", 100)
	// fmt.Println(list)
	var res []ToTalResponse
	for _, el := range list {
		total := ToTalResponse{
			Address: el.Address,
			UAtr:    fmt.Sprintf("%f", el.Total.UAtr),
			UCor:    fmt.Sprintf("%f", el.Total.UCor),
			UHar:    fmt.Sprintf("%f", el.Total.UHar),
			UOrd:    fmt.Sprintf("%f", el.Total.UOrd),
			SCOR:    fmt.Sprintf("%f", el.Total.SCOR),
			SORD:    fmt.Sprintf("%f", el.Total.SORD),
			Total:   fmt.Sprintf("%f", el.Total.Total),
		}
		res = append(res, total)
	}
	if err != nil {
		fmt.Println(err)
		c.String(404, err.Error())
		return
	}
	c.JSON(200, res)
}

type UAtrResponse struct {
	Address string `json:"address"`
	UAtr    string `json:"amount"`
}

func UatrRank(c *gin.Context) {
	list, err := db.Find("", "", "total.uatr", 100)
	var res []UAtrResponse
	for _, el := range list {
		atr := UAtrResponse{
			Address: el.Atreides.Address,
			UAtr:    fmt.Sprintf("%f", el.Total.UAtr),
		}
		res = append(res, atr)
	}
	if err != nil {
		fmt.Println(err)
		c.String(404, err.Error())
		return
	}
	c.JSON(200, res)
}

type UharResponse struct {
	Address string `json:"address"`
	UHar    string `json:"amount"`
}

func UHarRank(c *gin.Context) {
	list, err := db.Find("", "", "total.uhar", 100)
	var res []UharResponse
	for _, el := range list {
		uhar := UharResponse{
			Address: el.Harkonnen.Address,
			UHar:    fmt.Sprintf("%f", el.Total.UHar),
		}
		res = append(res, uhar)
	}
	if err != nil {
		fmt.Println(err)
		c.String(404, err.Error())
		return
	}
	c.JSON(200, res)
}

type UCorResponse struct {
	Address string `json:"address"`
	UCor    string `json:"amount"`
}

func UCorRank(c *gin.Context) {
	list, err := db.Find("", "", "total.ucor", 100)
	var res []UCorResponse
	for _, el := range list {
		ucor := UCorResponse{
			Address: el.Corrino.Address,
			UCor:    fmt.Sprintf("%f", el.Total.UCor),
		}
		res = append(res, ucor)
	}
	if err != nil {
		fmt.Println(err)
		c.String(404, err.Error())
		return
	}
	c.JSON(200, res)
}

type UOrdResponse struct {
	Address string `json:"address"`
	UOrd    string `json:"amount"`
}

func UOrdRank(c *gin.Context) {
	list, err := db.Find("", "", "total.ord", 100)
	var res []UOrdResponse
	for _, el := range list {
		uord := UOrdResponse{
			Address: el.Ordos.Address,
			UOrd:    fmt.Sprintf("%f", el.Total.UOrd),
		}
		res = append(res, uord)
	}
	if err != nil {
		fmt.Println(err)
		c.String(404, err.Error())
		return
	}
	c.JSON(200, res)
}

type ScorResponse struct {
	Address string `json:"address"`
	SCor    string `json:"amount"`
}

func SCorRank(c *gin.Context) {
	list, err := db.Find("", "", "total.scor", 100)
	var res []ScorResponse
	for _, el := range list {
		scor := ScorResponse{
			Address: el.Address,
			SCor:    fmt.Sprintf("%f", el.Total.SCOR),
		}
		res = append(res, scor)
	}
	if err != nil {
		fmt.Println(err)
		c.String(404, err.Error())
		return
	}
	c.JSON(200, res)
}

type SOrdResponse struct {
	Address string `json:"address"`
	Sord    string `json:"amount"`
}

func SOrdRank(c *gin.Context) {
	list, err := db.Find("", "", "total.sord", 100)
	var res []SOrdResponse
	for _, el := range list {
		sord := SOrdResponse{
			Address: el.Address,
			Sord:    fmt.Sprintf("%f", el.Total.SORD),
		}
		res = append(res, sord)
	}
	if err != nil {
		fmt.Println(err)
		c.String(404, err.Error())
		return
	}
	c.JSON(200, res)
}

type MyRewardResponse struct {
	Address string `json:"address"`
	UAtr    string `json:"uatr"`
	UHar    string `json:"uhar"`
	UCor    string `json:"ucor"`
	UOrd    string `json:"uord"`
	SCor    string `json:"scor"`
	SOrd    string `json:"sord"`
}

func UserReward(c *gin.Context) {
	// address := c.Query("address")
	address := c.Param("address")
	// address := "atreides1qqczsqkqh4nnrxz3xwcfzydfman9nsltnujta4"
	fmt.Println(address)
	account := account.Account{}
	filter := bson.D{{Key: "address", Value: address}}
	ok := db.FindOne(filter, &account)
	switch ok {
	case nil:
		myReward := MyRewardResponse{
			Address: account.Address,
			UAtr:    fmt.Sprintf("%f", account.Total.UAtr),
			UHar:    fmt.Sprintf("%f", account.Total.UHar),
			UCor:    fmt.Sprintf("%f", account.Total.UCor),
			UOrd:    fmt.Sprintf("%f", account.Total.UOrd),
			SCor:    fmt.Sprintf("%f", account.Total.SCOR),
			SOrd:    fmt.Sprintf("%f", account.Total.SORD),
		}
		c.JSON(200, myReward)
	case mongo.ErrNoDocuments:
		// fmt.Println(address)
		c.String(404, "Invalid Address")
	}
}
