package helper

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseReponse struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		// Description []string `json:"description"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

// Valid Response
func SuccessResponse(e echo.Context, data interface{}) error {
	response := BaseReponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "successful operator"
	response.Data = data
	return e.JSON(http.StatusOK, response)
}

// Invalid Response
func ErrorResponse(c echo.Context, status int, err string, errs error) error {
	response := BaseReponse{}
	response.Meta.Status = status
	// if errs != nil {
	// 	response.Meta.Description = []string{errs.Error()}
	// }
	response.Meta.Message = err
	return c.JSON(status, response)
}
