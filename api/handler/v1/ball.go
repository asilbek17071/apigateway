package v1

import (
	"context"
	"net/http"
	"time"

	pb "github.com/asilbek17071/apigateway/genproto/course_service"
	l "github.com/asilbek17071/apigateway/pkg/logger"
	"github.com/asilbek17071/apigateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// BallCreate ...
// @Summary BallCreate
// @Router /ball/create/ [post]
// @Description This API for creating a new ball
// @Tags ball
// @Accept  json
// @Produce  json
// @Param Ball request body models.Ball true "ballCreateRequest"
// @Success 200 {object} models.BallResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) BallCreate(c *gin.Context) {
	var (
		body        pb.Ball
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

	response, err := h.serviceManager.CourseService().BallCreate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create ball", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// BallGet ...
// @Router /ball/byid/{id} [get]
// @Summary BallGet
// @Description This API for getting ball BallList
// @Tags ball
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.BallResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) BallGet(c *gin.Context) {
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

	response, err := h.serviceManager.CourseService().BallGet(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get ball", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// BallList ...
// @Router /ball/list/ [get]
// @Summary BallList
// @Description This API for getting list of balls
// @Tags ball
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} models.BallsList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) BallList(c *gin.Context) {
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

	response, err := h.serviceManager.CourseService().BallList(
		ctx, &pb.ListReq{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list balls", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// BallUpdate ...
// @Router /ball/update/{id} [put]
// @Summary BallUpdate
// @Description This API for updating ball
// @Tags ball
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param Ball request body models.BallUpdate true "ballUpdateRequest"
// @Success 200 {object} models.BallResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) BallUpdate(c *gin.Context) {
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
		body        pb.Ball
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

	response, err := h.serviceManager.CourseService().BallUpdate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update ball", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// BallDelete ...
// @Router /ball/delete/{id} [delete]
// @Summary BallDelete
// @Description This API for deleting ball
// @Tags ball
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Empty
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) BallDelete(c *gin.Context) {
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

	response, err := h.serviceManager.CourseService().BallDelete(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete ball", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
