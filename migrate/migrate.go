package main

import (
	"go_ride_backend-golang/initializers"
	"go_ride_backend-golang/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
