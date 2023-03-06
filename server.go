package main

import (
	"myapp/config"
	"myapp/controller"
	"myapp/tool"
	"os"

	"myapp/middleware"

	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func init() {
	config.ConnectDB()
	tool.InitValidator()
}

func main() {
	db := config.GetDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := gin.Default()

	router.POST("/login", controller.EmployeeAccountLogin)
	router.POST("/register", controller.EmployeeAccountRegister)

	authGroup := router.Group("")
	authGroup.Use(middleware.Authorize())

	authGroup.GET("/employee/:id", controller.EmployeeGetDetail)
	authGroup.GET("/employees", controller.EmployeeGetlist)
	authGroup.POST("/employee", controller.EmployeeCreate)
	authGroup.PUT("/employee", controller.EmployeeUpdate)
	authGroup.DELETE("/employee/:id", controller.EmployeeDelete)

	router.Run("localhost:" + port)
}
