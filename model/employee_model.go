package model

import "time"

type Employee struct {
	EmployeeID int    `json:"employee_id"`
	AgentID    int    `json:"agent_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedLoc string `json:"created_loc"`
	CreatedBy string `json:"created_by"`
}

type NewEmployee struct {
	AgentID    int    `json:"agent_id" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone" validate:"required,e164"` 
	CreatedLoc string `json:"created_loc"`
	CreatedBy string `json:"created_by" validate:"required"`
}

type UpdateEmployee struct {
	AgentID    int    `json:"agent_id" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone" validate:"required,e164"`
}