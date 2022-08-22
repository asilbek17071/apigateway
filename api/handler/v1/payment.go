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

// PaymentCreate ...
// @Summary PaymentCreate
// @Router /payment/create/ [post]
// @Description This API for creating a new payment
// @Tags payment
// @Accept  json
// @Produce  json
// @Param Payment request body models.Payment true "paymentCreateRequest"
// @Success 200 {object} models.PaymentResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) PaymentCreate(c *gin.Context) {
	var (
		body        pb.Payment
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

	response, err := h.serviceManager.FinanceService().PaymentCreate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create payment", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// PaymentGet ...
// @Router /payment/byid/{id} [get]
// @Summary PaymentGet
// @Description This API for getting payment PaymentList
// @Tags payment
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.PaymentResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) PaymentGet(c *gin.Context) {
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

	response, err := h.serviceManager.FinanceService().PaymentGet(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get payment", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// PaymentList ...
// @Router /payment/list/ [get]
// @Summary PaymentList
// @Description This API for getting list of payments
// @Tags payment
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} models.PaymentsList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) PaymentList(c *gin.Context) {
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

	response, err := h.serviceManager.FinanceService().PaymentList(
		ctx, &pb.ListReq{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list payments", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// PaymentUpdate ...
// @Router /payment/update/{id} [put]
// @Summary PaymentUpdate
// @Description This API for updating payment
// @Tags payment
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param Payment request body models.PaymentUpdate true "paymentUpdateRequest"
// @Success 200 {object} models.PaymentResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) PaymentUpdate(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}
	var (
		body        pb.Payment
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

	body.Id = params.Id

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.FinanceService().PaymentUpdate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update payment", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// PaymentDelete ...
// @Router /payment/delete/{id} [delete]
// @Summary PaymentDelete
// @Description This API for deleting payment
// @Tags payment
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Empty
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) PaymentDelete(c *gin.Context) {
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

	response, err := h.serviceManager.FinanceService().PaymentDelete(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete payment", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// PaymentSearchList ...
// @Router /payment/searchlist/ [get]
// @Summary PaymentSearchList
// @Description This API for getting list of payments
// @Tags payment
// @Accept  json
// @Produce  json
// @Param id query string false "From"
// @Param from query string false "From"
// @Param to query string false "To"
// @Success 200 {object} models.PaymentsList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) PaymentSearchList(c *gin.Context) {
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

	response, err := h.serviceManager.FinanceService().PaymentSearchList(
		ctx, &pb.SearchReq{
			Dan:   params.From,
			Gacha: params.To,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list payments", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *handlerV1) PaymentSearch(c *gin.Context) {
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

	response, err := h.serviceManager.FinanceService().PaymentSearch(
		ctx, &pb.PaymentSearchReq{
			Id:    params.Id,
			Dan:   params.From,
			Gacha: params.To,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list payments", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
