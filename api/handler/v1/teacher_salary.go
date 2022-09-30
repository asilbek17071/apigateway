package v1

import (
	"context"
	"net/http"
	"time"

	pb "github.com/asilbek17071/apigateway/genproto/teacher_service"
	l "github.com/asilbek17071/apigateway/pkg/logger"
	"github.com/asilbek17071/apigateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// TeacherSalaryCreate ...
// @Summary TeacherSalaryCreate
// @Router /salary/teacher/create/ [post]
// @Description This API for creating a new teacher_salary
// @Tags teacher_salary
// @Accept  json
// @Produce  json
// @Param TeacherSalary request body models.TeacherSalary true "teacher_salaryCreateRequest"
// @Success 200 {object} models.TeacherSalaryResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) TeacherSalaryCreate(c *gin.Context) {
	var (
		body        pb.TeacherSalary
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

	response, err := h.serviceManager.TeacherService().TeacherSalaryCreate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create direction", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// TeacherSalaryGet ...
// @Router /salary/teacher/byid/{id} [get]
// @Summary TeacherSalaryGet
// @Description This API for getting teacher_salary TeacherSalaryList
// @Tags teacher_salary
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.TeacherSalaryResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) TeacherSalaryGet(c *gin.Context) {
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

	response, err := h.serviceManager.TeacherService().TeacherSalaryGet(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get teacher_salary", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// TeacherSalaryList ...
// @Router /salary/teacher/list/ [get]
// @Summary TeacherSalaryList
// @Description This API for getting list of teacher_salarys
// @Tags teacher_salary
// @Accept  json
// @Produce  json
// @Param id query string false "Bo'tga teacher yoki assistantni id si"
// @Success 200 {object} models.TeachersSalaryList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) TeacherSalaryList(c *gin.Context) {
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

	response, err := h.serviceManager.TeacherService().TeacherSalaryList(
		ctx, &pb.SalaryReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list teacher_salarys", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// TeacherSalaryUpdate ...
// @Router /salary/teacher/update/{id} [put]
// @Summary TeacherSalaryUpdate
// @Description This API for updating teacher_salary
// @Tags teacher_salary
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param TeacherSalary request body models.TeacherSalaryUpdate true "teacher_salaryUpdateRequest"
// @Success 200 {object} models.TeacherSalaryResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) TeacherSalaryUpdate(c *gin.Context) {
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
		body        pb.TeacherSalary
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

	response, err := h.serviceManager.TeacherService().TeacherSalaryUpdate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update teacher_salary", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// TeacherSalaryDelete ...
// @Router /salary/teacher/delete/{id} [delete]
// @Summary TeacherSalaryDelete
// @Description This API for deleting teacher_salary
// @Tags teacher_salary
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Empty
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) TeacherSalaryDelete(c *gin.Context) {
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

	response, err := h.serviceManager.TeacherService().TeacherSalaryDelete(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete teacher_salary", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
