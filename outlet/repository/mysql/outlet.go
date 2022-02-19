package mysql

import (
	"fmt"

	"github.com/ilhampraset/testcasebe/domain"
	"gorm.io/gorm"
)

type OutletRepository struct {
	db *gorm.DB
}

func NewOutletRepository(db *gorm.DB) domain.IOutletRepository {
	return &OutletRepository{db}
}

func (o OutletRepository) GetMonthlyRevenueByOutletName(merchantName string, year int, month int) (*[]domain.OutletMontlyReport, error) {
	var outlets *[]domain.OutletMontlyReport
	from := fmt.Sprintf("%d-%d-1", year, month)
	end := fmt.Sprintf("%d-%d-30", year, month)
	o.db.Debug().Table("merchants").Select("merchants.merchant_name,outlets.outlet_name,sum(transactions.bill_total) as revenue, transactions.created_at").
		Joins("inner join users on users.id = merchants.user_id").
		Joins("inner join outlets on outlets.merchant_id = merchants.id").
		Joins("inner join transactions on transactions.outlet_id = outlets.id").
		Where("DATE(transactions.created_at) between ? and ? and merchants.merchant_name = ? ", from, end, merchantName).
		Group("DATE(transactions.created_at), outlets.id").
		Find(&outlets)
	return outlets, nil
}
