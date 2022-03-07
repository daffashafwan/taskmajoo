package transactions

import (
	"context"
	"time"

	"github.com/daffashafwan/taskmajoo/app/middlewares"
	"github.com/daffashafwan/taskmajoo/business/merchants"
	"github.com/daffashafwan/taskmajoo/business/outlets"
)

type TransactionUsecase struct {
	OutletRepo     outlets.Repository
	MerchantRepo   merchants.Repository
	Repo           Repository
	contextTimeout time.Duration
	ConfigJWT      middlewares.ConfigJWT
}

func NewTransactionUsecase(repo Repository, merchantRepo merchants.Repository, outletRepo outlets.Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &TransactionUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		MerchantRepo:   merchantRepo,
		OutletRepo:     outletRepo,
		contextTimeout: timeout,
	}
}

func (pc *TransactionUsecase) GetByMerchantId(ctx context.Context, id int, pagination Pagination) ([]DomainRecap, error) {
	domainRecap := []DomainRecap{}
	merchant, err := pc.MerchantRepo.GetByUserId(ctx, id)
	transaction, err := pc.Repo.GetByMerchantId(ctx, id)
	start := time.Date(2021, time.November, 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)
	for d := start; d.After(end) == false; d = d.AddDate(0, 0, 1) {
		temp := DomainRecap{}
		for _, user := range transaction {
			if user.UpdatedAt.Format("2006-01-02") == d.Format("2006-01-02") {
				temp.Date = user.UpdatedAt.Format("2006-01-02")
				temp.Omzet = temp.Omzet + user.BillTotal
			}
		}
		if temp.Omzet == 0 {
			temp.Date = d.Format("2006-01-02")
		}
		temp.MerchantName = merchant.MerchantName
		domainRecap = append(domainRecap, temp)
	}
	if err != nil {
		return []DomainRecap{}, err
	}
	firstEntry := (pagination.Page - 1) * pagination.Limit
	lastEntry := firstEntry + pagination.Limit

	if lastEntry > len(domainRecap) {
		lastEntry = len(domainRecap)
	}
	return domainRecap[firstEntry:lastEntry], nil
}

func (pc *TransactionUsecase) GetByMerchantIdWithOutlet(ctx context.Context, id int, pagination Pagination) ([]DomainRecapOutlet, error) {
	domainRecapOutlet := []DomainRecapOutlet{}
	merchant, err := pc.MerchantRepo.GetByUserId(ctx, id)
	transaction, err := pc.Repo.GetByMerchantId(ctx, id)
	start := time.Date(2021, time.November, 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)
	for d := start; d.After(end) == false; d = d.AddDate(0, 0, 1) {
		temp := DomainRecapOutlet{}
		for _, user := range transaction {
			if user.UpdatedAt.Format("2006-01-02") == d.Format("2006-01-02") {
				temp.Date = user.UpdatedAt.Format("2006-01-02")
				outlet,_ := pc.OutletRepo.GetById(ctx, user.OutletId)
				temp.OutletName = outlet.OutletName
				temp.Omzet = temp.Omzet + user.BillTotal
			}
		}
		if temp.Omzet == 0 {
			temp.OutletName = ""
			temp.Date = d.Format("2006-01-02")
		}
		temp.MerchantName = merchant.MerchantName
		domainRecapOutlet = append(domainRecapOutlet, temp)
	}
	if err != nil {
		return []DomainRecapOutlet{}, err
	}
	firstEntry := (pagination.Page - 1) * pagination.Limit
	lastEntry := firstEntry + pagination.Limit

	if lastEntry > len(domainRecapOutlet) {
		lastEntry = len(domainRecapOutlet)
	}
	return domainRecapOutlet[firstEntry:lastEntry], nil
}
