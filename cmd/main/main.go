package main

import (
	"github.com/anirudh/project-1/pkg/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run()
}
