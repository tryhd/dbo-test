package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tryhd/dbo-test/app/config/middleware"
	"github.com/tryhd/dbo-test/app/controllers"
	"github.com/tryhd/dbo-test/app/models"
)

func CustomerRoute(route *gin.Engine) {
	m := models.NewCustomerModels()
	m.Init()
	c := controllers.NewCustomerController(m)
	authRoutes := route.Group("api/v1/customer")
	{
		authRoutes.Use(middleware.JwtAuthMiddleware())
		authRoutes.POST("/create", c.RegisterCustomer)
		authRoutes.GET("/all", c.GetAllCustomer)
		authRoutes.GET("/:id", c.DetailCustomer)
		authRoutes.GET("/", c.FindCustomer)
		authRoutes.PUT("/:id", c.UpdateCustomer)
		authRoutes.DELETE("/:id", c.DeleteCustomer)

	}
}
