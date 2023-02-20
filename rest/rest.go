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

	// r.GET("/", Root)
	r.GET("/:account", UserReward)
	r.Run(port)
	defer wg.Done()
}
