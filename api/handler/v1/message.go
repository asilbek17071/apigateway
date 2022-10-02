package v1

import (
	"context"
	"net/http"
	"time"

	pb "github.com/asilbek17071/apigateway/genproto/messenger_service"
	l "github.com/asilbek17071/apigateway/pkg/logger"
	"github.com/asilbek17071/apigateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// MessageCreate ...
// @Summary MessageCreate
// @Router /message/create/ [post]
// @Description This API for creating a new message
// @Tags message
// @Accept  json
// @Produce  json
// @Param Message request body models.Message true "directionCreateRequest"
// @Success 200 {object} models.MessageResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) MessageCreate(c *gin.Context) {
	var (
		body        pb.Message
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

	response, err := h.serviceManager.MessengerService().MessageCreate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create direction", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// Notification ...
// @Router /message/notification/ [get]
// @Summary Notification
// @Description This API for getting list of messages
// @Tags message
// @Accept  json
// @Produce  json
// @Param id query string true "Id"
// @Param permission_receiver query string true "PermissionReceiver"
// @Success 200 {object} models.MessageList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) Notification(c *gin.Context) {
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

	response, err := h.serviceManager.MessengerService().Notification(
		ctx, &pb.ByIdReq{
			Id:                  params.Id,
			PersmissionReceiver: params.PermissionReceiver,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list spendings", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// SendList ...
// @Router /message/send/list [get]
// @Summary SendList
// @Description This API for getting list of messages
// @Tags message
// @Accept  json
// @Produce  json
// @Param id query string true "Id"
// @Param persmission_sender query string true "PermissionSender"
// @Success 200 {object} models.MessageList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) SendList(c *gin.Context) {
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

	response, err := h.serviceManager.MessengerService().SendList(
		ctx, &pb.Send{
			Id:                params.Id,
			PersmissionSender: params.PermissionSender,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list spendings", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// StatusUpdate ...
// @Router /message/status/update/{id} [put]
// @Summary StatusUpdate
// @Description This API for updating message
// @Tags message
// @Accept  json
// @Produce  json
// @Param id query string true "Id"
// @Param message_id query string true "MessageId"
// @Param persmission_sender query string true "PermissionSender"
// @Param persmission_receiver query string true "PermissionReceiver"
// @Success 200 {object} models.FinResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) StatusUpdate(c *gin.Context) {
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

	response, err := h.serviceManager.MessengerService().StatusUpdate(
		ctx, &pb.Status{
			Id:                 params.Id,
			MessageId:          params.MessageId,
			PermissionSender:   params.PermissionSender,
			PermissionReceiver: params.PermissionReceiver,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list spendings", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// MessageDelete ...
// @Router /message/delete/{id} [delete]
// @Summary MessageDelete
// @Description This API for deleting message
// @Tags message
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param persmission_sender query string true "PermissionSender"
// @Param persmission_receiver query string true "PermissionReceiver"
// @Success 200 {object} models.FinResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) MessageDelete(c *gin.Context) {
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

	response, err := h.serviceManager.MessengerService().MessageDelete(
		ctx, &pb.Del{
			MessageId:          params.Id,
			PermissionSender:   params.PermissionSender,
			PermissionReceiver: params.PermissionReceiver,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete spending", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
