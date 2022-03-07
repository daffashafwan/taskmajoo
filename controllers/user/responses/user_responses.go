package responses

import (
	"github.com/daffashafwan/taskmajoo/business/users"
)

type UserResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	JWTToken string `json:"jwtToken"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		Id:       domain.Id,
		Name:     domain.Name,
		Username: domain.Username,
		JWTToken: domain.JWTToken,
	}
}

func FromListDomain(domain []users.Domain) []UserResponse {
	var list []UserResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}
