package routes

import (
	"watchamovie-payment/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	presenter := factory.Init()
	// Initiate Echo & JWT
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	// Invoice
	e.POST("/invoice/add", presenter.InvoicePresentation.CreateInvoiceHandler)
	e.POST("/transactions/callback", presenter.InvoicePresentation.CallbackHandler)
	return e
}
