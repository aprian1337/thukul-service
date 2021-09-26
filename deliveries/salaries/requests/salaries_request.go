package requests

import (
	"aprian1337/thukul-service/business/salaries"
)

type SalariesRequest struct {
	Minimal float64 `json:"minimal"`
	Maximal float64 `json:"maximal"`
}

func (s *SalariesRequest) ToDomain() salaries.Domain {
	return salaries.Domain{
		Minimal: s.Minimal,
		Maximal: s.Maximal,
	}
}
