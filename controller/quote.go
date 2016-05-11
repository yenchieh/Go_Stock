package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/url"
	"fmt"
	"github.com/go_stock_with_gin/common"
	"net/http"
	"github.com/go_stock_with_gin/model"
	"encoding/json"
	"errors"
)

var yahooFinanceAPI string = "http://query.yahooapis.com/v1/public/yql"

func GetQuoteBySymbol(c *gin.Context){
	var symbolQuote model.YahooData
	symbol := c.Query("symbol")
	fmt.Printf("Getting Quote from symbol: %s\n", symbol)
	if(len(symbol) == 0){
		common.RenderError(c, http.StatusBadRequest, errors.New("Symbol is required"), "Symbol is required")
		return
	}

	log.Printf("Get %s quote data", symbol)
	resultRaw, err := getQuoteFromYahoo(symbol)

	if err != nil {
		common.RenderError(c, http.StatusInternalServerError, err, "Internal Error")
	}

	if err := json.Unmarshal(resultRaw, &symbolQuote); err != nil {
		common.RenderError(c, http.StatusInternalServerError, err, "Error on parsing JSON")
	}

	c.JSON(http.StatusOK, symbolQuote)
}


func getQuoteFromYahoo(symbol string) ([]byte, error){
	var Url *url.URL

	Url, err := url.Parse(yahooFinanceAPI)

	if err != nil {
		log.Fatal(err)
	}
	query := fmt.Sprintf("select finance, Name, Symbol, Change, PercentChange, DaysLow, DaysHigh, Open,PreviousClose, Volume from yahoo.finance.quotes where symbol in (\"%s\")", symbol)
	parameters := url.Values{}
	parameters.Add("q", query)
	parameters.Add("format", "json")
	parameters.Add("env", "http://datatables.org/alltables.env")

	Url.RawQuery = parameters.Encode()

	resultRaw, err := common.RequestService(Url)

	if err != nil {
		return nil, err
	}

	return resultRaw, nil
}
