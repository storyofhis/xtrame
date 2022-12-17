package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/storyofhis/xtrame/common"
	"github.com/storyofhis/xtrame/httpserver/controllers/params"
	"github.com/storyofhis/xtrame/httpserver/services"
)

type CategoryController struct {
	svc services.CategorySvc
}

func NewCategoryController(svc services.CategorySvc) *CategoryController {
	return &CategoryController{
		svc: svc,
	}
}

func (control *CategoryController) CreateCategory(ctx *gin.Context) {
	var req params.CreateCategory
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	claims, exists := ctx.Get("userData")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token doesn't exists",
		})
		return
	}
	userData := claims.(*common.CustomClaims)
	userRole := userData.Role
	if userRole != "admin" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}
	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := control.svc.CreateCategory(ctx, &req)
	WriteJsonResponse(ctx, response)
}

func (control *CategoryController) UpdateCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var req params.UpdateCategory
	err = ctx.ShouldBind(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	claims, exist := ctx.Get("userData")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token doesn't exists",
		})
		return
	}

	userData := claims.(*common.CustomClaims)
	userRole := userData.Role
	if userRole != "admin" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := control.svc.UpdateCategory(ctx, &req, uint(id))
	WriteJsonResponse(ctx, response)
}
