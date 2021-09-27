package activities

import (
	"aprian1337/thukul-service/business/activities"
	"gorm.io/gorm"
	"time"
)

type Activities struct {
	ID        int `gorm:"primaryKey"`
	PocketId  int
	Name      string
	Type      string
	Nominal   float64
	Note      string
	Date      string    `gorm:"type:date"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

type Total struct {
	ID    int
	Total int64
}

func (data *Activities) ToDomain() activities.Domain {
	return activities.Domain{
		ID:        data.ID,
		PocketId:  data.PocketId,
		Name:      data.Name,
		Type:      data.Type,
		Nominal:   data.Nominal,
		Note:      data.Note,
		Date:      data.Date,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func FromDomain(domain activities.Domain) Activities {
	return Activities{
		ID:        domain.ID,
		PocketId:  domain.PocketId,
		Name:      domain.Name,
		Type:      domain.Type,
		Nominal:   domain.Nominal,
		Note:      domain.Note,
		Date:      domain.Date,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ToListDomain(data []Activities) []activities.Domain {
	var list []activities.Domain
	for _, v := range data {
		list = append(list, v.ToDomain())
	}
	return list
}
