package outlet

import (
	"context"
	"errors"

	"github.com/daffashafwan/taskmajoo/business/outlets"
	"gorm.io/gorm"
)

type OutletRepo struct {
	DB *gorm.DB
}

func CreateOutletRepo(conn *gorm.DB) outlets.Repository {
	return &OutletRepo{
		DB: conn,
	}
}

func (rep *OutletRepo) GetById(ctx context.Context, id int) (outlets.Domain, error) {
	var data Outlets
	err := rep.DB.Table("Outlets").Find(&data, "id=?", id)
	if err.Error != nil {
		return outlets.Domain{}, errors.New("data not found")
	}
	return data.ToDomain(), nil
}


