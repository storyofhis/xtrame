package views

import (
	"time"

	"github.com/storyofhis/xtrame/httpserver/repository/models"
)

type CreateTicket struct {
	Id          uint          `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Price       float32       `json:"price"`
	Seat        int           `json:"seat"`
	Duration    time.Duration `json:"duration"`
	CategoryId  uint          `json:"category_id"`
	CreatedAt   time.Time     `json:"created_at"`
}

type GetTickets struct {
	Id          uint          `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Price       float32       `json:"price"`
	Seat        int           `json:"seat"`
	Duration    time.Duration `json:"duration"`
	CategoryId  uint          `json:"category_id"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type GetTicketById struct {
	Id          uint          `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Price       float32       `json:"price"`
	Seat        int           `json:"seat"`
	Duration    time.Duration `json:"duration"`
	CategoryId  uint          `json:"category_id"`
	UserId      uint          `json:"user_id"`
	UserTickets UserTickets   `json:"user"`
}

type UserTickets struct {
	Id       uint        `json:"id"`
	FullName string      `json:"full_name"`
	NickName string      `json:"nick_name"`
	UserName string      `json:"username"`
	Email    string      `json:"email"`
	Age      uint        `json:"age"`
	Role     models.Role `json:"role"`
}

type UpdateTicket struct {
	Id          uint          `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Price       float32       `json:"price"`
	Seat        int           `json:"seat"`
	Duration    time.Duration `json:"duration"`
	CategoryId  uint          `json:"category_id"`
	UpdatedAt   time.Time     `json:"updated_at"`
}
