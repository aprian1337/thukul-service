package favorites

import (
	"aprian1337/thukul-service/business/favorites"
	"aprian1337/thukul-service/repository/databases/coins"
	"aprian1337/thukul-service/repository/databases/users"
	"time"
)

type Favorites struct {
	ID        int `gorm:"primarykey"`
	UserId    int
	User      users.Users `gorm:"foreignKey:user_id"`
	CoinId    int
	Coin      coins.Coins `gorm:"foreignKey:coin_id"`
	CreatedAt time.Time   `gorm:"autoCreateTime"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime"`
}

func (data *Favorites) ToDomain() favorites.Domain {
	return favorites.Domain{
		ID:        data.ID,
		UserId:    data.UserId,
		CoinId:    data.CoinId,
		Coin:      data.Coin,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func FromDomain(domain favorites.Domain) Favorites {
	return Favorites{
		ID:        domain.ID,
		UserId:    domain.UserId,
		CoinId:    domain.CoinId,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ToListDomain(data []Favorites) []favorites.Domain {
	var list []favorites.Domain
	for _, v := range data {
		list = append(list, v.ToDomain())
	}
	return list
}
