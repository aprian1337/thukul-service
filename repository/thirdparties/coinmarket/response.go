package coinmarket

import (
	"aprian1337/thukul-service/business/coinmarket"
	"time"
)

type SymbolResponse struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       interface{} `json:"notice"`
	} `json:"status"`
	Data []struct {
		ID                  int         `json:"id"`
		Name                string      `json:"name"`
		Symbol              string      `json:"symbol"`
		Slug                string      `json:"slug"`
		Rank                int         `json:"rank"`
		IsActive            int         `json:"is_active"`
		FirstHistoricalData time.Time   `json:"first_historical_data"`
		LastHistoricalData  time.Time   `json:"last_historical_data"`
		Platform            interface{} `json:"platform"`
	} `json:"data"`
}

type SymbolData struct {
	ID                  int         `json:"id"`
	Name                string      `json:"name"`
	Symbol              string      `json:"symbol"`
	Slug                string      `json:"slug"`
	Rank                int         `json:"rank"`
	IsActive            int         `json:"is_active"`
	FirstHistoricalData time.Time   `json:"first_historical_data"`
	LastHistoricalData  time.Time   `json:"last_historical_data"`
	Platform            interface{} `json:"platform"`
}

func (resp *SymbolResponse) ToDomain() coinmarket.Domain {
	return coinmarket.Domain{
		Symbol: resp.Data[0].Symbol,
		Name:   resp.Data[0].Name,
	}
}

type PriceResponse struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       interface{} `json:"notice"`
	} `json:"status"`
	Data struct {
		ID          int       `json:"id"`
		Symbol      string    `json:"symbol"`
		Name        string    `json:"name"`
		Amount      float64   `json:"amount"`
		LastUpdated time.Time `json:"last_updated"`
		Quote       struct {
			CurrencyIdr struct {
				Price       float64   `json:"price"`
				LastUpdated time.Time `json:"last_updated"`
			} `json:"2794"`
		} `json:"quote"`
	} `json:"data"`
}
