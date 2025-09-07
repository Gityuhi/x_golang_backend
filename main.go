package main

import ("github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	router.Use(cors.New(config))

	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "good",
		})
	})
	router.Run()
}
