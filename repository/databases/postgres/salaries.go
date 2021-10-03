package postgres

import (
	"aprian1337/thukul-service/business/salaries"
	"aprian1337/thukul-service/repository/databases/records"
	"context"
	"gorm.io/gorm"
)

type SalariesRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresSalariesRepository(conn *gorm.DB) *SalariesRepository {
	return &SalariesRepository{
		ConnPostgres: conn,
	}
}

func (repo *SalariesRepository) SalariesGetList(_ context.Context, search string) ([]salaries.Domain, error) {
	var data []records.Salaries
	err := repo.ConnPostgres.Find(&data)
	if err.Error != nil {
		return []salaries.Domain{}, err.Error
	}
	return records.SalariesToListDomain(data), nil
}

func (repo *SalariesRepository) SalariesGetById(_ context.Context, id uint) (salaries.Domain, error) {
	var data records.Salaries
	err := repo.ConnPostgres.First(&data, "id=?", id)
	if err.Error != nil {
		return salaries.Domain{}, err.Error
	}
	return data.SalariesToDomain(), nil
}

func (repo *SalariesRepository) SalariesCreate(_ context.Context, domain salaries.Domain) (salaries.Domain, error) {
	salary := records.SalariesFromDomain(domain)
	err := repo.ConnPostgres.Create(&salary)
	if err.Error != nil {
		return salaries.Domain{}, err.Error
	}
	return salary.SalariesToDomain(), nil
}
func (repo *SalariesRepository) SalariesUpdate(_ context.Context, domain salaries.Domain) (salaries.Domain, error) {
	salary := records.SalariesFromDomain(domain)
	err := repo.ConnPostgres.First(&salary)
	if err.Error != nil {
		return salaries.Domain{}, err.Error
	}
	salary.Minimal = domain.Minimal
	salary.Maximal = domain.Maximal
	repo.ConnPostgres.Save(&salary)
	return salary.SalariesToDomain(), nil
}

func (repo *SalariesRepository) SalariesDelete(_ context.Context, id uint) (int64, error) {
	salary := records.Salaries{}
	err := repo.ConnPostgres.Delete(&salary, id)
	if err.Error != nil {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
