package wallets

import (
	"aprian1337/thukul-service/business/wallets"
	"aprian1337/thukul-service/repository/databases/users"
	"time"
)

type Wallets struct {
	ID        int `gorm:"primaryKey"`
	UserId    int
	User      users.Users `gorm:"foreignKey:user_id"`
	Total     float64
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (data *Wallets) ToDomain() wallets.Domain {
	return wallets.Domain{
		Id:     data.ID,
		UserId: data.UserId,
		Total:  data.Total,
	}
}

func FromDomain(domain wallets.Domain) Wallets {
	return Wallets{
		ID:     domain.Id,
		UserId: domain.UserId,
		Total:  domain.Total,
	}
}

func ToListDomain(data []Wallets) []wallets.Domain {
	var list []wallets.Domain
	for _, v := range data {
		list = append(list, v.ToDomain())
	}
	return list
}
