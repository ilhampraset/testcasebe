package domain

import "time"

type Merchant struct {
	ID           int
	UserID       int
	MerchantName string
	CreatedBy    int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UpdatedBy    int
	Outlet       []Outlet
	Transaction  []Transaction
}

type MerchantMontlyReport struct {
	MerchantName string    `json:"merchant_name"`
	Reveneu      int       `gorm:"column:revenue" json:"omzet"`
	CreatedAt    time.Time `json:"tanggal"`
}

type IMerchantRepository interface {
	GetMonthlyRevenueByMerchantName(merchantName string, year int, month int) (*[]MerchantMontlyReport, error)
}

type IMerchantUsecase interface {
	GetMonthlyRevenue(merchantName string, year int, month int) (*[]MerchantMontlyReport, error)
}
