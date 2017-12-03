package controller

import (
	"io/ioutil"
	"net/http"

	"fmt"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/yenchieh/Go_Stock/config"
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

type StockDetailIn struct {
	MetaData   MetaDataIn             `json:"Meta Data"`
	TimeSeries map[string]StockFeedIn `json:"Time Series (5min)"`
}

type MetaDataIn struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	Interval      string `json:"4. Interval"`
	OutputSize    string `json:"5. Output Size"`
	TimeZone      string `json:"6. Time Zone"`
}

type StockFeedIn struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

type StockDetail struct {
	MetaData   MetaData             `json:"meta_data"`
	TimeSeries map[string]StockFeed `json:"time_series"`
}

type MetaData struct {
	Information   string `json:"information"`
	Symbol        string `json:"symbol"`
	LastRefreshed string `json:"last_refreshed"`
	Interval      string `json:"interval"`
	OutputSize    string `json:"output_size"`
	TimeZone      string `json:"time_zone"`
}

type StockFeed struct {
	Open   string `json:"open"`
	High   string `json:"high"`
	Low    string `json:"low"`
	Close  string `json:"close"`
	Volume string `json:"volume"`
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

func StockDetailByTicker(c *gin.Context) {
	ticker := c.Param("ticker")

	key := config.Env.AlphaVantageKey

	if ticker == "" {
		renderJSONError(c, nil, http.StatusBadRequest, "Parameter 'ticker' is missing")
		return
	}

	baseURL := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%s&interval=5min&apikey=%s", ticker, key)

	resp, err := http.Get(baseURL)
	if err != nil {
		renderJSONError(c, errors.Wrapf(err, "Error on retrieve intraday. URL: %s", baseURL), http.StatusInternalServerError, "Error on retrieve intraday")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		renderJSONError(c, errors.Wrap(err, "Error on reading response body"), http.StatusInternalServerError, "Error on reading response body")
		return
	}

	var stockDetailIn StockDetailIn
	if err := json.Unmarshal(body, &stockDetailIn); err != nil {
		renderJSONError(c, errors.Wrapf(err, "Error on parse json. Response: %s", string(body)), http.StatusInternalServerError, "Error on parse json")
		return
	}

	timeSeries := make(map[string]StockFeed)
	for key, value := range stockDetailIn.TimeSeries {
		timeSeries[key] = StockFeed(value)
	}
	stockDetail := StockDetail{
		MetaData:   MetaData(stockDetailIn.MetaData),
		TimeSeries: timeSeries,
	}

	c.JSON(http.StatusOK, stockDetail)
}
