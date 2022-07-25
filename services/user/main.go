package main

import (
	"net/http"
	"os"

	"github.com/bl4nc/rocket_delivery/useCases/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := setupRouter()
	_ = r.Run(":" + os.Getenv("APP_PORT"))

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.POST("/user", user.CreateUser)

	return r

}
