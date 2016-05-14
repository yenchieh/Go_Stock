package model

import (
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"bytes"
)

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
		Quote *StockQuote `json:"quote"`
		Quotes []StockQuote `json:"quotes"`
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

	//Handle ArrayStruct or Struct
	structOrArray struct {
		parent *YahooQuote
		s StockQuote
		a []StockQuote
	}

	middleQuote struct {
		Load structOrArray `json:"quote"`
	}


)

func (this *structOrArray) UnmarshalJSON(data []byte) error{
	d := json.NewDecoder(bytes.NewBuffer(data))
	t, err := d.Token()
	if err != nil {
		return err
	}

	if t == json.Delim('['){
		if err := json.Unmarshal(data, &this.a); err != nil {
			return err
		}
		return nil
	}

	if err := json.Unmarshal(data, &this.s); err != nil{
		return err
	}

	return nil
}

func (this *YahooQuote) UnmarshalJSON(data []byte) error {
	mq := middleQuote{}
	if err := json.Unmarshal(data, &mq); err != nil{
		return err
	}
	this.Quote = &mq.Load.s
	this.Quotes = mq.Load.a
	return nil
}