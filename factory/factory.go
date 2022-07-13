package factory

import (
	"log"
	"watchamovie-payment/config"
	"watchamovie-payment/driver"

	// Invoice Domain
	inbus "watchamovie-payment/features/invoice/business"
	indata "watchamovie-payment/features/invoice/data"
	inpres "watchamovie-payment/features/invoice/presentation"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type presenter struct {
	InvoicePresentation inpres.InvoiceHandler
}

func Init() presenter {
	//Initiate client for Midtrans Snap
	var s snap.Client
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	s.New(config.MidtransApi, midtrans.Sandbox)

	var c coreapi.Client
	c.New(config.MidtransApi, midtrans.Sandbox)

	// Invoice
	invoiceData := indata.NewMySqlInvoice(driver.DB)
	invoiceBusiness := inbus.NewBusinessInvoice(invoiceData, s, c)

	return presenter{
		InvoicePresentation: *inpres.NewHandlerInvoice(invoiceBusiness),
	}
}
