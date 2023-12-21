package routes

import (
	"go_crud/controllers"

	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {
	r.POST("/addpost", controllers.AddPost)
}
