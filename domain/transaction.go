package domain

import "time"

type Transaction struct {
	ID         int
	MerchantID int
	OutletID   int
	BillTotal  float32
	CreatedBy  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UpdatedBy  int
}

type TransactionRepository interface {
	getMonthlyReportByMerchantName(merchantName string) (*Merchant, error)
	getMonthlyReportByMerchantNameAndOutletName(merchantName, OutletName string) (*Merchant, error)
}
