package usecase

import (
	"fmt"
	"strings"
	"time"

	"github.com/ilhampraset/testcasebe/domain"
)

type OutletUsecase struct {
	repo domain.IOutletRepository
}

func NewOutletUseCase(repo domain.IOutletRepository) domain.IOutletUsecase {
	return &OutletUsecase{repo}
}

func (uc *OutletUsecase) GetMonthlyRevenue(merchantName string, year int, month int) (*[]domain.OutletReport, error) {

	out, err := uc.repo.GetMonthlyRevenueByOutletName(merchantName, year, month)

	if err != nil {
		return nil, err
	}
	outlets := make([]domain.OutletReport, 30)
	hmap := make(map[string]int)
	mapOutlet := make(map[string]bool)

	for _, val := range *out {
		date := strings.Split(val.CreatedAt.String(), " ")
		key := fmt.Sprintf("%s&%s", val.OutletName, date[0])
		hmap[key] = val.Reveneu
		mapOutlet[val.OutletName] = true
	}

	incDate := 1
	for i := 0; i < 30; i++ {
		input := fmt.Sprintf("%d-%d-%d", year, month, incDate)
		incDate++
		layout := "2006-01-2"
		t, _ := time.Parse(layout, input)

		date := strings.Split(t.String(), " ")
		outlets[i].MerchantName = (*out)[0].MerchantName
		outinfo := outletInfo(mapOutlet, hmap, date[0])
		outlets[i].Outlet = append(outlets[i].Outlet, outinfo...)

		outlets[i].CreatedAt = t

	}

	return &outlets, nil
}

func outletInfo(omap map[string]bool, hmap map[string]int, date string) []domain.OutletRevenueInfo {
	res := make([]domain.OutletRevenueInfo, 2)
	keys := make([]string, len(omap))

	i := 0
	for k := range omap {
		keys[i] = k
		res[i].OutletName = keys[i]

		if val, ok := hmap[fmt.Sprintf("%s&%s", keys[i], date)]; ok {
			res[i].Revenue = val
		}

		i++
	}

	return res
}
