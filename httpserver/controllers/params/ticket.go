package params

import "time"

type CreateTicket struct {
	Title       string        `json:"title" validate:"required"`
	Description string        `json:"description" validate:"required"`
	Price       float32       `json:"price" validate:"required"`
	Seat        int           `json:"seat" validate:"required"`
	CategoryId  uint          `json:"category_id" validate:"required"`
	Duration    time.Duration `json:"duration"`
}

type UpdateTicket struct {
	Title       string        `json:"title" validate:"required"`
	Description string        `json:"description"`
	Price       float32       `json:"price" validate:"required"`
	Seat        int           `json:"seat" validate:"required"`
	CategoryId  uint          `json:"category_id" validate:"required"`
	Duration    time.Duration `json:"duration"`
}
