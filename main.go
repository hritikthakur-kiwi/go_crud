package main

import (
	"go_crud/initializers"
	"go_crud/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	routes.UserRoutes(r)
	routes.PostRoutes(r)
	r.Run()
}
