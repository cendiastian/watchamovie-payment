package data

import (
	"errors"
	"fmt"
	"watchamovie-payment/features/invoice"

	"gorm.io/gorm"
)

type InvoiceData struct {
	DB *gorm.DB
}

func NewMySqlInvoice(DB *gorm.DB) invoice.Data {
	return &InvoiceData{DB}
}

func (inData *InvoiceData) CreateInvoice(data invoice.InvoiceCore) (invoice.InvoiceCore, error) {
	convData := toInvoiceRecord(data)

	if err := inData.DB.Create(&convData).Error; err != nil {
		return invoice.InvoiceCore{}, err
	}
	return toInvoiceCore(convData), nil
}
func (inData *InvoiceData) GetInvoice(Id int, UserId string) (invoice.InvoiceCore, error) {
	var convData Invoice

	if err := inData.DB.First(&convData, "ID= ?", Id).Error; err != nil {
		return invoice.InvoiceCore{}, err
	}
	if convData.UserId != UserId {
		return invoice.InvoiceCore{}, errors.New("UserId is different")
	}
	return toInvoiceCore(convData), nil
}

func (inData *InvoiceData) UpdateInvoice(data invoice.InvoiceCore) error {
	fmt.Println("Isi data di data : ", data)
	// var singleData Invoice
	convData := toInvoiceRecord(data)
	fmt.Println("Isi convData di data : ", convData)
	err := inData.DB.Debug().Model(&Invoice{}).Where("id = ?", data.Id).Updates(&convData).Error
	if err != nil {
		return err
	}
	err = inData.UpdatePaymentLink(convData.PaymentLink, convData.ID)
	if err != nil {
		return err
	}
	return nil
}

func (inData *InvoiceData) UpdateTransactionStatus(transactionID int64, PaymentStatus string) error {
	err := inData.DB.Model(&Invoice{}).Where("id = ?", transactionID).Update("payment_status", PaymentStatus).Error
	if err != nil {
		return err
	}
	return nil
}

func (inData *InvoiceData) UpdatePaymentLink(url string, id uint) error {
	err := inData.DB.Debug().Model(&Invoice{}).Where("id = ?", id).Update("payment_link", url).Error
	if err != nil {
		return err
	}
	return nil
}
