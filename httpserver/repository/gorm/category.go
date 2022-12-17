package gorm

import (
	"context"
	"time"

	"github.com/storyofhis/xtrame/httpserver/repository"
	"github.com/storyofhis/xtrame/httpserver/repository/models"
	"gorm.io/gorm"
)

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) repository.CategoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (repo *categoryRepo) CreateCategory(ctx context.Context, category *models.Category) error {
	category.CreatedAt = time.Now()
	if err := repo.db.WithContext(ctx).Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (repo *categoryRepo) UpdateCategory(ctx context.Context, category *models.Category, id uint) error {
	category.UpdatedAt = time.Now()
	if err := repo.db.WithContext(ctx).Where("id = ?", id).Updates(category).Error; err != nil {
		return err
	}
	return nil
}

func (repo *categoryRepo) GetCategories(ctx context.Context) ([]models.Category, error) {
	var category []models.Category
	if err := repo.db.WithContext(ctx).Find(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (repo *categoryRepo) GetCategoryById(ctx context.Context, id uint) (*models.Category, error) {
	var category models.Category
	if err := repo.db.WithContext(ctx).Where("id = ?", id).Updates(category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (repo *categoryRepo) DeleteCategory(ctx context.Context, id uint) error {
	var category models.Category
	category.DeletedAt = time.Now()
	if err := repo.db.WithContext(ctx).Where("id = ?", id).Delete(category).Error; err != nil {
		return err
	}
	return nil
}
