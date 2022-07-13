package business

import (
	"errors"
	"fmt"
	"time"
	"watchamovie-payment/features/invoice"
	"watchamovie-payment/helper"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type InvoiceBusiness struct {
	invoiceData        invoice.Data
	midtransClient     snap.Client
	midtransCoreClient coreapi.Client
}

func NewBusinessInvoice(inData invoice.Data, midtransClient snap.Client, midtransCoreClient coreapi.Client) invoice.Business {
	return &InvoiceBusiness{inData, midtransClient, midtransCoreClient}
}

func (inBusiness *InvoiceBusiness) CreateInvoice(dataReq invoice.InvoiceCore) (invoice.InvoiceCore, error) {
	t := time.Now()
	if dataReq.PaymentTerms == 7 {
		dataReq.PaymentDue = t.Add(time.Hour * 24 * 7)
	} else if dataReq.PaymentTerms == 10 {
		dataReq.PaymentDue = t.Add(time.Hour * 24 * 10)
	} else if dataReq.PaymentTerms == 30 {
		dataReq.PaymentDue = t.Add(time.Hour * 24 * 30)
	}

	data, err := inBusiness.invoiceData.CreateInvoice(dataReq)
	if err != nil {
		return invoice.InvoiceCore{}, err
	}
	if data.PaymentStatus == "unpaid" {
		resp, errMidtrans := inBusiness.midtransClient.CreateTransaction(&snap.Request{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  fmt.Sprintf("%d", data.Id),
				GrossAmt: int64(data.Total),
			},
			Expiry: &snap.ExpiryDetails{
				Unit:     "day",
				Duration: int64(data.PaymentTerms),
			},
			CreditCard: &snap.CreditCardDetails{
				Secure: true,
			},
		})

		if errMidtrans != nil {
			return invoice.InvoiceCore{}, errMidtrans
		}
		data.PaymentLink = resp.RedirectURL
	}

	err = inBusiness.invoiceData.UpdateInvoice(data)
	if err != nil {
		return invoice.InvoiceCore{}, err
	}

	return data, nil
}
func (inBusiness *InvoiceBusiness) UpdateInvoice(data invoice.InvoiceCore) (invoice.InvoiceCore, error) {
	if helper.IsEmpty(data.PaymentStatus) {
		return invoice.InvoiceCore{}, errors.New("invalid data")
	}

	t := time.Now()
	if data.PaymentTerms == 7 {
		data.PaymentDue = t.Add(time.Hour * 24 * 7)
	} else if data.PaymentTerms == 10 {
		data.PaymentDue = t.Add(time.Hour * 24 * 10)
	} else if data.PaymentTerms == 30 {
		data.PaymentDue = t.Add(time.Hour * 24 * 30)
	}

	if data.PaymentStatus == "unpaid" {
		resp, errMidtrans := inBusiness.midtransClient.CreateTransaction(&snap.Request{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  fmt.Sprintf("%d", data.Id),
				GrossAmt: int64(data.Total),
			},
			Expiry: &snap.ExpiryDetails{
				Unit:     "day",
				Duration: int64(data.PaymentTerms),
			},
			CreditCard: &snap.CreditCardDetails{
				Secure: true,
			},
		})

		if errMidtrans != nil {
			return invoice.InvoiceCore{}, errMidtrans
		}
		data.PaymentLink = resp.RedirectURL
	}

	err := inBusiness.invoiceData.UpdateInvoice(data)
	if err != nil {
		return invoice.InvoiceCore{}, err
	}
	return data, nil
}

func (inBusiness *InvoiceBusiness) UpdateTransactionStatus(transactionID int64) error {
	trans, err := inBusiness.midtransCoreClient.CheckTransaction(fmt.Sprintf("%d", transactionID))
	if err != nil {
		return err
	}

	if trans.TransactionStatus == "capture" || trans.TransactionStatus == "settlement" {
		PaymentStatus := "paid"
		errUpd := inBusiness.invoiceData.UpdateTransactionStatus(transactionID, PaymentStatus)
		if errUpd != nil {
			return errUpd
		}
	}

	return nil
}
