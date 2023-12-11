package main

import (
	"go_crud/controllers"
	"go_crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "helloj",
		})
	})
	r.POST("/addUser", controllers.Create)
	r.POST("/login", controllers.Login)
	r.Run()
}
