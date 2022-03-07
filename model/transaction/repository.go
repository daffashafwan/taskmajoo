package transaction

import (
	"context"
	"errors"

	"github.com/daffashafwan/taskmajoo/business/merchants"
	"gorm.io/gorm"
)

type MerchantRepo struct {
	DB *gorm.DB
}

func CreateMerchantRepo(conn *gorm.DB) merchants.Repository {
	return &MerchantRepo{
		DB: conn,
	}
}

func (rep *MerchantRepo) GetByUserId(ctx context.Context, id int) (merchants.Domain, error) {
	var data Merchant
	err := rep.DB.Table("Merchants").Find(&data, "user_id=?", id)
	if err.Error != nil {
		return merchants.Domain{}, errors.New("data not found")
	}
	return data.ToDomain(), nil
}


