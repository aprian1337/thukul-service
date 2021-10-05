package responses

import "aprian1337/thukul-service/business/cryptos"

type CryptosResponse struct {
	Id     int         `json:"id"`
	Amount float64     `json:"amount"`
	Coin   interface{} `json:"coin,omitempty"`
}

func FromDomainCrypto(d cryptos.Domain) CryptosResponse {
	return CryptosResponse{
		Id:     d.ID,
		Coin:   d.Coin,
		Amount: d.Qty,
	}
}

func FromDomainCryptoList(d []cryptos.Domain) []CryptosResponse {
	var data []CryptosResponse
	for _, v := range d {
		data = append(data, FromDomainCrypto(v))
	}
	return data
}
