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
