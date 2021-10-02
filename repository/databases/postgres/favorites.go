package postgres

import (
	"aprian1337/thukul-service/business/favorites"
	favorites2 "aprian1337/thukul-service/repository/databases/records"
	"context"
	"gorm.io/gorm"
)

type FavoritesRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresFavoritesRepository(conn *gorm.DB) *FavoritesRepository {
	return &FavoritesRepository{
		ConnPostgres: conn,
	}
}

func (repo *FavoritesRepository) FavoritesGetList(ctx context.Context, userId int) ([]favorites.Domain, error) {
	var data []favorites2.Favorites
	err := repo.ConnPostgres.Joins("Coin").Find(&data, "user_id=?", userId)
	if err.Error != nil {
		return []favorites.Domain{}, err.Error
	}
	return favorites2.FavoritesToListDomain(data), nil
}

func (repo *FavoritesRepository) FavoritesGetById(ctx context.Context, userId int, wishlistId int) (favorites.Domain, error) {
	var data favorites2.Favorites
	err := repo.ConnPostgres.Joins("Coin").First(&data, "user_id=? AND favorites.id=?", userId, wishlistId)
	if err.Error != nil {
		return favorites.Domain{}, err.Error
	}
	return data.FavoritesToDomain(), nil
}

func (repo *FavoritesRepository) FavoritesCreate(ctx context.Context, domain favorites.Domain) (favorites.Domain, error) {
	favorite := favorites2.FavoritesFromDomain(domain)
	err := repo.ConnPostgres.Create(&favorite)
	if err.Error != nil {
		return favorites.Domain{}, err.Error
	}
	return favorite.FavoritesToDomain(), nil
}

func (repo *FavoritesRepository) FavoritesDelete(ctx context.Context, userId int, favoriteId int) (int64, error) {
	data := favorites2.Favorites{}
	err := repo.ConnPostgres.Delete(&data, "user_id=? AND id=?", userId, favoriteId)
	if err.Error != nil {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func (repo *FavoritesRepository) FavoritesCheck(ctx context.Context, userId int, coinId int) (int64, error) {
	var count int64
	err := repo.ConnPostgres.Model(&favorites2.Favorites{}).Where("user_id = ? AND coin_id=?", userId, coinId).Count(&count)
	if err.Error != nil {
		return 0, err.Error
	}
	return count, nil
}
