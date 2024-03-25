// CompileDaemon -command="./go_ride_backend-golang"
package main

import (
	"github.com/gin-gonic/gin"

	"go_ride_backend-golang/initializers"
	routers "go_ride_backend-golang/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.LoadConfigFile()
	initializers.ConnectToDB()
}

func main() {
	// Initialize Gin engine
	r := gin.Default()

	// Setup routes
	routers.SetupRoutes(r)

	// Run the Server
	r.Run()
}
