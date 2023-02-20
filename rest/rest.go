package rest

import (
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const port = ":3306"

func Start(wg *sync.WaitGroup) {
	r := gin.Default()
	gin.ForceConsoleColor()

	//CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		MaxAge:       12 * time.Hour,
	}))

	r.GET("/", Root)
	r.GET("/uatr", UatrRank)
	r.GET("/uhar", UHarRank)
	r.GET("/ucor", UCorRank)
	r.GET("/uord", UOrdRank)
	r.GET("/scor", SCorRank)
	r.GET("/sord", SOrdRank)
	r.GET("/account/:address", UserReward)
	r.Run(port)
	defer wg.Done()
}
