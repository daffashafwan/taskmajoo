package merchant

import (
	"net/http"
	"strconv"
	"github.com/daffashafwan/taskmajoo/business/merchants"
	"github.com/daffashafwan/taskmajoo/controllers/merchant/responses"
	"github.com/daffashafwan/taskmajoo/helpers/response"
	"github.com/labstack/echo/v4"
)

type MerchantController struct {
	MerchantUsecase merchants.Usecase
}

func NewMerchantController(merchantUsecase merchants.Usecase) *MerchantController {
	return &MerchantController{
		MerchantUsecase: merchantUsecase,
	}
}

func (MerchantController MerchantController) GetByUserId(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	merchant, _ := MerchantController.MerchantUsecase.GetByUserId(ctx, convId)
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(merchant))
}