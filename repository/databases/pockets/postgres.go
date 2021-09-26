package pockets

import (
	"aprian1337/thukul-service/business/pockets"
	"context"
	"gorm.io/gorm"
)

type PostgresPocketsRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresPocketsRepository(conn *gorm.DB) *PostgresPocketsRepository {
	return &PostgresPocketsRepository{
		ConnPostgres: conn,
	}
}

func (repo *PostgresPocketsRepository) GetList(ctx context.Context, id int) ([]pockets.Domain, error) {
	var data []Pockets
	if id > 0 {
		err := repo.ConnPostgres.Find(&data, "user_id=?", id)
		if err.Error != nil {
			return []pockets.Domain{}, err.Error
		}
	} else {
		err := repo.ConnPostgres.Find(&data)
		if err.Error != nil {
			return []pockets.Domain{}, err.Error
		}
	}

	return ToListDomain(data), nil
}
func (repo *PostgresPocketsRepository) GetById(ctx context.Context, id int) (pockets.Domain, error) {
	var data Pockets
	err := repo.ConnPostgres.First(&data, "id=?", id)
	if err.Error != nil {
		return pockets.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}
func (repo *PostgresPocketsRepository) Create(ctx context.Context, domain pockets.Domain) (pockets.Domain, error) {
	pocket := FromDomain(domain)
	err := repo.ConnPostgres.Create(&pocket)
	if err.Error != nil {
		return pockets.Domain{}, err.Error
	}
	return pocket.ToDomain(), nil
}
func (repo *PostgresPocketsRepository) Update(ctx context.Context, id int, domain pockets.Domain) (pockets.Domain, error) {
	data := FromDomain(domain)
	data.ID = id
	err := repo.ConnPostgres.First(&data)
	if err.Error != nil {
		return pockets.Domain{}, err.Error
	}
	repo.ConnPostgres.Save(&data)
	return data.ToDomain(), nil
}
func (repo *PostgresPocketsRepository) Delete(ctx context.Context, id int) error {
	data := Pockets{}
	err := repo.ConnPostgres.Delete(&data, id)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
