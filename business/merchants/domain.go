package merchants

import (
	"context"
)

type Domain struct {
	Id   int
	UserId int 
	MerchantName string 
}

type Usecase interface {
	GetByUserId(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	GetByUserId(ctx context.Context, id int) (Domain, error)
}
