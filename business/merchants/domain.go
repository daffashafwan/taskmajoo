package merchants

import (
	"context"
	"time"
)

type Domain struct {
	Id   int
	Name string 
	Username string 
	Password string 
	JWTToken string
	CreatedAt time.Time 
	UpdatedAt time.Time 
}

type Usecase interface {
	Login(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, user_name string, password string) (Domain, error)
}
