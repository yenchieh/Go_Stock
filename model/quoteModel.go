package model

import "gopkg.in/mgo.v2/bson"

type (
	YahooData struct {
		Id bson.ObjectId `bson:"_id,omitempty"`
		Query YahooQuery `json:"query"`
	}

	YahooQuery struct {
		Count int `json:"count"`
		Created string `json:"created"`
		Lang string `json:"lang"`
		Results YahooQuote `json:"results"`
	}

	YahooQuote struct {
		Quote StockQuote `json:"quote"`
	}

	StockQuote struct {
		Change string `json:"change"`
		PercentChange string `json:"percentChange"`
		DaysLow string `json:"daysLow"`
		DaysHigh string `json:"daysHigh"`
		Open string `json:"open"`
		PreviousClose string `json:"previousClose"`
		Symbol string `json:"symbol"`
		Name string `json:"name"`
		Volume string `json:"volume"`
	}


)