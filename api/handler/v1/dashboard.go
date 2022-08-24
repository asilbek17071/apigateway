package v1

import (
	"context"
	"net/http"
	"time"

	l "github.com/asilbek17071/apigateway/pkg/logger"

	pbf "github.com/asilbek17071/apigateway/genproto/finance_service"
	pb "github.com/asilbek17071/apigateway/genproto/student_service"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// StudentActive ...
// @Router /dashboard/student/ [get]
// @Summary StudentActive
// @Description This API for getting list of studentActive
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Success 200 {object} models.RespList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) StudentActive(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.StudentService().StudentActive(ctx, &pb.EmptyRes{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list dashboard", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// StudentTopList ...
// @Router /dashboard/toplist/ [get]
// @Summary StudentTopList
// @Description This API for getting list of studentTopList
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Success 200 {object} models.RespList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) StudentTopList(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.StudentService().StudentTopList(ctx, &pb.EmptyRes{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list dashboard", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// StudentDeleteList ...
// @Router /dashboard/deletelist/ [get]
// @Summary StudentDeleteList
// @Description This API for getting list of studentDeleteList
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Success 200 {object} models.RespList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) StudentDeleteList(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.StudentService().StudentDeleteList(ctx, &pb.EmptyRes{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list dashboard", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// StudentProfit ...
// @Router /dashboard/student_profit/ [get]
// @Summary StudentProfit
// @Description This API for getting list of studentProfit
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ProfitRespList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) StudentProfit(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.StudentService().StudentProfit(ctx, &pb.EmptyRes{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list dashboard", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// LeadActive ...
// @Router /dashboard/lead/ [get]
// @Summary LeadActive
// @Description This API for getting list of leadactive
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Success 200 {object} models.RespList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) LeadActive(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.StudentService().LeadActive(ctx, &pb.EmptyRes{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list dashboard", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Profit ...
// @Router /dashboard/profit/ [get]
// @Summary Profit
// @Description This API for getting list of profit
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ProfitRespList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) Profit(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.FinanceService().Profit(ctx, &pbf.EmptyRes{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list dashboard", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
