package model

import (
	"github.com/go_stock_with_gin/data"
	"github.com/go_stock_with_gin/config"
	"database/sql"
	"log"
)

type (
	Stocks struct {
		Id     uint32 `json:"_id,omitempty"`
		UserId int `json:"userId"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
		Type string `json:"type"`
	}
)


func UserWatchList(email string) (watchList []Stocks, err error){
	db := data.GetDatabase()
	defer db.Close()

	result, err := db.Query("SELECT s.symbol, s.name, u.id FROM stocks s JOIN user_stocks_connection us on s.id = us.stock_id JOIN user u on u.id = us.user_id WHERE type = ? AND u.email = ?", config.WATCH_LIST_STOCK, email)

	if err != nil && err != sql.ErrNoRows {
		return watchList, err
	} else if err == sql.ErrNoRows {
		return watchList, nil
	}

	for result.Next() {
		var stock Stocks
		if err := result.Scan(&stock.Symbol, &stock.Name, &stock.UserId); err != nil {
			log.Printf("Error on converting sql to struct, %v", err)
		}else{
			watchList = append(watchList, stock)
		}
	}

	return watchList, nil

}