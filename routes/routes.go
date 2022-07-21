package routes

import (
	"net/http"
	"watchamovie-payment/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	presenter := factory.Init()
	// Initiate Echo & JWT
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// AllowOrigins: []string{"https://watchamovie.herokuapp.com", "http://localhost:3000"},
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	// Invoice
	e.POST("/invoice/add", presenter.InvoicePresentation.CreateInvoiceHandler)
	e.POST("/transactions/callback", presenter.InvoicePresentation.CallbackHandler)
	e.GET("/invoice", presenter.InvoicePresentation.GetInvoice)
	return e
}
