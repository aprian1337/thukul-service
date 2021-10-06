package records

import (
	"aprian1337/thukul-service/business/favorites"
	"time"
)

type Favorites struct {
	ID        int `gorm:"primarykey"`
	UserId    int
	User      Users `gorm:"foreignKey:user_id"`
	CoinId    int
	Coin      Coins     `gorm:"foreignKey:coin_id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (data *Favorites) FavoritesToDomain() favorites.Domain {
	return favorites.Domain{
		ID:     data.ID,
		UserId: data.UserId,
		CoinId: data.CoinId,
		Coin:      data.Coin,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func FavoritesFromDomain(domain favorites.Domain) Favorites {
	return Favorites{
		ID:        domain.ID,
		UserId:    domain.UserId,
		CoinId:    domain.CoinId,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FavoritesToListDomain(data []Favorites) []favorites.Domain {
	var list []favorites.Domain
	for _, v := range data {
		list = append(list, v.FavoritesToDomain())
	}
	return list
}
