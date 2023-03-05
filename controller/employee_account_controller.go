package controller

import (
	"myapp/model"
	"myapp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EmployeeAccountLogin(c *gin.Context) {
	var (
		s = service.GetService()
		err error

		param model.EmployeeAccountLoginParam

		response *model.EmployeeAccountLoginResponse
	)

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err = s.EmployeeAccountLogin(c, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func EmployeeAccountRegister(c *gin.Context) {
	var (
		s = service.GetService()
		err error

		param model.EmployeeAccountRegisterParam

		employeeAccount *model.EmployeeAccount
	)

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employeeAccount, err = s.EmployeeAccountRegister(c, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, employeeAccount)
}