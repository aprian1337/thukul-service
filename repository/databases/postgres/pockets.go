package postgres

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/pockets"
	pockets2 "aprian1337/thukul-service/repository/databases/records"
	"context"
	"gorm.io/gorm"
)

type PocketsRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresPocketsRepository(conn *gorm.DB) *PocketsRepository {
	return &PocketsRepository{
		ConnPostgres: conn,
	}
}

func (repo *PocketsRepository) PocketsGetList(_ context.Context, id int) ([]pockets.Domain, error) {
	var data []pockets2.Pockets
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

	return pockets2.PocketsToListDomain(data), nil
}
func (repo *PocketsRepository) PocketsGetById(_ context.Context, userId int, pocketId int) (pockets.Domain, error) {
	var data pockets2.Pockets
	err := repo.ConnPostgres.First(&data, "user_id = ? AND id = ?", userId, pocketId)
	if err.Error != nil {
		return pockets.Domain{}, err.Error
	}
	return data.PocketsToDomain(), nil
}
func (repo *PocketsRepository) PocketsCreate(_ context.Context, domain pockets.Domain) (pockets.Domain, error) {
	pocket := pockets2.PocketsFromDomain(domain)
	var user pockets2.Users
	err := repo.ConnPostgres.First(&user, "id=?", pocket.UserId)
	if err.Error != nil {
		return pockets.Domain{}, businesses.ErrUserIdNotFound
	}
	err = repo.ConnPostgres.Create(&pocket)
	if err.Error != nil {
		return pockets.Domain{}, err.Error
	}
	return pocket.PocketsToDomain(), nil
}
func (repo *PocketsRepository) PocketsUpdate(_ context.Context, domain pockets.Domain, userId int, pocketId int) (pockets.Domain, error) {
	data := pockets2.PocketsFromDomain(domain)
	pocket := pockets2.Pockets{}
	err := repo.ConnPostgres.First(&pocket, "user_id = ? AND id = ?", userId, pocketId)
	if err.Error != nil {
		return pockets.Domain{}, err.Error
	}
	pocket.Name = data.Name
	repo.ConnPostgres.Save(&pocket)
	return pocket.PocketsToDomain(), nil
}
func (repo *PocketsRepository) PocketsDelete(_ context.Context, userId int, pocketId int) (int64, error) {
	data := pockets2.Pockets{}
	err := repo.ConnPostgres.Delete(&data, "user_id = ? AND pocket_id = ?", userId, pocketId)
	if err.Error != nil {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
