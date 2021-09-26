package responses

import (
	"aprian1337/thukul-service/business/pockets"
	"time"
)

type PocketsResponse struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain pockets.Domain) PocketsResponse {
	return PocketsResponse{
		ID:        domain.ID,
		UserId:    domain.UserId,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromListDomain(domain []pockets.Domain) []PocketsResponse {
	var result []PocketsResponse
	for _, v := range domain {
		result = append(result, FromDomain(v))
	}
	return result
}
