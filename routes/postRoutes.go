package routes

import (
	"go_crud/controllers"
	auth "go_crud/middleware"

	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {
	r.POST("/addpost", auth.VerifyUser, controllers.AddPost)
}
