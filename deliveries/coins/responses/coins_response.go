package responses

import "aprian1337/thukul-service/business/coins"

type CoinsResponse struct {
	Id     int    `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

func CoinsFromDomain(domain coins.Domain) CoinsResponse {
	return CoinsResponse{
		Id:     domain.Id,
		Symbol: domain.Symbol,
		Name:   domain.Name,
	}
}

func ListCoinsFromDomain(domain []coins.Domain) []CoinsResponse {
	var res []CoinsResponse
	for _, v := range domain {
		res = append(res, CoinsFromDomain(v))
	}
	return res
}
