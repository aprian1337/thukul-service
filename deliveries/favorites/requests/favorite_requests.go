package requests

import (
	"aprian1337/thukul-service/business/favorites"
)

type FavoriteRequest struct {
	CoinSymbol string `json:"coin_symbol"`
}

func (fav *FavoriteRequest) ToDomain() favorites.Domain {
	return favorites.Domain{
		Symbol: fav.CoinSymbol,
	}
}
