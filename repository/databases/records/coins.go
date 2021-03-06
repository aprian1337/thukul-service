package records

import (
	"aprian1337/thukul-service/business/coins"
	"time"
)

type Coins struct {
	ID        int `gorm:"primaryKey"`
	Symbol    string
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (data *Coins) CoinsToDomain() coins.Domain {
	return coins.Domain{
		Id:        data.ID,
		Symbol:    data.Symbol,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func CoinsToListDomain(list []Coins) []coins.Domain {
	var data []coins.Domain
	for _, v := range list {
		data = append(data, v.CoinsToDomain())
	}
	return data
}

func CoinsFromDomain(domain coins.Domain) Coins {
	return Coins{
		ID:        domain.Id,
		Symbol:    domain.Symbol,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
