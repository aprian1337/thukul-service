package activities

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/helpers"
	"context"
	"fmt"
	"time"
)

type ActivityUsecase struct {
	Repo    Repository
	Timeout time.Duration
}

func NewActivityUsecase(repo Repository, timeout time.Duration) *ActivityUsecase {
	return &ActivityUsecase{
		Repo:    repo,
		Timeout: timeout,
	}
}

func (uc *ActivityUsecase) GetList(ctx context.Context, pocketId int) ([]Domain, error) {
	pockets, err := uc.Repo.ActivitiesGetAll(ctx, pocketId)
	if err != nil {
		return []Domain{}, err
	}
	return pockets, nil
}

func (uc *ActivityUsecase) GetById(ctx context.Context, pocketId int, id int) (Domain, error) {
	data, err := uc.Repo.ActivitiesGetById(ctx, pocketId, id)
	if err != nil {
		return Domain{}, err
	}
	return data, nil
}

func (uc *ActivityUsecase) Create(ctx context.Context, domain Domain, pocketId int) (Domain, error) {
	if domain.Name == "" || domain.Type == "" || domain.Nominal == 0 || domain.Date == "" {
		return Domain{}, businesses.ErrBadRequest
	}
	if domain.Type != "income" && domain.Type != "expense" {
		return Domain{}, businesses.ErrTypeActivity
	}
	if !helpers.IsDate(domain.Date) {
		return Domain{}, businesses.ErrInvalidDate
	}

	fmt.Println(pocketId)
	pockets, err := uc.Repo.ActivitiesCreate(ctx, domain, pocketId)
	if err != nil {
		return Domain{}, err
	}
	return pockets, nil
}

func (uc *ActivityUsecase) Update(ctx context.Context, domain Domain, pocketId int, id int) (Domain, error) {
	if !helpers.IsDate(domain.Date) {
		return Domain{}, businesses.ErrInvalidDate
	}
	pockets, err := uc.Repo.ActivitiesUpdate(ctx, domain, pocketId, id)
	if err != nil {
		return Domain{}, err
	}
	return pockets, nil
}

func (uc *ActivityUsecase) Delete(ctx context.Context, id int, pocketId int) error {
	rowsAffected, err := uc.Repo.ActivitiesDelete(ctx, id, pocketId)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return businesses.ErrNothingDestroy
	}
	return nil
}

func (uc *ActivityUsecase) GetTotal(ctx context.Context, userId int, pocketId int, kind string) (int64, error) {
	total, err := uc.Repo.ActivitiesGetTotal(ctx, userId, pocketId, kind)
	if err != nil {
		return 0, err
	}
	return total, nil
}
