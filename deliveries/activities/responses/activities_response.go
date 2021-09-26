package responses

import (
	"aprian1337/thukul-service/business/activities"
	"time"
)

type ActivitiesResponse struct {
	ID        int       `json:"id"`
	PocketId  int       `json:"pocket_id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Nominal   float64   `json:"nominal"`
	Note      string    `json:"note"`
	Date      string    `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain activities.Domain) ActivitiesResponse {
	return ActivitiesResponse{
		ID:        domain.ID,
		PocketId:  domain.PocketId,
		Name:      domain.Name,
		Type:      domain.Type,
		Nominal:   domain.Nominal,
		Note:      domain.Note,
		Date:      domain.Date,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromListDomain(domain []activities.Domain) []ActivitiesResponse {
	var result []ActivitiesResponse
	for _, v := range domain {
		result = append(result, FromDomain(v))
	}
	return result
}
