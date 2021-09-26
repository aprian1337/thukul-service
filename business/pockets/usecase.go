package pockets

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/helpers"
	"context"
	"strconv"
	"time"
)

type PocketUsecase struct {
	Repo    Repository
	Timeout time.Duration
}

func NewPocketUsecase(repo Repository, timeout time.Duration) *PocketUsecase {
	return &PocketUsecase{
		Repo:    repo,
		Timeout: timeout,
	}
}

func (pu *PocketUsecase) GetList(ctx context.Context, id string) ([]Domain, error) {
	convId := 0
	if id != "" {
		if helpers.IsInt(id) == false {
			return []Domain{}, businesses.ErrInvalidId
		}
		convId, _ = strconv.Atoi(id)
	}

	pockets, err := pu.Repo.GetList(ctx, convId)
	if err != nil {
		return []Domain{}, err
	}
	return pockets, nil
}
func (pu *PocketUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	pockets, err := pu.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return pockets, nil
}
func (pu *PocketUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	if domain.UserId == 0 || domain.Name == "" {
		return Domain{}, businesses.ErrBadRequest
	}
	domain.TotalNominal = 0
	pockets, err := pu.Repo.Create(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return pockets, nil
}
func (pu *PocketUsecase) Update(ctx context.Context, id int, domain Domain) (Domain, error) {
	pockets, err := pu.Repo.Update(ctx, id, domain)
	if err != nil {
		return Domain{}, err
	}
	return pockets, nil
}

func (pu *PocketUsecase) Delete(ctx context.Context, id int) error {
	err := pu.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
