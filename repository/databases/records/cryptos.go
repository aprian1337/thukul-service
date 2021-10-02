package records

import (
	"aprian1337/thukul-service/business/cryptos"
	"time"
)

type Cryptos struct {
	ID        int `gorm:"primaryKey"`
	UserId    int
	CoinId    int
	Qty       float64
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func CryptosFromDomain(domain cryptos.Domain) Cryptos {
	return Cryptos{
		UserId: domain.UserId,
		CoinId: domain.CoinId,
		Qty:    domain.Qty,
	}
}

func (c *Cryptos) CryptosToDomain() cryptos.Domain {
	return cryptos.Domain{
		ID:        c.ID,
		UserId:    c.UserId,
		CoinId:    c.CoinId,
		Qty:       c.Qty,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
