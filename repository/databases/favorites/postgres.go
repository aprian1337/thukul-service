package favorites

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/favorites"
	"aprian1337/thukul-service/repository/databases/users"
	"context"
	"gorm.io/gorm"
)

type PostgresFavoritesRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresFavoritesRepository(conn *gorm.DB) *PostgresFavoritesRepository {
	return &PostgresFavoritesRepository{
		ConnPostgres: conn,
	}
}

func (repo *PostgresFavoritesRepository) GetList(ctx context.Context, userId int) ([]favorites.Domain, error) {
	var data []Favorites
	err := repo.ConnPostgres.Find(&data, "user_id=?", userId)
	if err.Error != nil {
		return []favorites.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (repo *PostgresFavoritesRepository) GetById(ctx context.Context, userId int, wishlistId int) (favorites.Domain, error) {
	var data Favorites
	err := repo.ConnPostgres.First(&data, "user_id=? AND id=?", userId, wishlistId)
	if err.Error != nil {
		return favorites.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (repo *PostgresFavoritesRepository) Create(ctx context.Context, domain favorites.Domain, userId int) (favorites.Domain, error) {
	pocket := FromDomain(domain)
	var user users.Users
	err := repo.ConnPostgres.First(&user, "user_id=?", userId)
	if err.Error != nil {
		return favorites.Domain{}, businesses.ErrUserIdNotFound
	}
	err = repo.ConnPostgres.Create(&pocket)
	if err.Error != nil {
		return favorites.Domain{}, err.Error
	}
	return pocket.ToDomain(), nil
}

func (repo *PostgresFavoritesRepository) Delete(ctx context.Context, userId int, wishlistId int) (int64, error) {
	data := Favorites{}
	err := repo.ConnPostgres.Delete(&data, "user_id=? AND id=?", userId, wishlistId)
	if err.Error != nil {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
