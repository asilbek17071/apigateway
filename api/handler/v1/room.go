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

// RoomCreate ...
// @Summary RoomCreate
// @Router /room/create/ [post]
// @Description This API for creating a new room
// @Tags room
// @Accept  json
// @Produce  json
// @Param Room request body models.Room true "roomCreateRequest"
// @Success 200 {object} models.RoomResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) RoomCreate(c *gin.Context) {
	var (
		body        pb.Room
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

	response, err := h.serviceManager.SystemService().RoomCreate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create room", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// RoomGet ...
// @Router /room/byid/{id} [get]
// @Summary RoomGet
// @Description This API for getting room RoomList
// @Tags room
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.RoomResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) RoomGet(c *gin.Context) {
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

	response, err := h.serviceManager.SystemService().RoomGet(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get room", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// RoomList ...
// @Router /room/list/ [get]
// @Summary RoomList
// @Description This API for getting list of rooms
// @Tags room
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} models.RoomsList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) RoomList(c *gin.Context) {
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

	response, err := h.serviceManager.SystemService().RoomList(
		ctx, &pb.ListReq{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list rooms", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// RoomUpdate ...
// @Router /room/update/{id} [put]
// @Summary RoomUpdate
// @Description This API for updating room
// @Tags room
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param Room request body models.RoomUpdate true "roomUpdateRequest"
// @Success 200 {object} models.RoomResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) RoomUpdate(c *gin.Context) {
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
		body        pb.Room
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

	response, err := h.serviceManager.SystemService().RoomUpdate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update room", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// RoomDelete ...
// @Router /room/delete/{id} [delete]
// @Summary RoomDelete
// @Description This API for deleting room
// @Tags room
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Empty
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) RoomDelete(c *gin.Context) {
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

	response, err := h.serviceManager.SystemService().RoomDelete(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete room", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
