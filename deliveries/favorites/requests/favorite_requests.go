package requests

import (
	"aprian1337/thukul-service/business/favorites"
	coins2 "aprian1337/thukul-service/repository/databases/coins"
)

type FavoriteRequest struct {
	CoinSymbol string `json:"coin_symbol"`
}

func (fav *FavoriteRequest) ToDomain() favorites.Domain {
	return favorites.Domain{
		Coin: coins2.Coins{
			Symbol: fav.CoinSymbol,
		},
	}
}
