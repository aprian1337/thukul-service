package coins

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

func (data *Coins) ToDomain() coins.Domain {
	return coins.Domain{
		Id:        data.ID,
		Symbol:    data.Symbol,
		Name:      data.Name,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func FromDomain(domain coins.Domain) Coins {
	return Coins{
		ID:        domain.Id,
		Symbol:    domain.Symbol,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
