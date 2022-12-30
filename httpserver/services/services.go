package services

import (
	"context"

	"github.com/storyofhis/xtrame/httpserver/controllers/params"
	"github.com/storyofhis/xtrame/httpserver/controllers/views"
)

type EmailSvc interface {
	ConnectEmail() *views.Response
}

type UserSvc interface {
	Register(ctx context.Context, user *params.Register) *views.Response
	Login(ctx context.Context, user *params.Login) *views.Response
}

type TicketSvc interface {
	CreateTicket(ctx context.Context, ticket *params.CreateTicket) *views.Response
	UpdateTicket(ctx context.Context, ticket *params.UpdateTicket, id uint) *views.Response
	GetTickets(ctx context.Context) *views.Response
	GetTicketById(ctx context.Context, id uint) *views.Response
	DeleteTicket(ctx context.Context, id uint) *views.Response
}

type CategorySvc interface {
	CreateCategory(ctx context.Context, category *params.CreateCategory) *views.Response
	UpdateCategory(ctx context.Context, category *params.UpdateCategory, id uint) *views.Response
	GetCategories(ctx context.Context) *views.Response
}
