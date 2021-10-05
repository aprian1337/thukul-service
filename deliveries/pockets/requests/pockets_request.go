package requests

import (
	"aprian1337/thukul-service/business/pockets"
)

type PocketsRequest struct {
	UserId int
	Name   string `json:"name"`
}

type TotalRequest struct {
	Type string `json:"type"`
}

func (s *PocketsRequest) ToDomain() pockets.Domain {
	return pockets.Domain{
		UserId: s.UserId,
		Name:   s.Name,
	}
}
