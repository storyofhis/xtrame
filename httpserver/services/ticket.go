package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/storyofhis/xtrame/httpserver/controllers/params"
	"github.com/storyofhis/xtrame/httpserver/controllers/views"
	"github.com/storyofhis/xtrame/httpserver/repository"
	"github.com/storyofhis/xtrame/httpserver/repository/models"
	"gorm.io/gorm"
)

type ticketSvc struct {
	repos        repository.TicketRepo
	userRepo     repository.UserRepo
	categoryRepo repository.CategoryRepo
}

func NewTicketSvc(repos repository.TicketRepo, userRepo repository.UserRepo, categoryRepo repository.CategoryRepo) TicketSvc {
	return &ticketSvc{
		repos:        repos,
		userRepo:     userRepo,
		categoryRepo: categoryRepo,
	}
}

func (svc *ticketSvc) CreateTicket(ctx context.Context, ticket *params.CreateTicket) *views.Response {
	if _, err := svc.categoryRepo.GetCategoryById(ctx, ticket.CategoryId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, errors.New("category id is invalid"))
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	// request
	model := models.Tickets{
		Title:       ticket.Title,
		Description: ticket.Description,
		Price:       ticket.Price,
		Seat:        ticket.Seat,
		CategoryId:  ticket.CategoryId,
		// Duration:    ticket.Duration,
	}

	err := svc.repos.CreateTicket(ctx, &model)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	// response
	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.CreateTicket{
		Id:          model.Id,
		Title:       model.Title,
		Description: model.Description,
		Price:       model.Price,
		Seat:        model.Seat,
		Duration:    model.Duration,
		CategoryId:  model.CategoryId,
		CreatedAt:   model.CreatedAt,
	})
}

func (svc *ticketSvc) UpdateTicket(ctx context.Context, ticket *params.UpdateTicket, id uint) *views.Response {
	_, err := svc.categoryRepo.GetCategoryById(ctx, ticket.CategoryId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, errors.New("category id is invalid"))
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	model, err := svc.repos.GetTicketById(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, errors.New("category id is invalid"))
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	/** request */
	model.Title = ticket.Title
	model.Description = ticket.Description
	model.Price = ticket.Price
	model.Seat = ticket.Seat
	model.CategoryId = ticket.CategoryId
	model.Duration = ticket.Duration

	err = svc.repos.UpdateTicket(ctx, model, id)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_OK, views.UpdateTicket{
		Id:          model.Id,
		Title:       model.Title,
		Description: model.Description,
		Price:       model.Price,
		Seat:        model.Seat,
		Duration:    model.Duration,
		CategoryId:  model.CategoryId,
		UpdatedAt:   model.UpdatedAt,
	})
}

func (svc *ticketSvc) GetTickets(ctx context.Context) *views.Response {
	tickes, err := svc.repos.GetTickets(ctx)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	response := make([]views.GetTickets, 0)
	for i := 0; i < len(tickes); i++ {
		response = append(response, views.GetTickets{
			Id:          tickes[i].Id,
			Title:       tickes[i].Title,
			Description: tickes[i].Description,
			Price:       tickes[i].Price,
			Seat:        tickes[i].Seat,
			Duration:    tickes[i].Duration,
			CategoryId:  tickes[i].CategoryId,
			CreatedAt:   tickes[i].CreatedAt,
			UpdatedAt:   tickes[i].UpdatedAt,
		})
	}

	return views.SuccessResponse(http.StatusOK, views.M_OK, response)
}

func (svc *ticketSvc) GetTicketById(ctx context.Context, id uint) *views.Response {
	ticket, err := svc.repos.GetTicketById(ctx, id)
	// user, err := svc.userRepo.FindUserById(ctx, ticket.UserId)
	category, err := svc.categoryRepo.GetCategoryById(ctx, ticket.CategoryId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, errors.New("category id is invalid"))
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	// response
	return views.SuccessResponse(http.StatusOK, views.M_OK, views.GetTicketById{
		Id:          ticket.Id,
		Title:       ticket.Title,
		Description: ticket.Description,
		Price:       ticket.Price,
		Seat:        ticket.Seat,
		Duration:    ticket.Duration,
		CategoryId:  category.Id,
		// UserId:      user.Id,
		// UserTickets: views.UserTickets{
		// 	Id:       user.Id,
		// 	FullName: user.FullName,
		// 	NickName: user.NickName,
		// 	UserName: user.UserName,
		// 	Email:    user.Email,
		// 	Age:      user.Age,
		// 	Role:     user.Role,
		// },
	})
}

func (svc *ticketSvc) DeleteTicket(ctx context.Context, id uint) *views.Response {
	ticket, err := svc.repos.GetTicketById(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, errors.New("product with this id is not found"))
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	err = svc.repos.DeleteTicket(ctx, ticket, id)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_TICKET_SUCCESSFULLY_DELETED, nil)
}
