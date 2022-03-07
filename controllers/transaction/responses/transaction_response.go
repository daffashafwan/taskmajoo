package responses

import (
	"github.com/daffashafwan/taskmajoo/business/transactions"
)

type TransactionResponse struct {
	MerchantName   string `json:"merchant_name"`
	Omzet  float64     `json:"omzet"`
	Date string`json:"date"`
}

type TransactionResponseOutlet struct {
	MerchantName   string `json:"merchant_name"`
	Omzet  float64     `json:"omzet"`
	Outlet string `json:"outlet"`
	Date string`json:"date"`
}

func FromDomain(domain transactions.DomainRecap) TransactionResponse {
	return TransactionResponse{
		MerchantName: domain.MerchantName,
		Omzet: domain.Omzet,
		Date: domain.Date,
	}
}

func FromListDomain(domain []transactions.DomainRecap) []TransactionResponse {
	var list []TransactionResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}

func FromDomainOutlet(domain transactions.DomainRecapOutlet) TransactionResponseOutlet {
	return TransactionResponseOutlet{
		MerchantName: domain.MerchantName,
		Omzet: domain.Omzet,
		Outlet: domain.OutletName,
		Date: domain.Date,
	}
}

func FromListDomainOutlet(domain []transactions.DomainRecapOutlet) []TransactionResponseOutlet {
	var list []TransactionResponseOutlet
	for _, v := range domain {
		list = append(list, FromDomainOutlet(v))
	}
	return list
}
