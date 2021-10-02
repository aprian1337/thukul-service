package users

import (
	"aprian1337/thukul-service/business/users"
	"aprian1337/thukul-service/repository/databases/salaries"
	"gorm.io/gorm"
	"time"
)

type Users struct {
	ID       uint `gorm:"primaryKey"`
	SalaryId int
	Salary   salaries.Salaries `gorm:"foreignKey:SalaryId"`
	Name     string
	Password string
	IsAdmin  int `gorm:"type:smallint; default:0"`
	Email    string
	Phone    string `gorm:"size:18"`
	Gender   string `gorm:"size:8"`
	Birthday string `gorm:"type:date"`
	Address  string `gorm:"type:text"`
	Company  string
	//Wallets   *[]wallets.Wallets
	IsValid   int            `gorm:"type:smallint; default:0"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *Users) ToDomain() users.Domain {
	return users.Domain{
		ID:        user.ID,
		SalaryId:  user.SalaryId,
		Name:      user.Name,
		Password:  user.Password,
		IsAdmin:   user.IsAdmin,
		Email:     user.Email,
		Phone:     user.Phone,
		Gender:    user.Gender,
		Birthday:  user.Birthday,
		Address:   user.Address,
		Company:   user.Company,
		IsValid:   user.IsValid,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func FromDomain(domain *users.Domain) Users {
	return Users{
		ID:        domain.ID,
		SalaryId:  domain.SalaryId,
		Name:      domain.Name,
		Password:  domain.Password,
		IsAdmin:   domain.IsAdmin,
		Email:     domain.Email,
		Phone:     domain.Phone,
		Gender:    domain.Gender,
		Birthday:  domain.Birthday,
		Address:   domain.Address,
		Company:   domain.Company,
		IsValid:   domain.IsValid,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ToListDomain(data []Users) []users.Domain {
	list := []users.Domain{}
	for _, v := range data {
		list = append(list, v.ToDomain())
	}
	return list
}
