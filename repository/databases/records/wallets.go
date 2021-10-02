package records

import (
	"aprian1337/thukul-service/business/wallets"
	"time"
)

type Wallets struct {
	ID        int `gorm:"primaryKey"`
	UserId    int
	User      Users `gorm:"foreignKey:user_id"`
	Total     float64
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (data *Wallets) WalletsToDomain() wallets.Domain {
	return wallets.Domain{
		Id:     data.ID,
		UserId: data.UserId,
		Total:  data.Total,
	}
}

func WalletsFromDomain(domain wallets.Domain) Wallets {
	return Wallets{
		ID:     domain.Id,
		UserId: domain.UserId,
		Total:  domain.Total,
	}
}

func WalletsToListDomain(data []Wallets) []wallets.Domain {
	var list []wallets.Domain
	for _, v := range data {
		list = append(list, v.WalletsToDomain())
	}
	return list
}
