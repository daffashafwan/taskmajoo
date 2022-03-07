package merchants

import (
	"context"
	"time"

	"github.com/daffashafwan/taskmajoo/app/middlewares"
	errors "github.com/daffashafwan/taskmajoo/helpers/errors"
)

type UserUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
	ConfigJWT      middlewares.ConfigJWT
}

func NewUserUsecase(repo Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &UserUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

// core bisinis login
func (uc *UserUsecase) Login(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Username == "" {
		return Domain{}, errors.ErrUsernamePasswordNotFound
	}

	if domain.Password == "" {
		return Domain{}, errors.ErrUsernamePasswordNotFound
	}

	user, err := uc.Repo.Login(ctx, domain.Username, domain.Password)

	if err != nil {
		return Domain{}, err
	}
	user.JWTToken, err = uc.ConfigJWT.GenerateTokenJWT(user.Id, 0)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
