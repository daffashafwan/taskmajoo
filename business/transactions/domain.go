package transactions

import (
	"context"
	"time"
)

type Domain struct {
	Id   int
	MerchantId int
	OutletId int
	BillTotal float64 
	MerchantName string 
	UpdatedAt time.Time
}

type DomainRecap struct{
	MerchantName string
	Omzet float64
	Date string
}

type DomainRecapOutlet struct{
	MerchantName string
	OutletName string
	Omzet float64
	Date string
}

type Pagination struct{
	Limit int 
	Page  int  
}

type Usecase interface {
	GetByMerchantId(ctx context.Context, id int, pagination Pagination) ([]DomainRecap, error)
	GetByMerchantIdWithOutlet(ctx context.Context, id int, pagination Pagination) ([]DomainRecapOutlet, error)
}

type Repository interface {
	GetByMerchantId(ctx context.Context, id int) ([]Domain, error)
	GetByMerchantIdWithOutlet(ctx context.Context, id int) ([]Domain, error)
}
