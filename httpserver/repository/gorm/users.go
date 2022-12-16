package gorm

import (
	"context"
	"strings"
	"time"

	"github.com/storyofhis/xtrame/httpserver/repository"
	"github.com/storyofhis/xtrame/httpserver/repository/models"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repository.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (repo *userRepo) CreateUser(ctx context.Context, user *models.Users) error {
	user.CreatedAt = time.Now()
	err := repo.db.WithContext(ctx).Create(user).Error
	return err
}

func (repo *userRepo) FindUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	user := new(models.Users)
	err := repo.db.WithContext(ctx).Where("LOWER(email) = ?", strings.ToLower(email)).Take(user).Error
	return user, err
}

func (repo *userRepo) FindUserById(ctx context.Context, id uint) (*models.Users, error) {
	user := new(models.Users)
	err := repo.db.WithContext(ctx).Where("id = ?", id).Take(user).Error
	return user, err
}
