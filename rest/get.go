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

	list, err := db.Find("", "", "total.total", 100)
	// fmt.Println(list)
	var res []ToTalResponse
	for _, el := range list {
		total := ToTalResponse{
			Address: el.Address,
			UAtr:    el.Total.UAtr,
			UCor:    el.Total.UCor,
			UHar:    el.Total.UHar,
			UOrd:    el.Total.UOrd,
			SCOR:    el.Total.SCOR,
			SORD:    el.Total.SORD,
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
	UAtr    string `json:"uatr"`
}

func UatrRank(c *gin.Context) {
	list, err := db.Find("", "", "total.uatr", 100)
	var res []UAtrResponse
	for _, el := range list {
		atr := UAtrResponse{
			Address: el.Address,
			UAtr:    el.Total.UAtr,
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
func UHarRank(c *gin.Context) {
	list, err := db.Find("", "", "total.uhar", 100)

	if err != nil {
		fmt.Println(err)
		c.String(404, err.Error())
		return
	}
	c.JSON(200, list)
}
func UCorRank(c *gin.Context) {
	list, err := db.Find("", "", "total.ucor", 100)

	if err != nil {
		fmt.Println(err)
		c.String(404, err.Error())
		return
	}
	c.JSON(200, list)
}
func UOrdRank(c *gin.Context) {
	list, err := db.Find("", "", "total.ord", 100)

	if err != nil {
		fmt.Println(err)
		c.String(404, err.Error())
		return
	}
	c.JSON(200, list)
}
func SCorRank(c *gin.Context) {
	list, err := db.Find("", "", "total.scor", 100)

	if err != nil {
		fmt.Println(err)
		c.String(404, err.Error())
		return
	}
	c.JSON(200, list)
}
func SOrdRank(c *gin.Context) {
	list, err := db.Find("", "", "total.sord", 100)

	if err != nil {
		fmt.Println(err)
		c.String(404, err.Error())
		return
	}
	c.JSON(200, list)
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
		// fmt.Println(ok, account)
		c.JSON(200, account)
	case mongo.ErrNoDocuments:
		// fmt.Println(address)
		c.String(404, "Not Account")
	}
}
