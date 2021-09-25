package salaries

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

func DomainToSalaries(domain salaries.Domain) Salaries {
	return Salaries{
		ID:      domain.ID,
		Minimal: domain.Minimal,
		Maximal: domain.Maximal,
	}
}

func (sal *Salaries) ToDomain() salaries.Domain {
	return salaries.Domain{
		ID:        sal.ID,
		Minimal:   sal.Minimal,
		Maximal:   sal.Maximal,
		CreatedAt: sal.CreatedAt,
		UpdatedAt: sal.UpdatedAt,
	}
}

func ToListDomain(domain []Salaries) []salaries.Domain {
	var result []salaries.Domain
	for _, v := range domain {
		result = append(result, v.ToDomain())
	}
	return result
}
