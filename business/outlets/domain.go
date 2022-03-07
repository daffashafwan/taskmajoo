package outlets

import (
	"context"
)

type Domain struct {
	Id   int 
	OutletName string 
}

type Usecase interface {
	GetById(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	GetById(ctx context.Context, id int) (Domain, error)
}
