package controller

import (
	"myapp/model"
	"myapp/service"
	"myapp/tool"
	"net/http"
	"strconv"

	"myapp/middleware"

	"github.com/gin-gonic/gin"
)

func EmployeeGetDetail(c *gin.Context) {
	var (
		s       = service.GetService()
		err     error
		ctxData = middleware.AuthContext(c)

		id       = c.Param("id")
		IntID, _ = strconv.Atoi(id)

		employee *model.Employee
	)

	if ctxData.ID != IntID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "cannot access another employee data"})
		return
	}

	employee, err = s.EmployeeGetByID(c, IntID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func EmployeeGetlist(c *gin.Context) {
	var (
		s   = service.GetService()
		err error

		employees []*model.Employee
	)

	employees, err = s.EmployeeGetAll(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}

func EmployeeCreate(c *gin.Context) {
	var (
		s   = service.GetService()
		err error

		newEmployee model.NewEmployee

		employee *model.Employee
	)

	if err := c.ShouldBindJSON(&newEmployee); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = tool.ValidateStruct(newEmployee)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err = s.EmployeeCreate(c, newEmployee)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func EmployeeUpdate(c *gin.Context) {
	var (
		s   = service.GetService()
		err error

		updateEmployee model.UpdateEmployee

		employee *model.Employee
	)

	if err := c.ShouldBindJSON(&updateEmployee); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = tool.ValidateStruct(updateEmployee)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err = s.EmployeeUpdate(c, updateEmployee)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func EmployeeDelete(c *gin.Context) {
	var (
		s   = service.GetService()
		err error

		id       = c.Param("id")
		IntID, _ = strconv.Atoi(id)

		message string
	)

	message, err = s.EmployeeDelete(c, IntID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}
