package controllers

import (
	"go_ride_backend-golang/initializers"
	"go_ride_backend-golang/models"

	"github.com/gin-gonic/gin"
)

// PostUserData handles the creation of a new user.
func PostUserData(c *gin.Context) {
	// Bind request body to User struct
	var body models.User
	c.Bind(&body)

	// Create user record in the database
	result := initializers.DB.Create(&body)
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Respond with success message and created user data
	c.JSON(200, gin.H{
		"status": "success",
		"data":   body,
	})
}

// GetAllUsersData retrieves all user records from the database.
func GetAllUsersData(c *gin.Context) {
	// Fetch all users from the database
	var users []models.User
	initializers.DB.Find(&users)

	// Respond with success message and user data
	c.JSON(200, gin.H{
		"status": "success",
		"data":   users,
	})
}

// GetUserDataById retrieves a user record by its ID from the database.
func GetUserDataById(c *gin.Context) {
	// Get user ID from request parameters
	id := c.Param("id")

	// Find user by ID
	var user models.User
	initializers.DB.First(&user, id)

	// Respond with success message and user data
	c.JSON(200, gin.H{
		"status": "success",
		"data":   user,
	})
}

// UpdateUserData updates a user record in the database.
func UpdateUserData(c *gin.Context) {
	// Get user ID from request parameters
	id := c.Param("id")

	// Bind request body to User struct
	var body models.User
	c.Bind(&body)

	// Find user by ID
	var user models.User
	initializers.DB.First(&user, id)

	// Update user record with new data
	initializers.DB.Model(&user).Updates(body)

	// Respond with success message and updated user data
	c.JSON(200, gin.H{
		"status": "success",
		"data":   user,
	})
}

// DeleteUserData deletes a user record from the database.
func DeleteUserData(c *gin.Context) {
	// Get user ID from request parameters
	id := c.Param("id")

	// Delete user by ID
	initializers.DB.Delete(&models.User{}, id)

	// Respond with success message
	c.JSON(200, gin.H{
		"status": "success",
	})
}
