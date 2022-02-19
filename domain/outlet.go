package domain

import "time"

type Outlet struct {
	ID          int
	MerchantId  int
	OutletName  string
	CreatedBy   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UpdatedBy   int
	Transaction []Transaction
}

type OutletMontlyReport struct {
	MerchantName string    `json:"merchant_name"`
	OutletName   string    `json:"outlet_name"`
	Reveneu      int       `gorm:"column:revenue" json:"omzet"`
	CreatedAt    time.Time `json:"tanggal"`
}

type OutletRevenueInfo struct {
	OutletName string `json:"outlet_name"`
	Revenue    int    `json:"omzet"`
}
type OutletReport struct {
	MerchantName string              `json:"merchant_name"`
	Outlet       []OutletRevenueInfo `json:"outlet"`
	//Reveneu      int                 `gorm:"column:revenue" json:"omzet"`
	CreatedAt time.Time `json:"tanggal"`
}

type IOutletRepository interface {
	GetMonthlyRevenueByOutletName(merchantName string, year int, month int) (*[]OutletMontlyReport, error)
}

type IOutletUsecase interface {
	GetMonthlyRevenue(merchantName string, year int, month int) (*[]OutletReport, error)
}
