// CompileDaemon -command="./go_ride_backend-golang"
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"go_ride_backend-golang/initializers"
	routers "go_ride_backend-golang/routes"
)

func init() {
	initializers.LoadConfigFile()
	initializers.ConnectToDB()
}

func main() {	
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	// Initialize Gin engine
	r := gin.Default()

	// Setup routes
	routers.SetupRoutes(r)

	// Start the server on the specified port in .env file
	envFile, _ := godotenv.Read(".env")
	port := envFile["port"]

	fmt.Println("Server Started")
	fmt.Println("Running at port: " + port)

	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}
