package requests

import (
	"aprian1337/thukul-service/business/activities"
)

type ActivityRequest struct {
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	Nominal float64 `json:"nominal"`
	Note    string  `json:"note"`
	Date    string  `json:"date"`
}

func (s *ActivityRequest) ToDomain() activities.Domain {
	return activities.Domain{
		Name:    s.Name,
		Type:    s.Type,
		Nominal: s.Nominal,
		Note:    s.Note,
		Date:    s.Date,
	}
}
