package responses

import "aprian1337/thukul-service/business/wallets"

type TopUpResponse struct {
	Message string     `json:"message"`
	Data    *TopUpData `json:"data"`
}

type TopUpData struct {
	Total              float64 `json:"total"`
	NominalTransaction float64 `json:"nominal_transaction"`
}

func FromDomainWallets(domain wallets.Domain) TopUpData {
	return TopUpData{
		Total:              domain.Total,
		NominalTransaction: domain.NominalTransaction,
	}
}
