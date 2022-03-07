package responses

import (
	"github.com/daffashafwan/taskmajoo/business/merchants"
)

type MerchantResponse struct {
	Id           int    `json:id`
	UserId       int    `json:"userId"`
	MerchantName string `json:"merchantName"`
}

func FromDomain(domain merchants.Domain) MerchantResponse {
	return MerchantResponse{
		UserId: domain.UserId,
		Id: domain.Id,
		MerchantName: domain.MerchantName,
	}
}
