package responses

import (
	"aprian1337/thukul-service/business/salaries"
	"time"
)

type SalariesResponse struct {
	ID        uint      `json:"id"`
	Maximal   float64   `json:"maximal"`
	Minimal   float64   `json:"minimal"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain salaries.Domain) SalariesResponse {
	return SalariesResponse{
		ID:        domain.ID,
		Maximal:   domain.Maximal,
		Minimal:   domain.Minimal,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromListDomain(domain []salaries.Domain) []SalariesResponse {
	var result []SalariesResponse
	for _, v := range domain {
		result = append(result, FromDomain(v))
	}
	return result
}
