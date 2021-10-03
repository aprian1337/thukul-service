package salaries

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/helpers"
	"context"
	"time"
)

type SalaryUsecase struct {
	Repo    Repository
	Timeout time.Duration
}

func NewSalaryUsecase(repo Repository, timeout time.Duration) *SalaryUsecase {
	return &SalaryUsecase{
		Repo:    repo,
		Timeout: timeout,
	}
}

func (su *SalaryUsecase) GetList(ctx context.Context, search string) ([]Domain, error) {
	salaries, err := su.Repo.SalariesGetList(ctx, search)
	if err != nil {
		return []Domain{}, err
	}
	return salaries, nil
}
func (su *SalaryUsecase) GetById(ctx context.Context, id uint) (Domain, error) {
	salaries, err := su.Repo.SalariesGetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return salaries, nil
}
func (su *SalaryUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	salaries, err := su.Repo.SalariesCreate(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return salaries, nil
}
func (su *SalaryUsecase) Update(ctx context.Context, id uint, domain Domain) (Domain, error) {
	if helpers.IsZero(domain.Minimal) || helpers.IsZero(domain.Maximal) {
		return Domain{}, businesses.ErrBadRequest
	}
	domain.ID = id
	salaries, err := su.Repo.SalariesUpdate(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return salaries, nil
}

func (su *SalaryUsecase) Delete(ctx context.Context, id uint) error {
	rowAffected, err := su.Repo.SalariesDelete(ctx, id)
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return businesses.ErrNothingDestroy
	}
	return nil
}
