package controller

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Company struct {
	ID           int     `json:"id" db:"id"`
	Symbol       string  `json:"symbol" db:"symbol"`
	Name         string  `json:"name" db:"name"`
	LastSale     float32 `json:"last_sale" db:"last_sale"`
	MarketCap    float32 `json:"market_cap" db:"market_cap"`
	IPOYear      int     `json:"ipo_year" db:"ipo_year"`
	Sector       string  `json:"sector" db:"sector"`
	Industry     string  `json:"industry" db:"industry"`
	SummaryQuote string  `json:"summary_quote" db:"summary_quote"`
}

func SearchStock(c *gin.Context) {
	db := NewDB()
	defer db.Close()

	var company []Company

	keyword := fmt.Sprintf("%%%s%%", c.Query("keyword"))
	if err := db.Select(&company, "SELECT * FROM company_list WHERE symbol Ilike $1 or name Ilike $1 LIMIT 20", keyword); err != nil {
		renderJSONError(c, errors.Wrapf(err, "Error on retrieve company_list. Params: %s", c.Query("keyword")), http.StatusInternalServerError, "Error on retrieve company_list")
		return
	}

	c.JSON(http.StatusOK, company)

}
