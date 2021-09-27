package coinmarket

import (
	"aprian1337/thukul-service/business/coinmarket"
	"aprian1337/thukul-service/repository"
	"context"
	"encoding/json"
	"net/http"
)

type MarketCapAPI struct {
	Client         http.Client
	BaseUrl        string
	ApiKey         string
	EndpointSymbol string
}

func NewMarketCapAPI(api MarketCapAPI) *MarketCapAPI {
	return &MarketCapAPI{
		Client:         http.Client{},
		BaseUrl:        api.BaseUrl,
		ApiKey:         api.ApiKey,
		EndpointSymbol: api.EndpointSymbol,
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
	response := Response{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if len(response.Data) == 0 {
		return coinmarket.Domain{}, repository.ErrDataNotFound
	}
	return response.ToDomain(), err
}
