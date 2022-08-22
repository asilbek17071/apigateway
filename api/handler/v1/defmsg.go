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

// DefMsgCreate ...
// @Summary DefMsgCreate
// @Router /defmsg/create/ [post]
// @Description This API for creating a new defmsg
// @Tags defmsg
// @Accept  json
// @Produce  json
// @Param DefMsg request body models.DefMsg true "defmsgCreateRequest"
// @Success 200 {object} models.DefMsgResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) DefMsgCreate(c *gin.Context) {
	var (
		body        pb.DefMsg
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

	response, err := h.serviceManager.SystemService().DefMsgCreate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create defmsg", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// DefMsgGet ...
// @Router /defmsg/byid/{id} [get]
// @Summary DefMsgGet
// @Description This API for getting defmsg DefMsgList
// @Tags defmsg
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.DefMsgResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) DefMsgGet(c *gin.Context) {
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

	response, err := h.serviceManager.SystemService().DefMsgGet(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get defmsg", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DefMsgList ...
// @Router /defmsg/list/ [get]
// @Summary DefMsgList
// @Description This API for getting list of defmsgs
// @Tags defmsg
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} models.DefMsgsList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) DefMsgList(c *gin.Context) {
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

	response, err := h.serviceManager.SystemService().DefMsgList(
		ctx, &pb.ListReq{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list defmsgs", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DefMsgUpdate ...
// @Router /defmsg/update/{id} [put]
// @Summary DefMsgUpdate
// @Description This API for updating defmsg
// @Tags defmsg
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param DefMsg request body models.DefMsgUpdate true "defmsgUpdateRequest"
// @Success 200 {object} models.DefMsgResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) DefMsgUpdate(c *gin.Context) {
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
		body        pb.DefMsg
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

	response, err := h.serviceManager.SystemService().DefMsgUpdate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update defmsg", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DefMsgDelete ...
// @Router /defmsg/delete/{id} [delete]
// @Summary DefMsgDelete
// @Description This API for deleting defmsg
// @Tags defmsg
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Empty
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) DefMsgDelete(c *gin.Context) {
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

	response, err := h.serviceManager.SystemService().DefMsgDelete(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete defmsg", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
