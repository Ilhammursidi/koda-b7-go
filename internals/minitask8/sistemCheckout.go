package minitask8

import (
	"fmt"
)

type Bank struct {
	BankName string
}

func (b Bank) Bayar(amount int) error {
	fmt.Printf("[Bank] Pembyaran berhasil sebesar Rp %d melalui : %v\n", amount, b.BankName)
	return nil
}

type Online struct {
	EbankName string
}

func (o Online) Bayar(amount int) error {
	fmt.Printf("[E-Bank] Pembyaran berhasil sebesar Rp %d melalui : %v\n", amount, o.EbankName)
	return nil
}

type Pembayaran interface {
	Bayar(amount int) error
}

func MethodPay(metode Pembayaran, prices []int) {
	for _, price := range prices {
		err := metode.Bayar(price)
		if err != nil {
			fmt.Printf("[ERROR] %v\n", err)
		}
	}
}

type FictivePayment struct {
	records []int
}

func (f *FictivePayment) Pay(amount int) error {
	if amount <= 0 {
		return fmt.Errorf("biaya pembayaran tidak valid: %d (harus lebih dari 0)", amount)
	}
	f.records = append(f.records, amount)
	return nil
}

func (f *FictivePayment) GetRecords() []int {
	return f.records
}
