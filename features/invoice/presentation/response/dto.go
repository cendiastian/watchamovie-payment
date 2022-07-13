package response

import (
	"time"
	"watchamovie-payment/features/invoice"
)

type RespInvoice struct {
	ID            uint      `json:"id"`
	UserId        uint      `json:"user_id"`
	Item          string    `json:"item"`
	Total         int       `json:"total"`
	PaymentDue    time.Time `json:"payment_due"`
	PaymentStatus string    `json:"payment_status"`
	PaymentTerms  int       `json:"payment_terms"`
	PaymentLink   string    `json:"payment_link"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func ToInvoiceResponse(in invoice.InvoiceCore) RespInvoice {
	return RespInvoice{
		ID:            in.Id,
		UserId:        in.UserId,
		Item:          in.Item,
		Total:         in.Total,
		PaymentDue:    in.PaymentDue,
		PaymentStatus: in.PaymentStatus,
		PaymentTerms:  in.PaymentTerms,
		PaymentLink:   in.PaymentLink,
		CreatedAt:     in.CreatedAt,
		UpdatedAt:     in.UpdatedAt,
	}
}

func ToInvoiceResponseList(inList []invoice.InvoiceCore) []RespInvoice {
	convIn := []RespInvoice{}

	for _, invoice := range inList {
		convIn = append(convIn, ToInvoiceResponse(invoice))
	}
	return convIn
}
