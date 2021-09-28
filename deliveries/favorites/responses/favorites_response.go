package responses

import (
	"aprian1337/thukul-service/business/favorites"
	"time"
)

type FavoriteResponse struct {
	ID     int `json:"id"`
	UserId int `json:"user_id"`
	CoinId int `json:"coin_id"`
	Coins  struct {
		CoinSymbol string `json:"coin_symbol"`
		CoinName   string `json:"coin_name"`
	} `json:"coins"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain favorites.Domain) FavoriteResponse {
	return FavoriteResponse{
		ID:     domain.ID,
		UserId: domain.UserId,
		CoinId: domain.CoinId,
		Coins: struct {
			CoinSymbol string `json:"coin_symbol"`
			CoinName   string `json:"coin_name"`
		}(struct {
			CoinSymbol string
			CoinName   string
		}{
			CoinSymbol: domain.Coins.CoinSymbol,
			CoinName:   domain.Coins.CoinName,
		}),
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromListDomain(domain []favorites.Domain) []FavoriteResponse {
	var result []FavoriteResponse
	for _, v := range domain {
		result = append(result, FromDomain(v))
	}
	return result
}
