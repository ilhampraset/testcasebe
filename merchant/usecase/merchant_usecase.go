package usecase

import (
	"fmt"
	"strings"
	"time"

	"github.com/ilhampraset/testcasebe/domain"
)

type MerchantUsecase struct {
	repo domain.IMerchantRepository
}

func NewMerchantUseCase(repo domain.IMerchantRepository) domain.IMerchantUsecase {
	return &MerchantUsecase{repo}
}

func (uc *MerchantUsecase) GetMonthlyRevenue(merchantName string, year int, month int) (*[]domain.MerchantMontlyReport, error) {

	merch, err := uc.repo.GetMonthlyRevenueByMerchantName(merchantName, year, month)
	if err != nil {
		return nil, err
	}

	merchants := make([]domain.MerchantMontlyReport, 30)
	hmap := make(map[string]int)
	for _, val := range *merch {
		date := strings.Split(val.CreatedAt.String(), " ")
		hmap[date[0]] = val.Reveneu
	}

	if len(hmap) > 0 {
		incDate := 1
		for i := 0; i < 30; i++ {
			input := fmt.Sprintf("%d-%d-%d", year, month, incDate)
			incDate++
			layout := "2006-01-2"
			t, _ := time.Parse(layout, input)
			merchants[i].MerchantName = (*merch)[0].MerchantName
			date := strings.Split(t.String(), " ")
			if val, ok := hmap[date[0]]; ok {
				merchants[i].Reveneu = val
			}
			merchants[i].CreatedAt = t

		}
		return &merchants, nil
	}

	return nil, err
}
