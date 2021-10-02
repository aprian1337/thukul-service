package records

import (
	"aprian1337/thukul-service/business/pockets"
	"gorm.io/gorm"
	"time"
)

type Pockets struct {
	ID        int `gorm:"primaryKey"`
	UserId    int
	User      Users `gorm:"foreignKey:user_id"`
	Name      string
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type PocketsTotal struct {
	ID    int
	Total int64
}

func (data *Pockets) PocketsToDomain() pockets.Domain {
	return pockets.Domain{
		ID:        data.ID,
		UserId:    data.UserId,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func PocketsFromDomain(domain pockets.Domain) Pockets {
	return Pockets{
		ID:        domain.ID,
		UserId:    domain.UserId,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func PocketsToListDomain(data []Pockets) []pockets.Domain {
	var list []pockets.Domain
	for _, v := range data {
		list = append(list, v.PocketsToDomain())
	}
	return list
}
