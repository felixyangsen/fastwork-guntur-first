package service

import (
	"fmt"
	"myapp/model"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Service) EmployeeGetByID(c *gin.Context, id int) (*model.Employee, error) {
	var employee model.Employee

	if err := s.DB.First(&employee, id).Error; err != nil {
		return nil, err
	}

	return &employee, nil
}

func (s *Service) EmployeeGetAll(c *gin.Context) ([]*model.Employee, error) {
	var employees []*model.Employee

	if err := s.DB.Find(&employees).Error; err != nil {
		return nil, err
	}

	return employees, nil
}

func (s *Service) EmployeeCreate(c *gin.Context, newEmployee model.NewEmployee) (*model.Employee, error) {
	var (
		currentTime = time.Now()
		employee = model.Employee{
			AgentID: newEmployee.AgentID,
			Name: newEmployee.Name,
			Email: newEmployee.Email,
			Phone: newEmployee.Phone,	
			CreatedAt: currentTime,
			CreatedLoc: newEmployee.CreatedLoc,
			CreatedBy: newEmployee.CreatedBy,
		}
	)

	if err := s.DB.Create(&employee).Error; err != nil {
		return nil, err
	}

	return &employee, nil
}

func (s *Service) EmployeeUpdate(c *gin.Context, updateEmployee model.UpdateEmployee) (*model.Employee, error) {
	var (		
		employee = model.Employee{
			AgentID: updateEmployee.AgentID,
			Name: updateEmployee.Name,
			Email: updateEmployee.Email,
			Phone: updateEmployee.Phone,
		}
	)

	if err := s.DB.Updates(&employee).Error; err != nil {
		return nil, err
	}

	return &employee, nil
}

func (s *Service) EmployeeDelete(c *gin.Context, id int) (string, error) {
	if err := s.DB.Delete(&model.Employee{}, id).Error; err != nil {
		return "", err
	}
	
	return fmt.Sprintf("action success, delete employee id %d", id), nil
}