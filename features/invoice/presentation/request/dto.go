package request

import (
	"watchamovie-payment/features/invoice"
)

type ReqInvoice struct {
	UserId       uint   `json:"user_id"`
	Item         string `json:"item"`
	Total        int    `json:"total"`
	PaymentTerms int    `json:"payment_terms"`
	PaymentLink  string `json:"payment_link"`
}

type ReqInvoiceUpdate struct {
	Id            uint   `json:"id"`
	UserId        uint   `json:"user_id"`
	Item          string `json:"item"`
	Total         int    `json:"total"`
	PaymentTerms  int    `json:"payment_terms"`
	PaymentStatus string `json:"payment_status"`
	PaymentLink   string `json:"payment_link"`
}

func (reqdata *ReqInvoice) ToInvoiceCore() invoice.InvoiceCore {
	return invoice.InvoiceCore{
		UserId:       reqdata.UserId,
		Item:         reqdata.Item,
		Total:        reqdata.Total,
		PaymentTerms: reqdata.PaymentTerms,
		PaymentLink:  reqdata.PaymentLink,
	}
}

func (reqdata *ReqInvoiceUpdate) ToInvoiceCore() invoice.InvoiceCore {
	return invoice.InvoiceCore{
		Id:            reqdata.Id,
		UserId:        reqdata.UserId,
		Item:          reqdata.Item,
		Total:         reqdata.Total,
		PaymentTerms:  reqdata.PaymentTerms,
		PaymentStatus: reqdata.PaymentStatus,
		PaymentLink:   reqdata.PaymentLink,
	}
}
