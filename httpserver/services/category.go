package services

import (
	"context"
	"net/http"

	"github.com/storyofhis/xtrame/httpserver/controllers/params"
	"github.com/storyofhis/xtrame/httpserver/controllers/views"
	"github.com/storyofhis/xtrame/httpserver/repository"
	"github.com/storyofhis/xtrame/httpserver/repository/models"
)

type categorySvc struct {
	repo repository.CategoryRepo
	user repository.UserRepo
}

func NewCategorySvc(repo repository.CategoryRepo, user repository.UserRepo) CategorySvc {
	return &categorySvc{
		repo: repo,
		user: user,
	}
}

func (svc *categorySvc) CreateCategory(ctx context.Context, category *params.CreateCategory) *views.Response {
	model := models.Category{
		Type: category.Type,
	}
	if err := svc.repo.CreateCategory(ctx, &model); err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.CreateCategory{
		Id:        model.Id,
		UserId:    model.UserId,
		Type:      model.Type,
		CreatedAt: model.CreatedAt,
	})
}

func (svc *categorySvc) UpdateCategory(ctx context.Context, category *params.UpdateCategory, id uint) *views.Response {
	model := models.Category{
		Type: category.Type,
	}
	// err := svc.repo.UpdateCategory(ctx, &model, model.Id)
	err := svc.repo.UpdateCategory(ctx, &model, id)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
	return views.SuccessResponse(http.StatusOK, views.M_OK, views.UpdateCategory{
		Id:        model.Id,
		UserId:    model.UserId,
		Type:      model.Type,
		UpdatedAt: model.UpdatedAt,
	})
}

func (svc *categorySvc) GetCategories(ctx context.Context) *views.Response {

	return nil
}
