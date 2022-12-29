package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tryhd/dbo-test/app/helper"
	"github.com/tryhd/dbo-test/app/models"
	"github.com/tryhd/dbo-test/app/types"
)

type OrderController struct {
	m models.OrderModel
}

func NewOrderController(m models.OrderModel) *OrderController {
	return &OrderController{
		m: m,
	}
}

func (c *OrderController) RegisterOrder(context *gin.Context) {
	var req types.Order

	err := context.ShouldBind(&req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	ok, err := c.m.RegisterOrder(req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", ok)
	context.JSON(http.StatusOK, res)
}

func (c *OrderController) DetailOrder(context *gin.Context) {

	ok, err := c.m.DetailOrder(context.Param("id"))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", ok)
	context.JSON(http.StatusOK, res)
}

func (c *OrderController) DeleteOrder(context *gin.Context) {

	ok, err := c.m.DeleteOrder(context.Param("id"))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", ok)
	context.JSON(http.StatusOK, res)
}

func (c *OrderController) UpdateOrder(context *gin.Context) {

	var req types.OrderUpdate

	err := context.ShouldBind(&req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	req.ID = context.Param("id")

	ok, err := c.m.UpdateOrder(req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", ok)
	context.JSON(http.StatusOK, res)
}

func (c *OrderController) GetAllOrder(context *gin.Context) {
	limit := 2
	page := 1
	sort := "created_at asc"
	query := context.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break

		}
	}
	ok, err := c.m.GetAllOrder(limit, page, sort)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", ok)
	context.JSON(http.StatusOK, res)
}

func (c *OrderController) FindOrder(context *gin.Context) {
	var search string
	query := context.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "search":
			if len(queryValue) < 3 {
				res := helper.BuildErrorResponse("Failed to process request", "Must have more than 3 character", helper.EmptyObj{})
				context.JSON(http.StatusBadRequest, res)
				break
			}
			search = queryValue
			break
		}
	}
	ok, err := c.m.FindOrder(search)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", ok)
	context.JSON(http.StatusOK, res)
}
