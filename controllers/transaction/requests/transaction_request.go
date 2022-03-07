package requests

import(
	"github.com/daffashafwan/taskmajoo/business/transactions"
)

type Pagination struct {
	Limit int   `json:"limit"`
	Page  int   `json:"page"`
}

func (ul *Pagination) ToDomain() transactions.Pagination {
	return transactions.Pagination{
		Limit: ul.Limit,
		Page: ul.Page,
	}
}