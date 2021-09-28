package favorites

import (
	"aprian1337/thukul-service/business/favorites"
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

func (repo *PostgresFavoritesRepository) Create(ctx context.Context, domain favorites.Domain) (favorites.Domain, error) {
	favorite := FromDomain(domain)
	err := repo.ConnPostgres.Create(&favorite)
	if err.Error != nil {
		return favorites.Domain{}, err.Error
	}
	return favorite.ToDomain(), nil
}

func (repo *PostgresFavoritesRepository) Delete(ctx context.Context, userId int, favoriteId int) (int64, error) {
	data := Favorites{}
	err := repo.ConnPostgres.Delete(&data, "user_id=? AND id=?", userId, favoriteId)
	if err.Error != nil {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
