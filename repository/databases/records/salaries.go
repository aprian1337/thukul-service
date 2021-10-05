package records

import (
	"aprian1337/thukul-service/business/salaries"
	"gorm.io/gorm"
	"time"
)

type Salaries struct {
	ID        uint `gorm:"primaryKey"`
	Minimal   float64
	Maximal   float64
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func SalariesFromDomain(domain salaries.Domain) Salaries {
	return Salaries{
		ID:      domain.ID,
		Minimal: domain.Minimal,
		Maximal: domain.Maximal,
	}
}

func (sal *Salaries) SalariesToDomain() salaries.Domain {
	return salaries.Domain{
		ID:        sal.ID,
		Minimal:   sal.Minimal,
		Maximal:   sal.Maximal,
		CreatedAt: sal.CreatedAt,
		UpdatedAt: sal.UpdatedAt,
	}
}

func SalariesToListDomain(domain []Salaries) []salaries.Domain {
	var result []salaries.Domain
	for _, v := range domain {
		result = append(result, v.SalariesToDomain())
	}
	return result
}
