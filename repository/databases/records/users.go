package records

import (
	"aprian1337/thukul-service/business/users"
	"gorm.io/gorm"
	"time"
)

type Users struct {
	ID        uint `gorm:"primaryKey"`
	SalaryId  int
	Salary    Salaries `gorm:"foreignKey:SalaryId"`
	Name      string
	Password  string
	IsAdmin   int `gorm:"type:smallint; default:0"`
	Email     string
	Phone     string `gorm:"size:18"`
	Gender    string `gorm:"size:8"`
	Birthday  string `gorm:"type:date"`
	Address   string `gorm:"type:text"`
	Company   string
	Wallets   []Wallets      `gorm:"foreignKey:user_id"`
	IsValid   int            `gorm:"type:smallint; default:0"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *Users) UsersToDomain() users.Domain {
	return users.Domain{
		ID:       user.ID,
		SalaryId: user.SalaryId,
		Salary: struct {
			ID      uint
			Minimal float64
			Maximal float64
		}{ID: user.Salary.ID, Minimal: user.Salary.Minimal, Maximal: user.Salary.Maximal},
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

func UsersFromDomain(domain *users.Domain) Users {
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

func UsersToListDomain(data []Users) []users.Domain {
	var list []users.Domain
	for _, v := range data {
		list = append(list, v.UsersToDomain())
	}
	return list
}
