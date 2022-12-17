package repository

import (
	"context"

	"github.com/storyofhis/xtrame/httpserver/repository/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.Users) error
	FindUserByEmail(ctx context.Context, email string) (*models.Users, error)
	FindUserById(ctx context.Context, id uint) (*models.Users, error)
}

type TicketRepo interface {
	CreateTicket(ctx context.Context, ticket *models.Tickets) error
	GetTickets(ctx context.Context) ([]models.Tickets, error)
	GetTicketById(ctx context.Context, id uint) (*models.Tickets, error)
	UpdateTicket(ctx context.Context, ticket *models.Tickets, id uint) error
	DeleteTicket(ctx context.Context, ticket *models.Tickets, id uint) error
}

type CategoryRepo interface {
	CreateCategory(ctx context.Context, category *models.Category) error
	UpdateCategory(ctx context.Context, category *models.Category, id uint) error
	GetCategories(ctx context.Context) ([]models.Category, error)
	GetCategoryById(ctx context.Context, id uint) (*models.Category, error)
	DeleteCategory(ctx context.Context, id uint) error
}
