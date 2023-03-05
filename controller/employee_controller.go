package controller

import (
	"myapp/model"
	"myapp/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EmployeeGetDetail(c *gin.Context) {
	var (
		s = service.GetService()
		err error

		id = c.Param("id")
		IntID, _ = strconv.Atoi(id)

		employee *model.Employee
	)

	employee, err = s.EmployeeGetByID(c, IntID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func EmployeeGetlist(c *gin.Context) {
	var (
		s = service.GetService()
		err error

		employees []*model.Employee
	)

	employees, err = s.EmployeeGetAll(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}

func EmployeeCreate(c *gin.Context) {
	var (
		s = service.GetService()
		err error

		newEmployee model.NewEmployee
		
		employee *model.Employee
	)

	if err := c.ShouldBindJSON(&newEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err = s.EmployeeCreate(c, newEmployee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func EmployeeUpdate(c *gin.Context) {
	var (
		s = service.GetService()
		err error

		updateEmployee model.UpdateEmployee

		employee *model.Employee
	)

	if err := c.ShouldBindJSON(&updateEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err = s.EmployeeUpdate(c, updateEmployee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func EmployeeDelete(c *gin.Context) {
	var (
		s = service.GetService()
		err error

		id = c.Param("id")
		IntID, _ = strconv.Atoi(id)

		message string
	)

	message, err = s.EmployeeDelete(c,  IntID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, message)
}