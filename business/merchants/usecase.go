package merchants

import (
	"context"
	"time"

	"github.com/daffashafwan/taskmajoo/app/middlewares"
	errors "github.com/daffashafwan/taskmajoo/helpers/errors"
)

type MerchantUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
	ConfigJWT      middlewares.ConfigJWT
}

func NewMerchantUsecase(repo Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &MerchantUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (pc *MerchantUsecase) GetByUserId(ctx context.Context, id int) (Domain, error) {
	user, err := pc.Repo.GetByUserId(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if user.Id == 0 {
		return Domain{}, errors.ErrIDNotFound
	}
	return user, nil
}
