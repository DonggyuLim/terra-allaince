package rest

import (
	"fmt"

	"github.com/DonggyuLim/Alliance-Rank/account"
	"github.com/DonggyuLim/Alliance-Rank/db"

	"github.com/gin-gonic/gin"
)

// func Root(c *gin.Context) {
// 	list, err := db.Find("", "", "totalReward", 100)
// 	// fmt.Println(list)
// 	if err != nil {
// 		fmt.Println(err)
// 		c.String(404, err.Error())
// 		return
// 	}
// 	c.JSON(200, list)
// }

func UserReward(c *gin.Context) {
	a := c.Param("account")
	fmt.Println(a)

	bytes, ok := db.Get(a)
	var account account.Account

	account.FromBytes(bytes)
	if ok {
		c.JSON(200, account)
	} else {
		c.String(404, "Not Account")
		// fmt.Println(account)
		// c.JSON(404, account)
	}

	// if err != nil {

	// }
	// // fmt.Println(data)

}
