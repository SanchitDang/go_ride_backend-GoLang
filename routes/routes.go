package routers

import (
	"go_ride_backend-golang/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// User End points
	r.GET("/get-all-users", controllers.GetAllUsersData)
	r.GET("/get-user/:id", controllers.GetUserDataById)
	r.POST("/post-user", controllers.PostUserData)
	r.PUT("/update-user/:id", controllers.UpdateUserData)
	r.DELETE("/delete-user/:id", controllers.DeleteUserData)

}
