package favorites

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/coins"
	"context"
	"time"
)

type FavoriteUsecase struct {
	Repo     Repository
	RepoCoin coins.Repository
	Timeout  time.Duration
}

func NewFavoriteUsecase(repo Repository, repoCoin coins.Repository, timeout time.Duration) *FavoriteUsecase {
	return &FavoriteUsecase{
		Repo:     repo,
		RepoCoin: repoCoin,
		Timeout:  timeout,
	}
}

func (uc *FavoriteUsecase) GetList(ctx context.Context, userId int) ([]Domain, error) {
	data, err := uc.Repo.GetList(ctx, userId)
	if err != nil {
		return []Domain{}, err
	}
	return data, nil
}

func (uc *FavoriteUsecase) GetById(ctx context.Context, userId int, favoriteId int) (Domain, error) {
	data, err := uc.Repo.GetById(ctx, userId, favoriteId)
	if err != nil {
		return Domain{}, err
	}
	return data, nil
}

func (uc *FavoriteUsecase) Create(ctx context.Context, domain Domain, userId int) (Domain, error) {
	data, err := uc.Repo.Create(ctx, domain, userId)
	if err != nil {
		return Domain{}, err
	}
	return data, nil
}

func (uc *FavoriteUsecase) Delete(ctx context.Context, userId int, favoriteId int) error {
	rowsAffected, err := uc.Repo.Delete(ctx, userId, favoriteId)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return businesses.ErrNothingDestroy
	}
	return nil
}
