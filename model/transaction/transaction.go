package transaction

import (
	"github.com/daffashafwan/taskmajoo/business/transactions"
	//"github.com/daffashafwan/taskmajoo/model/merchant"
	"time"
)

type Transaction struct {
	Id          int `gorm:"primaryKey"`
	Merchant_id int
	//Merchants    merchant.Merchants `gorm:"foreignKey:Merchant_id;association_foreignkey:Id"`
	Outlet_id     int
	Bill_total    float64
	Merchant_name string
	Updated_at time.Time
}

func (transaction *Transaction) ToDomain() transactions.Domain {
	return transactions.Domain{
		Id:         transaction.Id,
		MerchantId: transaction.Merchant_id,
		//Merchants: transaction.Merchants,
		MerchantName: transaction.Merchant_name,
		OutletId:     transaction.Outlet_id,
		BillTotal:    transaction.Bill_total,
		UpdatedAt: transaction.Updated_at,
	}
}

func FromDomain(domain transactions.Domain) Transaction {
	return Transaction{
		Id:          domain.Id,
		Merchant_id: domain.MerchantId,
		Outlet_id:   domain.OutletId,
		Bill_total:  domain.BillTotal,
	}
}

func ToListDomain(data []Transaction) (result []transactions.Domain) {
	result = []transactions.Domain{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}
