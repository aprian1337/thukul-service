package coinmarket

import (
	"aprian1337/thukul-service/business/coinmarket"
	"aprian1337/thukul-service/repository"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type MarketCapAPI struct {
	Client         http.Client
	BaseUrl        string
	ApiKey         string
	EndpointSymbol string
	EndpointPrice  string
}

func NewMarketCapAPI(api MarketCapAPI) *MarketCapAPI {
	return &MarketCapAPI{
		Client:         http.Client{},
		BaseUrl:        api.BaseUrl,
		ApiKey:         api.ApiKey,
		EndpointSymbol: api.EndpointSymbol,
		EndpointPrice:  api.EndpointPrice,
	}
}

func (api *MarketCapAPI) GetBySymbol(ctx context.Context, symbol string) (coinmarket.Domain, error) {
	uri := api.BaseUrl + api.EndpointSymbol + symbol
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("X-CMC_PRO_API_KEY", api.ApiKey)
	resp, err := api.Client.Do(req)
	if err != nil {
		return coinmarket.Domain{}, err
	}
	response := SymbolResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if len(response.Data) == 0 {
		return coinmarket.Domain{}, repository.ErrDataNotFound
	}
	return response.ToDomain(), err
}

func (api *MarketCapAPI) GetPrice(ctx context.Context, symbol string, amount float64) (float64, error) {
	//GET IDR CURRENCY FROM MARKETCAP API DOCS
	uri := fmt.Sprintf("%v%vsymbol=%v&amount=%v", api.BaseUrl, api.EndpointPrice, symbol, amount)
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("X-CMC_PRO_API_KEY", api.ApiKey)
	resp, err := api.Client.Do(req)
	if err != nil {
		return 0, err
	}
	response := PriceResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	fmt.Println("URI")
	fmt.Println(uri)
	if err != nil {
		return 0, repository.ErrDataNotFound
	}
	return response.Data.Quote.CurrencyIdr.Price, err
}
