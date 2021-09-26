package requests

import (
	"aprian1337/thukul-service/business/pockets"
)

type PocketsRequest struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
}

func (s *PocketsRequest) ToDomain() pockets.Domain {
	return pockets.Domain{
		UserId: s.UserId,
		Name:   s.Name,
	}
}
