package rest

import (
	"sync"
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
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
	r.Use(gin.Recovery())
	memoryStore := persist.NewMemoryStore(1 * time.Minute)
	cache := cache.CacheByRequestPath(memoryStore, 10*time.Second)

	r.GET("/", Root)
	r.GET("/uatr", cache, UatrRank)
	r.GET("/uhar", cache, UHarRank)
	r.GET("/ucor", cache, UCorRank)
	r.GET("/uord", cache, UOrdRank)
	r.GET("/scor", cache, SCorRank)
	r.GET("/sord", cache, SOrdRank)
	r.GET("/account/:address", UserReward)
	r.Run(port)
	defer wg.Done()
}
