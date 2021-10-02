package records

import (
	"aprian1337/thukul-service/business/wishlists"
	"gorm.io/gorm"
	"gorm.io/gorm/utils/tests"
	"time"
)

type Wishlists struct {
	ID          int `gorm:"primarykey"`
	UserId      int
	User        tests.User `gorm:"foreignKey:user_id"`
	Name        string
	Nominal     float64
	TargetDate  string
	Priority    string `gorm:"size:6"`
	Note        string
	IsDone      int
	PicUrl      string
	WishlistUrl string
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (data *Wishlists) WishlistsToDomain() wishlists.Domain {
	return wishlists.Domain{
		ID:          data.ID,
		UserId:      data.UserId,
		Nominal:     data.Nominal,
		TargetDate:  data.TargetDate,
		Priority:    data.Priority,
		Note:        data.Note,
		Name:        data.Name,
		IsDone:      data.IsDone,
		PicUrl:      data.PicUrl,
		WishlistUrl: data.WishlistUrl,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func WishlistsFromDomain(domain wishlists.Domain) Wishlists {
	return Wishlists{
		ID:          domain.ID,
		UserId:      domain.UserId,
		Nominal:     domain.Nominal,
		TargetDate:  domain.TargetDate,
		Priority:    domain.Priority,
		Name:        domain.Name,
		Note:        domain.Note,
		IsDone:      domain.IsDone,
		PicUrl:      domain.PicUrl,
		WishlistUrl: domain.WishlistUrl,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func WishlistsToListDomain(data []Wishlists) []wishlists.Domain {
	var list []wishlists.Domain
	for _, v := range data {
		list = append(list, v.WishlistsToDomain())
	}
	return list
}
