package pockets

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/pockets"
	"aprian1337/thukul-service/repository/databases/users"
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

func (repo *PostgresPocketsRepository) GetList(_ context.Context, id int) ([]pockets.Domain, error) {
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
func (repo *PostgresPocketsRepository) GetById(_ context.Context, userId int, pocketId int) (pockets.Domain, error) {
	var data Pockets
	err := repo.ConnPostgres.First(&data, "user_id = ? AND id = ?", userId, pocketId)
	if err.Error != nil {
		return pockets.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}
func (repo *PostgresPocketsRepository) Create(_ context.Context, domain pockets.Domain) (pockets.Domain, error) {
	pocket := FromDomain(domain)
	var user users.Users
	err := repo.ConnPostgres.First(&user, "id=?", pocket.UserId)
	if err.Error != nil {
		return pockets.Domain{}, businesses.ErrUserIdNotFound
	}
	err = repo.ConnPostgres.Create(&pocket)
	if err.Error != nil {
		return pockets.Domain{}, err.Error
	}
	return pocket.ToDomain(), nil
}
func (repo *PostgresPocketsRepository) Update(_ context.Context, domain pockets.Domain, userId int, pocketId int) (pockets.Domain, error) {
	data := FromDomain(domain)
	pocket := Pockets{}
	err := repo.ConnPostgres.First(&pocket, "user_id = ? AND id = ?", userId, pocketId)
	if err.Error != nil {
		return pockets.Domain{}, err.Error
	}
	pocket.Name = data.Name
	repo.ConnPostgres.Save(&pocket)
	return pocket.ToDomain(), nil
}
func (repo *PostgresPocketsRepository) Delete(_ context.Context, userId int, pocketId int) (int64, error) {
	data := Pockets{}
	err := repo.ConnPostgres.Delete(&data, "user_id = ? AND pocket_id = ?", userId, pocketId)
	if err.Error != nil {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
