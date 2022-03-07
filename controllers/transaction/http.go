package transaction

import (
	"net/http"
	"strconv"

	"github.com/daffashafwan/taskmajoo/business/merchants"
	"github.com/daffashafwan/taskmajoo/business/transactions"
	"github.com/daffashafwan/taskmajoo/controllers/transaction/requests"
	"github.com/daffashafwan/taskmajoo/controllers/transaction/responses"
	"github.com/daffashafwan/taskmajoo/helpers/response"
	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	TransactionUsecase transactions.Usecase
	MerchantUsecase merchants.Usecase
}

func NewTransactionController(transactionUsecase transactions.Usecase, merchantUsecase merchants.Usecase) *TransactionController {
	return &TransactionController{
		TransactionUsecase: transactionUsecase,
		MerchantUsecase: merchantUsecase,
	}
}

func (TransactionController TransactionController) GetByMerchantId(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	pagination := requests.Pagination{}
	c.Bind(&pagination)
	merchant,_ := TransactionController.MerchantUsecase.GetByUserId(ctx,convId)

	transaction, _ := TransactionController.TransactionUsecase.GetByMerchantId(ctx, merchant.Id, pagination.ToDomain())

	return response.SuccessResponse(c,http.StatusOK, responses.FromListDomain(transaction))
}

func (TransactionController TransactionController) GetByMerchantIdWithOutlet(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	pagination := requests.Pagination{}
	c.Bind(&pagination)
	merchant,_ := TransactionController.MerchantUsecase.GetByUserId(ctx,convId)

	transaction, _ := TransactionController.TransactionUsecase.GetByMerchantIdWithOutlet(ctx, merchant.Id, pagination.ToDomain())
	return response.SuccessResponse(c,http.StatusOK, responses.FromListDomainOutlet(transaction))
}