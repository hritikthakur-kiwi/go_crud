package routes

import (
	"go_crud/controllers"
	auth "go_crud/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "helloj",
		})
	})
	r.POST("/addUser", controllers.Create)
	r.POST("/login", controllers.Login)
	r.PUT("/updateUser", auth.VerifyUser, controllers.UpdateUser)
}
