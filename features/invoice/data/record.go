package data

import (
	"fmt"
	"time"
	"watchamovie-payment/features/invoice"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	UserId        uint
	FullName      string
	Email         string
	Item          string
	Total         int
	PaymentDue    time.Time
	PaymentStatus string `gorm:"default:unpaid"`
	PaymentTerms  int
	PaymentLink   string
}

func toInvoiceRecord(in invoice.InvoiceCore) Invoice {
	fmt.Println("Isi payment due di record : ", in.PaymentDue)
	return Invoice{
		Model: gorm.Model{
			ID:        in.Id,
			CreatedAt: in.CreatedAt,
			UpdatedAt: in.UpdatedAt,
		},
		UserId:        in.UserId,
		Item:          in.Item,
		FullName:      in.FullName,
		Email:         in.Email,
		Total:         in.Total,
		PaymentDue:    in.PaymentDue,
		PaymentStatus: in.PaymentStatus,
		PaymentTerms:  in.PaymentTerms,
		PaymentLink:   in.PaymentLink,
	}
}

func toInvoiceCore(in Invoice) invoice.InvoiceCore {
	return invoice.InvoiceCore{
		Id:            in.ID,
		UserId:        in.UserId,
		Item:          in.Item,
		Total:         in.Total,
		FullName:      in.FullName,
		Email:         in.Email,
		PaymentDue:    in.PaymentDue,
		PaymentStatus: in.PaymentStatus,
		PaymentTerms:  in.PaymentTerms,
		CreatedAt:     in.CreatedAt,
		UpdatedAt:     in.UpdatedAt,
		PaymentLink:   in.PaymentLink,
	}
}

func toInvoiceCoreList(inList []Invoice) []invoice.InvoiceCore {
	convIn := []invoice.InvoiceCore{}

	for _, invoice := range inList {
		convIn = append(convIn, toInvoiceCore(invoice))
	}
	return convIn
}
