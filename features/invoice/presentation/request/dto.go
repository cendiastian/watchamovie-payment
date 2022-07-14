package request

import (
	"watchamovie-payment/features/invoice"
)

type ReqInvoice struct {
	UserId       uint   `json:"user_id"`
	Item         string `json:"item"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	Total        int    `json:"total"`
	PaymentTerms int    `json:"payment_terms"`
	PaymentLink  string `json:"payment_link"`
}

func (reqdata *ReqInvoice) ToInvoiceCore() invoice.InvoiceCore {
	return invoice.InvoiceCore{
		UserId:       reqdata.UserId,
		Item:         reqdata.Item,
		Total:        reqdata.Total,
		FullName:     reqdata.FullName,
		Email:        reqdata.Email,
		PaymentTerms: reqdata.PaymentTerms,
		PaymentLink:  reqdata.PaymentLink,
	}
}
