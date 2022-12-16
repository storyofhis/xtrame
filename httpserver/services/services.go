package services

import (
	"context"

	"github.com/storyofhis/xtrame/httpserver/controllers/params"
	"github.com/storyofhis/xtrame/httpserver/controllers/views"
)

type UserSvc interface {
	Register(ctx context.Context, user *params.Register) *views.Response
	Login(ctx context.Context, user *params.Login) *views.Response
}
