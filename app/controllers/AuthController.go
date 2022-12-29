package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tryhd/dbo-test/app/helper"
	"github.com/tryhd/dbo-test/app/models"
	"github.com/tryhd/dbo-test/app/types"
)

type AuthController struct {
	m models.AuthModel
}

func NewAuthController(m models.AuthModel) *AuthController {
	return &AuthController{
		m: m,
	}
}

func (c *AuthController) Login(context *gin.Context) {
	var req types.LoginRequest

	err := context.ShouldBind(&req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	login, err := c.m.Login(req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", login)
	context.JSON(http.StatusOK, res)
}

func (c *AuthController) Register(context *gin.Context) {
	var req types.RegisterRequest

	err := context.ShouldBind(&req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	register, err := c.m.RegisterAuth(req)

	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	if register {
		reqLogin := types.LoginRequest{
			Username: req.Username,
			Password: req.Password,
		}
		login, err := c.m.Login(reqLogin)
		if err != nil {
			res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
			context.JSON(http.StatusBadRequest, res)
			return
		}
		res := helper.BuildResponse(true, "OK", login)
		context.JSON(http.StatusOK, res)
		return
	}
}
