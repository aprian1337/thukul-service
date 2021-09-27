package activities

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/activities"
	"aprian1337/thukul-service/repository"
	"aprian1337/thukul-service/repository/databases/pockets"
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
func (repo *PostgresPocketsRepository) GetList(ctx context.Context, pocketId int) ([]activities.Domain, error) {
	var data []Activities
	if pocketId > 0 {
		err := repo.ConnPostgres.Find(&data, "pocket_id=?", pocketId)
		if err.Error != nil {
			return []activities.Domain{}, err.Error
		}
	} else {
		err := repo.ConnPostgres.Find(&data)
		if err.Error != nil {
			return []activities.Domain{}, err.Error
		}
	}

	return ToListDomain(data), nil
}

func (repo *PostgresPocketsRepository) GetById(ctx context.Context, pocketId int, id int) (activities.Domain, error) {
	var data Activities
	err := repo.ConnPostgres.First(&data, "id=? AND pocket_id=?", id)
	if err.Error != nil {
		return activities.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (repo *PostgresPocketsRepository) Create(ctx context.Context, domain activities.Domain, pocketId int) (activities.Domain, error) {
	data := FromDomain(domain)
	var pocket pockets.Pockets
	err := repo.ConnPostgres.First(&pocket, "id=?", pocketId)
	if err.Error != nil {
		return activities.Domain{}, businesses.ErrUserIdNotFound
	}
	err = repo.ConnPostgres.Create(&data)
	if err.Error != nil {
		return activities.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (repo *PostgresPocketsRepository) Update(ctx context.Context, domain activities.Domain, pocketId int, id int) (activities.Domain, error) {
	data := Activities{}
	dataTemp := FromDomain(domain)
	err := repo.ConnPostgres.First(&data, "id=?", id)
	if err.Error != nil {
		return activities.Domain{}, err.Error
	}
	repo.ConnPostgres.Save(&dataTemp)
	return data.ToDomain(), nil
}

func (repo *PostgresPocketsRepository) Delete(ctx context.Context, pocketId int, id int) (int64, error) {
	data := Activities{}
	err := repo.ConnPostgres.First(&data, "id=? AND pocket_id=?", id, pocketId)
	if err.Error != nil {
		return 0, err.Error
	}
	if err.RowsAffected == 0 {
		return 0, repository.ErrDataNotFound
	}

	err = repo.ConnPostgres.Delete(&data, id)
	if err.Error != nil {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
