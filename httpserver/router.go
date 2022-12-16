package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/storyofhis/xtrame/httpserver/controllers"
)

type router struct {
	router *gin.Engine

	user *controllers.UserController
}

func NewRouter(r *gin.Engine, user *controllers.UserController) *router {
	return &router{
		router: r,

		user: user,
	}
}

func (r *router) Start(port string) {
	// user
	r.router.POST("/v1/register", r.user.Register)
	r.router.POST("/v1/login", r.user.Login)

	// run
	r.router.Run(port)
}
