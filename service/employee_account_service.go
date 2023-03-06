package service

import (
	"fmt"
	"myapp/model"
	"myapp/tool"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *Service) EmployeeAccountLogin(c *gin.Context, request model.EmployeeAccountLoginParam) (*model.EmployeeAccountLoginResponse, error) {
	var (
		err             error
		employeeAccount *model.EmployeeAccount
		accessToken     string
	)

	employeeAccount, err = s.EmployeeAccountGetByLoginUsername(c, request.LoginUsername)
	if err != nil {
		return nil, err
	}
	if employeeAccount == nil {
		return nil, fmt.Errorf("account with such username not found")
	}

	if ok := tool.CheckPasswordHash(request.LoginPassword, employeeAccount.LoginPassword); !ok {
		return nil, fmt.Errorf("username or password not match")
	}

	accessToken = tool.TokenCreate(employeeAccount.EmployeeID)

	return &model.EmployeeAccountLoginResponse{
		EmployeeAccount: employeeAccount,
		AccessToken:     accessToken,
	}, nil
}

func (s *Service) EmployeeAccountRegister(c *gin.Context, request model.EmployeeAccountRegisterParam) (*model.EmployeeAccount, error) {
	var (
		err             error
		employeeAccount *model.EmployeeAccount
	)

	employeeAccount, err = s.EmployeeAccountGetByLoginUsername(c, request.LoginUsername)
	if err != nil {
		return nil, err
	}

	if employeeAccount != nil {
		return nil, fmt.Errorf("username has been used before, please choose another username")
	}

	employeeAccount, err = s.EmployeeAccountCreate(c, request)
	if err != nil {
		return nil, err
	}

	return employeeAccount, nil
}

func (s *Service) EmployeeAccountGetByLoginUsername(c *gin.Context, loginUsername string) (*model.EmployeeAccount, error) {
	var (
		employeeAccount model.EmployeeAccount
	)

	if err := s.DB.Where("login_username = ?", loginUsername).First(&employeeAccount).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &employeeAccount, nil
}

func (s *Service) EmployeeAccountUpdateLoginedAt(c *gin.Context, accountID int) error {
	var (
		currentTime = time.Now()
	)

	if err := s.DB.Model(&model.EmployeeAccount{}).Where("account_id = ?", accountID).UpdateColumn("logined_at", currentTime).Error; err != nil {
		return err
	}

	return nil
}

func (s *Service) EmployeeAccountCreate(c *gin.Context, request model.EmployeeAccountRegisterParam) (*model.EmployeeAccount, error) {
	var (
		currentTime     = time.Now()
		employeeAccount = model.EmployeeAccount{
			EmployeeID:    request.EmployeeID,
			LoginUsername: request.LoginUsername,
			LoginPassword: tool.HashPassword(request.LoginPassword),
			RoleID:        request.RoleID,
			CreatedAt:     currentTime,
		}
	)

	if err := s.DB.Create(&employeeAccount).Error; err != nil {
		return nil, err
	}

	return &employeeAccount, nil
}
