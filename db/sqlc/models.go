// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Task struct {
	ID int64 `json:"id"`
	// unique per owner
	Name            string    `json:"name"`
	Owner           string    `json:"owner"`
	Content         string    `json:"content"`
	Deleted         bool      `json:"deleted"`
	ContentChangeAt time.Time `json:"content_change_at"`
	DeletedAt       time.Time `json:"deleted_at"`
	CreatedAt       time.Time `json:"created_at"`
}

type User struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	// non negative
	DailyCap int64 `json:"daily_cap"`
	// non negative
	DailyQuantity    int64     `json:"daily_quantity"`
	PasswordChangeAt time.Time `json:"password_change_at"`
	CreatedAt        time.Time `json:"created_at"`
}
