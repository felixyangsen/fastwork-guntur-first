package main

import (
	"myapp/config"
	"myapp/controller"
	"os"

	"myapp/middleware"

	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func init() {
	config.ConnectDB()
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

	authRouter := router.Use(middleware.Authorize()) 

	authRouter.GET("/employee:id", controller.EmployeeGetDetail)
	authRouter.GET("/employees", controller.EmployeeGetlist)
	authRouter.POST("/employee", controller.EmployeeCreate)
	authRouter.PUT("/employee", controller.EmployeeUpdate)
	authRouter.DELETE("/employee:id", controller.EmployeeDelete)

	router.Run("localhost:"+port)
}