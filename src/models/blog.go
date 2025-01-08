package models

import "time"

type BlogStructModel struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	UserCreated string    `json:"user_created"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
