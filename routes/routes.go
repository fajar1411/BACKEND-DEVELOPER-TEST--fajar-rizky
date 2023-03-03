package routes

import (
	customer "test/fitur/customer"
	handlercostumer "test/fitur/customer/handler"
	family "test/fitur/family"
	handlerfamily "test/fitur/family/handler"
	"test/fitur/national"
	handlernational "test/fitur/national/handler"
	"test/middlewares"

	"github.com/labstack/echo/v4"
)

func NewHandlerCostumer(Service customer.CostumerService, e *echo.Echo) {
	handlers := &handlercostumer.CustomerHandler{
		CostumerServices: Service,
	}

	e.POST("/costumer/form", handlers.FormData)
	e.POST("/costumer/login", handlers.Login)
	e.GET("/costumer/profile", handlers.Profile, middlewares.JWTMiddleware())
	e.PUT("/costumer", handlers.Update, middlewares.JWTMiddleware())
	e.DELETE("/costumer", handlers.Delete, middlewares.JWTMiddleware())
}
func NewHandlerFamily(Service family.FamilyService, e *echo.Echo) {
	handlers := &handlerfamily.FamilyHandler{
		FamilyServices: Service,
	}

	e.POST("/costumer/family", handlers.AddFamily, middlewares.JWTMiddleware())
	e.GET("/costumer/family", handlers.MyFamily, middlewares.JWTMiddleware())
	e.PUT("/costumer/family/:id", handlers.UpdateFamily, middlewares.JWTMiddleware())
	e.DELETE("/costumer/family/:id", handlers.DeleteFamily, middlewares.JWTMiddleware())
}
func NewHandlerNational(Service national.NationalService, e *echo.Echo) {
	handlers := &handlernational.NationalHandler{
		NationalServices: Service,
	}

	e.POST("/national", handlers.AddNational)

}
