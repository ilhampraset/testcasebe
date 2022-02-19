package mysql

import (
	"fmt"

	"github.com/ilhampraset/testcasebe/domain"
	"gorm.io/gorm"
)

type MerchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) domain.IMerchantRepository {
	return &MerchantRepository{db}
}

func (m MerchantRepository) GetMonthlyRevenueByMerchantName(merchantName string, year int, month int) (*[]domain.MerchantMontlyReport, error) {
	var merchants *[]domain.MerchantMontlyReport
	from := fmt.Sprintf("%d-%d-1", year, month)
	end := fmt.Sprintf("%d-%d-30", year, month)
	m.db.Table("merchants").Select("merchants.merchant_name,sum(transactions.bill_total) as revenue, transactions.created_at").
		Joins("inner join users on users.id = merchants.user_id").
		Joins("inner join transactions on transactions.merchant_id = merchants.id").
		Where("DATE(transactions.created_at) between ? and ? and merchants.merchant_name = ? ", from, end, merchantName).
		Group("DATE(transactions.created_at)").
		Find(&merchants)

	return merchants, nil
}
