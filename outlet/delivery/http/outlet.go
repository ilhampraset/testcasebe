package http

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilhampraset/testcasebe/auth/utils"
	"github.com/ilhampraset/testcasebe/domain"
)

type OutletHandler struct {
	uc  domain.IOutletUsecase
	jwt utils.Token
}

type Info struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
}
type MonthlyReportDto struct {
	Info Info                  `json:"info"`
	Data []domain.OutletReport `json:"data"`
}

func NewOutletHandler(uc domain.IOutletUsecase, jwt utils.Token) *OutletHandler {
	return &OutletHandler{uc, jwt}
}

func (h *OutletHandler) MonthlyReport(c *gin.Context) {
	var result MonthlyReportDto

	page, _ := strconv.Atoi(c.Query("page"))

	per_page, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	month, _ := strconv.Atoi(c.DefaultQuery("bulan", "11"))
	year, _ := strconv.Atoi(c.DefaultQuery("tahun", "2021"))

	report, err := h.uc.GetMonthlyRevenue(c.Query("merchant_name"), year, month)
	if err != nil {

		c.AbortWithStatusJSON(404, gin.H{"message": "something wrong"})
	}
	result.Info.Total = 0
	result.Info.Page = page
	result.Info.PerPage = per_page
	if report != nil {
		total := len(*report)
		start := 0
		end := per_page
		if page > 1 {
			start = (page - 1) * per_page
			end = page * per_page
			if end > total {
				end = total
			}
		}
		fmt.Println(start)
		fmt.Println(end)
		result.Data = (*report)[start:end]
		result.Info.Total = total
		c.JSON(200, result)
	} else {

		c.JSON(200, result)
	}

}
