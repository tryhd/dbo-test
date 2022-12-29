package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tryhd/dbo-test/app/helper"
	"github.com/tryhd/dbo-test/app/models"
	"github.com/tryhd/dbo-test/app/types"
)

type CustomerController struct {
	m models.CustomerModel
}

func NewCustomerController(m models.CustomerModel) *CustomerController {
	return &CustomerController{
		m: m,
	}
}

func (c *CustomerController) RegisterCustomer(context *gin.Context) {
	var req types.Customer

	err := context.ShouldBind(&req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	ok, err := c.m.RegisterCustomer(req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", ok)
	context.JSON(http.StatusOK, res)
}

func (c *CustomerController) DetailCustomer(context *gin.Context) {

	ok, err := c.m.DetailCustomer(context.Param("id"))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", ok)
	context.JSON(http.StatusOK, res)
}

func (c *CustomerController) DeleteCustomer(context *gin.Context) {

	ok, err := c.m.DeleteCustomer(context.Param("id"))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", ok)
	context.JSON(http.StatusOK, res)
}

func (c *CustomerController) UpdateCustomer(context *gin.Context) {

	var req types.CustomerUpdate

	err := context.ShouldBind(&req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	req.ID = context.Param("id")

	ok, err := c.m.UpdateCustomer(req)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", ok)
	context.JSON(http.StatusOK, res)
}

func (c *CustomerController) GetAllCustomer(context *gin.Context) {
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
	ok, err := c.m.GetAllCustomer(limit, page, sort)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", ok)
	context.JSON(http.StatusOK, res)
}

func (c *CustomerController) FindCustomer(context *gin.Context) {
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
	ok, err := c.m.FindCustomer(search)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", ok)
	context.JSON(http.StatusOK, res)
}
