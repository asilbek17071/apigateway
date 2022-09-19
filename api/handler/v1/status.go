package v1

import (
	"context"
	"net/http"
	"time"

	pb "github.com/asilbek17071/apigateway/genproto/finance_service"
	l "github.com/asilbek17071/apigateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// StatusTeacherSalary ...
// @Router /status/teacher_salary/ [put]
// @Summary StatusTeacherSalary
// @Description This API for updating status
// @Tags status
// @Accept  json
// @Produce  json
// @Param Status request body models.StatusReq true "statusUpdateRequest"
// @Success 200 {object} models.FinResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) StatusTeacherSalary(c *gin.Context) {
	var (
		body        pb.StatusReq
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

	response, err := h.serviceManager.FinanceService().StatusTeacherSalary(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update spending", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// StatusAdminSalary ...
// @Router /status/admin_salary/ [put]
// @Summary StatusAdminSalary
// @Description This API for updating status
// @Tags status
// @Accept  json
// @Produce  json
// @Param Status request body models.StatusReq true "statusUpdateRequest"
// @Success 200 {object} models.FinResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) StatusAdminSalary(c *gin.Context) {
	var (
		body        pb.StatusReq
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

	response, err := h.serviceManager.FinanceService().StatusAdminSalary(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update spending", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// StatusPersonalSalary ...
// @Router /status/personal_salary/ [put]
// @Summary StatusPersonalSalary
// @Description This API for updating status
// @Tags status
// @Accept  json
// @Produce  json
// @Param Status request body models.StatusReq true "statusUpdateRequest"
// @Success 200 {object} models.FinResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) StatusPersonalSalary(c *gin.Context) {
	var (
		body        pb.StatusReq
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

	response, err := h.serviceManager.FinanceService().StatusPersonalSalary(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update spending", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// StatusSpending ...
// @Router /status/spending/ [put]
// @Summary StatusSpending
// @Description This API for updating status
// @Tags status
// @Accept  json
// @Produce  json
// @Param Status request body models.StatusSpendingReq true "statusUpdateRequest"
// @Success 200 {object} models.FinResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) StatusSpending(c *gin.Context) {
	var (
		body        pb.StatusSpendingReq
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

	response, err := h.serviceManager.FinanceService().StatusSpending(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update spending", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
