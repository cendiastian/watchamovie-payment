package business

import (
	"fmt"
	"time"
	"watchamovie-payment/features/invoice"

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
	fmt.Println(dataReq.PaymentTerms)
	if dataReq.PaymentTerms == 7 {
		dataReq.PaymentDue = t.Add(time.Hour * 24 * 7)
	} else if dataReq.PaymentTerms == 10 {
		dataReq.PaymentDue = t.Add(time.Hour * 24 * 10)
	} else if dataReq.PaymentTerms == 30 {
		dataReq.PaymentDue = t.Add(time.Hour * 24 * 30)
	} else {
		dataReq.PaymentDue = t.Add(time.Hour * 24 * 1)
	}

	data, err := inBusiness.invoiceData.CreateInvoice(dataReq)
	if err != nil {
		return invoice.InvoiceCore{}, err
	}
	var item []midtrans.ItemDetails
	item = append(item, midtrans.ItemDetails{
		Name:  dataReq.Item,
		Price: int64(dataReq.Total),
		Qty:   1,
	})
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
		Items: &item,
		CustomerDetail: &midtrans.CustomerDetails{
			FName: dataReq.FullName,
			Email: dataReq.Email,
		},
	})

	if errMidtrans != nil {
		return invoice.InvoiceCore{}, errMidtrans
	}
	data.PaymentLink = resp.RedirectURL

	err = inBusiness.invoiceData.UpdateInvoice(data)
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
func (inBusiness *InvoiceBusiness) GetInvoice(Id int, UserId string) (invoice.InvoiceCore, error) {
	data, err := inBusiness.invoiceData.GetInvoice(Id, UserId)
	if err != nil {
		return invoice.InvoiceCore{}, err
	}

	return data, nil
}
