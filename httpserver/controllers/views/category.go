package views

import "time"

type CreateCategory struct {
	Id        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateCategory struct {
	Id        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetCategories struct {
	Id        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
