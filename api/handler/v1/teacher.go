package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "github.com/asilbek17071/apigateway/genproto/teacher_service"
	l "github.com/asilbek17071/apigateway/pkg/logger"
	"github.com/asilbek17071/apigateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// TeacherCreate ...
// @Summary TeacherCreate
// @Router /teacher/create/ [post]
// @Description This API for creating a new teacher
// @Tags teacher
// @Accept  json
// @Produce  json
// @Param Teacher request body models.Teacher true "teacherCreateRequest"
// @Success 200 {object} models.TeacherResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) TeacherCreate(c *gin.Context) {
	var (
		body        pb.Teacher
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

	response, err := h.serviceManager.TeacherService().TeacherCreate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create teacher", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// TeacherGet ...
// @Router /teacher/byid/{id} [get]
// @Summary TeacherGet
// @Description This API for getting teacher TeacherList
// @Tags teacher
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.TeacherResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) TeacherGet(c *gin.Context) {
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

	response, err := h.serviceManager.TeacherService().TeacherGet(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get teacher", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// TeacherGet ...
// @Router /teacher/byid/{id} [get]
// @Summary TeacherGet
// @Description This API for getting teacher TeacherList
// @Tags teacher
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.TeacherResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) TeacherByName(c *gin.Context) {
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

	response, err := h.serviceManager.TeacherService().TeacherByName(
		ctx, &pb.TTT{
			Permission: params.Permission,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get teacher", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// TeacherGet ...
// @Router /signup/teacher/{token} [get]
// @Summary TeacherGet
// @Description This API for getting teacher TeacherList
// @Tags teacher
// @Accept  json
// @Produce  json
// @Param token path string true "TOEKN"
// @Success 200 {object} models.TeacherResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) TeacherSignUp(c *gin.Context) {
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

	response, err := h.serviceManager.TeacherService().TeacherSignUp(
		ctx, &pb.Token{
			Token: params.Token,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get teacher", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// TeacherList ...
// @Router /teacher/list/ [get]
// @Summary TeacherList
// @Description This API for getting list of teachers
// @Tags teacher
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} models.TeachersList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) TeacherList(c *gin.Context) {
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

	response, err := h.serviceManager.TeacherService().TeacherList(
		ctx, &pb.ListReq{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list teachers", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// TeacherLogin ...
// @Router /teacher/login/ [post]
// @Summary TeacherLogin
// @Description This API for getting list of teachers
// @Tags teacher
// @Accept  json
// @Produce  json
// @Param Teacher request body models.Login true "teacherCreateRequest"
// @Success 200 {object} models.TeachersList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) TeacherLogin(c *gin.Context) {
	var (
		body        pb.LoginReq
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	fmt.Println(body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.TeacherService().TeacherLogin(
		ctx, &pb.LoginReq{
			Email:    body.Email,
			Password: body.Password,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list teachers", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// TeacherUpdate ...
// @Router /teacher/update/{id} [put]
// @Summary TeacherUpdate
// @Description This API for updating teacher
// @Tags teacher
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param Teacher request body models.TeacherUpdate true "teacherUpdateRequest"
// @Success 200 {object} models.TeacherResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) TeacherUpdate(c *gin.Context) {
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
		body        pb.Teacher
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

	response, err := h.serviceManager.TeacherService().TeacherUpdate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update teacher", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// TeacherDelete ...
// @Router /teacher/delete/{id} [delete]
// @Summary TeacherDelete
// @Description This API for deleting teacher
// @Tags teacher
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Empty
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) TeacherDelete(c *gin.Context) {
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

	response, err := h.serviceManager.TeacherService().TeacherDelete(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete teacher", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
