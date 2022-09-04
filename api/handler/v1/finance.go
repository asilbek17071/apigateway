package v1

import (
	"context"
	"net/http"
	"time"

	pb "github.com/asilbek17071/apigateway/genproto/finance_service"
	l "github.com/asilbek17071/apigateway/pkg/logger"
	"github.com/asilbek17071/apigateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// SalaryFinanceCreate ...
// @Summary SalaryFinanceCreate
// @Router /finance/salary/create/ [post]
// @Description This API for creating a new finance
// @Tags finance
// @Accept  json
// @Produce  json
// @Param Finance request body models.Finance true "financeCreateRequest"
// @Success 200 {object} models.FinResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) SalaryFinanceCreate(c *gin.Context) {
	var (
		body        pb.EmptyReq
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.FinanceService().SalaryFinanceCreate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create direction", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// SalaryFinanceGet ...
// @Router /finance/salary/ [get]
// @Summary SalaryFinanceGet
// @Description This API for getting finance SalaryFinanceList
// @Tags finance
// @Accept  json
// @Produce  json
// @Param from path string true "From"
// @Success 200 {object} models.FinanceRespList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) SalaryFinanceGet(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.FinanceService().SalaryFinanceGet(
		ctx, &pb.Request{
			Date: params.From,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get direction", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
