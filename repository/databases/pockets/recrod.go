package pockets

import (
	"aprian1337/thukul-service/business/pockets"
	"aprian1337/thukul-service/repository/databases/users"
	"gorm.io/gorm"
	"time"
)

type Pockets struct {
	ID        int `gorm:"primaryKey"`
	UserId    int
	User      users.Users `gorm:"foreignKey:user_id"`
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

func (data *Pockets) ToDomain() pockets.Domain {
	return pockets.Domain{
		ID:        data.ID,
		UserId:    data.UserId,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func FromDomain(domain pockets.Domain) Pockets {
	return Pockets{
		ID:        domain.ID,
		UserId:    domain.UserId,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ToListDomain(data []Pockets) []pockets.Domain {
	var list []pockets.Domain
	for _, v := range data {
		list = append(list, v.ToDomain())
	}
	return list
}
