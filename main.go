// CompileDaemon -command="./go_ride_backend-golang"
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"

	"go_ride_backend-golang/database"
	"go_ride_backend-golang/initializers"
	routers "go_ride_backend-golang/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {	
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	// Initialize the database
	DB, err := database.InitializeDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer DB.Close()

	// Initialize Gin engine
	r := gin.Default()

	// Setup routes
	routers.SetupRoutes(r, DB)

	// Start the server on the specified port in .env file
	envFile, _ := godotenv.Read(".env")
	port := envFile["port"]

	fmt.Println("Server Started")
	fmt.Println("Running at port: " + port)

	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}
