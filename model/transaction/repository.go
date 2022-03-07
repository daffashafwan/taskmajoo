package transaction

import (
	"context"
	"errors"

	"github.com/daffashafwan/taskmajoo/business/transactions"
	"gorm.io/gorm"
)

type TransactionRepo struct {
	DB *gorm.DB
}

func CreateTransactionRepo(conn *gorm.DB) transactions.Repository {
	return &TransactionRepo{
		DB: conn,
	}
}

func (rep *TransactionRepo) GetByMerchantId(ctx context.Context, id int) ([]transactions.Domain, error) {
	var data []Transaction
	err := rep.DB.Table("Transactions").Select("Transactions.id, Transactions.merchant_id, Transactions.outlet_id, Transactions.bill_total, Transactions.updated_at, Merchants.merchant_name").Joins("left join Merchants on Merchants.id = Transactions.merchant_id").Find(&data, "merchant_id=?", id)
	if err.Error != nil {
		return []transactions.Domain{}, errors.New("data not found")
	}
	return ToListDomain(data), nil
}

func (rep *TransactionRepo) GetByMerchantIdWithOutlet(ctx context.Context, id int) ([]transactions.Domain, error) {
	var data []Transaction
	err := rep.DB.Table("Transactions").Select("Transactions.id, Transactions.merchant_id, Transactions.outlet_id, Transactions.bill_total, Transactions.updated_at, Merchants.merchant_name").Joins("left join Merchants on Merchants.id = Transactions.merchant_id").Find(&data, "merchant_id=?", id)
	if err.Error != nil {
		return []transactions.Domain{}, errors.New("data not found")
	}
	return ToListDomain(data), nil
}


