package model

import "time"

type EmployeeAccount struct {
	AccountID     int       `json:"account_id"`
	EmployeeID    int       `json:"employee_id"`
	LoginUsername string    `json:"login_username"`
	LoginPassword string    `json:"login_password"`
	RoleID        int       `json:"role_id"`
	CreatedAt     time.Time `json:"created_at"`
	LoginedAt *time.Time `json:"logined_at"`
	Rev *string `json:"rev"`
}

type EmployeeAccountLoginParam struct {
	LoginUsername string    `json:"login_username"`
	LoginPassword string    `json:"login_password"`
}

type EmployeeAccountLoginResponse struct {
	EmployeeAccount *EmployeeAccount `json:"employee_account"`
	AccessToken string `json:"access_token"`
}

type EmployeeAccountRegisterParam struct {
	EmployeeID    int       `json:"employee_id" validate:"required"`
	LoginUsername string    `json:"login_username" validate:"required"`
	LoginPassword string    `json:"login_password" validate:"required,alphanum"`
	RoleID        int       `json:"role_id" validate:"required"`
}