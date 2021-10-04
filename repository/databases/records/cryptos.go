package records

import (
	"aprian1337/thukul-service/business/cryptos"
	"time"
)

type Cryptos struct {
	ID        int `gorm:"primaryKey"`
	UserId    int
	User      Users `gorm:"foreignKey:user_id"`
	CoinId    int
	Coin      Coins `gorm:"foreignKey:coin_id"`
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
		ID:     c.ID,
		UserId: c.UserId,
		CoinId: c.CoinId,
		Coin: cryptos.Coin(struct {
			Name   string
			Symbol string
		}{Name: c.Coin.Name, Symbol: c.Coin.Symbol}),
		Qty:       c.Qty,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func CryptosToListDomain(d []Cryptos) []cryptos.Domain {
	var cryptoList []cryptos.Domain
	for _, v := range d {
		cryptoList = append(cryptoList, v.CryptosToDomain())
	}
	return cryptoList
}
