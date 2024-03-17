package routers

import (
	"database/sql"
	"golang-API-save-data-sql/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {

	// User End points
	r.POST("/post-user", func(c *gin.Context) {
		services.PostUserData(c.Writer, c.Request, db)
	})
	r.GET("/get-user", func(c *gin.Context) {
		services.GetUserData(c.Writer, c.Request, db)
	})

	// Feedback(Rating) End points
	r.POST("/post-feedback", func(c *gin.Context) {
		services.PostFeedback(c.Writer, c.Request, db)
	})
	r.GET("/get-feedback", func(c *gin.Context) {
		services.GetFeedback(c.Writer, c.Request, db)
	})
}
