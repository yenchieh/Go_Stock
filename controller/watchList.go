package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go_stock_with_gin/model"
	"github.com/go_stock_with_gin/common"
	"net/http"
	"log"
	"github.com/go_stock_with_gin/config"
	"github.com/go_stock_with_gin/data"
	"errors"
)

type AddStockParams struct {
	Email string `json:"email"`
	Stocks []string `json:"stocks"`

}

type response struct {
	Success bool `json:"success"`
	WatchList []model.Stocks `json:"watchList"`
}

type quoteResponse struct {
	Success bool `json:"success"`
	WatchList model.YahooData `json:"watchList"`
}

func AddStockToWatchList(c *gin.Context){
	var addStockParams AddStockParams
	if err := c.BindJSON(&addStockParams); err != nil {
		common.RenderError(c, http.StatusBadRequest, err, "Error on convert JSON parameter")
		return
	}

	log.Printf("Email: %s, StockName: %s", addStockParams.Email, addStockParams.Stocks)


	user, err := model.GetUserByEmail(addStockParams.Email)

	if err != nil {
		common.RenderError(c, http.StatusBadRequest, err, "Error on finding user")
		return
	}

	//var stocks []model.Stocks
	for _, symbol := range addStockParams.Stocks {
		stock := model.Stocks{Symbol: symbol, Type: config.WATCH_LIST_STOCK}
		user.Stocks = append(user.Stocks, stock)
	}

	//user.Stocks = stocks

	if err := user.SaveStock(); err != nil {
		common.RenderError(c, http.StatusInternalServerError, err, "Error on saving list to database")
		return
	}

	log.Printf("Success saving stock list. %v", user.Stocks)

	//Get all of user's watch stocks
	db := data.GetDatabase()
	defer db.Close()

	watchStocks, err := db.Query("SELECT u.id, s.symbol, us.type FROM user u join user_stocks_connection us on u.id = us.user_id join stocks s on s.id = us.stock_id WHERE type = ?", config.WATCH_LIST_STOCK)

	if err != nil {
		common.RenderError(c, http.StatusInternalServerError, err, "Error on getting watch list")
		return
	}

	response := response{
		Success: true,
	}

	for watchStocks.Next() {
		var stock model.Stocks
		watchStocks.Scan(&stock.UserId, &stock.Symbol, &stock.Type)
		response.WatchList = append(response.WatchList, stock)
	}



	c.JSON(http.StatusOK, response)
}

func UserWatchListByEmail(c *gin.Context){
	email := c.Query("email")

	if email == "" {
		common.RenderError(c, http.StatusBadRequest, errors.New("Email is required."), "Email is required")
		return
	}

	stocks, err := model.UserWatchList(email)

	if err != nil {
		common.RenderError(c, http.StatusInternalServerError, err, "Error on getting user's watch list")
		return
	}

	quotes, err := GetQuoteByStocks(stocks)

	if err != nil {
		common.RenderError(c, http.StatusInternalServerError, err, "Error on getting quotes from Yahoo")
		return
	}

	response := quoteResponse{
		WatchList: quotes,
		Success: true,
	}

	c.JSON(http.StatusOK, response)


}