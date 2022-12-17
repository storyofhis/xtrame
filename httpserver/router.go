package httpserver

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/storyofhis/xtrame/common"
	"github.com/storyofhis/xtrame/httpserver/controllers"
)

type router struct {
	router *gin.Engine

	user     *controllers.UserController
	category *controllers.CategoryController
	ticket   *controllers.TicketController
}

func NewRouter(r *gin.Engine, user *controllers.UserController, category *controllers.CategoryController, ticket *controllers.TicketController) *router {
	return &router{
		router: r,

		user:     user,
		category: category,
		ticket:   ticket,
	}
}

func (r *router) Start(port string) {
	// user
	r.router.POST("/v1/register", r.user.Register)
	r.router.POST("/v1/login", r.user.Login)

	// category
	r.router.POST("/v1/category", r.verifyToken, r.category.CreateCategory)
	r.router.PUT("/v1/category/:id", r.verifyToken, r.category.UpdateCategory)

	// ticket
	r.router.POST("/v1/ticket", r.verifyToken, r.ticket.CreateTicket)
	r.router.GET("/v1/ticket", r.verifyToken, r.ticket.GetTickets)
	r.router.GET("/v1/ticket/:id", r.verifyToken, r.ticket.GetTicketById)
	r.router.PUT("/v1/ticket/:id", r.verifyToken, r.ticket.UpdateTicket)
	r.router.DELETE("/v1/ticket/:id", r.verifyToken, r.ticket.DeleteTickets)

	// run
	r.router.Run(port)
}

func (r *router) verifyToken(ctx *gin.Context) {
	bearerToken := strings.Split(ctx.Request.Header.Get("Authorization"), "Bearer ")
	if len(bearerToken) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid bearer token",
		})
		return
	}
	claims, err := common.ValidateToken(bearerToken[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Set("userData", claims)
}
