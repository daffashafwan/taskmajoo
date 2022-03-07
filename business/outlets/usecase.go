package outlets

import (
	"context"
	"time"

	"github.com/daffashafwan/taskmajoo/app/middlewares"
	errors "github.com/daffashafwan/taskmajoo/helpers/errors"
)

type OutletUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
	ConfigJWT      middlewares.ConfigJWT
}

func NewOutletUsecase(repo Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &OutletUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (pc *OutletUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	user, err := pc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if user.Id == 0 {
		return Domain{}, errors.ErrIDNotFound
	}
	return user, nil
}
