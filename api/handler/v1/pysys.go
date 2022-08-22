package v1

import (
	"context"
	"net/http"
	"time"

	pb "github.com/asilbek17071/apigateway/genproto/system_service"
	l "github.com/asilbek17071/apigateway/pkg/logger"
	"github.com/asilbek17071/apigateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// PySysCreate ...
// @Summary PySysCreate
// @Router /payment_system/create/ [post]
// @Description This API for creating a new payment_system
// @Tags payment_system
// @Accept  json
// @Produce  json
// @Param PySys request body models.PySys true "payment_systemCreateRequest"
// @Success 200 {object} models.PySysResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) PySysCreate(c *gin.Context) {
	var (
		body        pb.PySys
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

	response, err := h.serviceManager.SystemService().PySysCreate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create payment_system", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// PySysGet ...
// @Router /payment_system/list/{id} [get]
// @Summary PySysGet
// @Description This API for getting payment_system PySysList
// @Tags payment_system
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.PySysResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) PySysGet(c *gin.Context) {
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

	response, err := h.serviceManager.SystemService().PySysGet(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get payment_system", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// PySysList ...
// @Router /payment_system/list/ [get]
// @Summary PySysList
// @Description This API for getting list of payment_systems
// @Tags payment_system
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} models.PySyssList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) PySysList(c *gin.Context) {
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

	response, err := h.serviceManager.SystemService().PySysList(
		ctx, &pb.ListReq{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list payment_systems", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// PySysUpdate ...
// @Router /payment_system/update/ [put]
// @Summary PySysUpdate
// @Description This API for updating payment_system
// @Tags payment_system
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param PySys request body models.PySysUpdate true "payment_systemUpdateRequest"
// @Success 200 {object} models.PySysResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) PySysUpdate(c *gin.Context) {
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
		body        pb.PySys
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

	response, err := h.serviceManager.SystemService().PySysUpdate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update payment_system", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// PySysDelete ...
// @Router /payment_system/delete/{id} [delete]
// @Summary PySysDelete
// @Description This API for deleting payment_system
// @Tags payment_system
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Empty
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) PySysDelete(c *gin.Context) {
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

	response, err := h.serviceManager.SystemService().PySysDelete(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete payment_system", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
