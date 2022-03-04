package routes

import (
	"github.com/anirudh/project-1/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterRoutes = func(router *gin.Engine) {
	router.GET("/users", func(c *gin.Context) {
		controllers.GetUsers(c)
	})

	router.POST("/users", func(c *gin.Context) {
		controllers.AddUser(c)
	})
}
