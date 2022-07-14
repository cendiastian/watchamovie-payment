package invoice

import "time"

type InvoiceCore struct {
	Id            uint
	UserId        uint
	FullName      string
	Email         string
	Item          string
	Total         int
	PaymentDue    time.Time
	PaymentStatus string
	PaymentTerms  int
	PaymentLink   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Business interface {
	CreateInvoice(data InvoiceCore) (InvoiceCore, error)
	UpdateTransactionStatus(transactionID int64) error
}

type Data interface {
	CreateInvoice(data InvoiceCore) (invoice InvoiceCore, err error)
	UpdateInvoice(data InvoiceCore) error
	UpdateTransactionStatus(transactionID int64, PaymentStatus string) error
}
