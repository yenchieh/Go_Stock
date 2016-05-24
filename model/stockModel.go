package model

type (
	Stocks struct {
		Id     uint32 `json:"_id,omitempty"`
		UserId int `json:"userId"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
		Type string `json:"type"`
	}
)
