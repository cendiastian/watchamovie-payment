package invoice

import "time"

type InvoiceCore struct {
	Id            uint
	UserId        string
	FullName      string
	Email         string
	Item          string
	Total         int
	PaymentDue    time.Time
	Expired       int
	PaymentStatus string
	PaymentTerms  int
	PaymentLink   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Business interface {
	CreateInvoice(data InvoiceCore) (InvoiceCore, error)
	GetInvoice(id int, UserId string) (InvoiceCore, error)
	UpdateTransactionStatus(transactionID int64) error
}

type Data interface {
	CreateInvoice(data InvoiceCore) (invoice InvoiceCore, err error)
	UpdateInvoice(data InvoiceCore) error
	GetInvoice(id int, UserId string) (InvoiceCore, error)
	UpdateTransactionStatus(transactionID int64, PaymentStatus string) error
}
