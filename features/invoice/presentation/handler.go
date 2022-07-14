package presentation

import (
	"net/http"
	"strconv"
	"watchamovie-payment/features/invoice"
	"watchamovie-payment/features/invoice/presentation/request"
	"watchamovie-payment/helper"

	"github.com/labstack/echo/v4"
)

type InvoiceHandler struct {
	invoiceBusiness invoice.Business
}

func NewHandlerInvoice(invoiceBusiness invoice.Business) *InvoiceHandler {
	return &InvoiceHandler{invoiceBusiness}
}

func (inHandler *InvoiceHandler) CreateInvoiceHandler(e echo.Context) error {
	newInvoice := request.ReqInvoice{}

	if err := e.Bind(&newInvoice); err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	data, err := inHandler.invoiceBusiness.CreateInvoice(newInvoice.ToInvoiceCore())
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}

	return helper.SuccessResponse(e, data)
}
func (inHandler *InvoiceHandler) CallbackHandler(e echo.Context) error {
	var body = map[string]interface{}{}
	err := e.Bind(&body)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	orderID, exist := body["order_id"].(string)
	if !exist {
		return helper.ErrorResponse(e, http.StatusBadRequest, "id does not exist", err)
	}

	transactionID, err := strconv.Atoi(orderID)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "id not valid", err)
	}

	err = inHandler.invoiceBusiness.UpdateTransactionStatus(int64(transactionID))
	if err != nil {
		return err
	}
	return helper.SuccessResponse(e, nil)
}
