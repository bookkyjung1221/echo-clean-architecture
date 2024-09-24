package model

import "time"

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}

type TaskResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskResponseDetail struct {
	TaskResponse []TaskResponse
	Code         int    `json:"code"`
	Message      string `json:"message"`
}

type TaskResponseOneDetail struct {
	TaskResponse TaskResponse
	Code         int    `json:"code"`
	Message      string `json:"message"`
}

type TaskResponseDelete struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type TaskResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
