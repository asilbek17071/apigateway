package v1

import (
	"context"
	"net/http"
	"time"

	pb "github.com/asilbek17071/apigateway/genproto/user_service"
	l "github.com/asilbek17071/apigateway/pkg/logger"
	"github.com/asilbek17071/apigateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// PersonalSalaryCreate ...
// @Summary PersonalSalaryCreate
// @Router /salary/personal/create [post]
// @Description This API for creating a new personalSalary
// @Tags personalSalary
// @Accept  json
// @Produce  json
// @Param PersonalSalary request body models.PersonalSalary true "personalSalaryCreateRequest"
// @Success 200 {object} models.PersonalSalaryResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) PersonalSalaryCreate(c *gin.Context) {
	var (
		body        pb.PersonalSalary
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

	response, err := h.serviceManager.UserService().PersonalSalaryCreate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create personal_salary", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// PersonalSalaryGet ...
// @Router /salary/personal/byid/{id} [get]
// @Summary PersonalSalaryGet
// @Description This API for getting personalSalary PersonalSalaryList
// @Tags personalSalary
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.PersonalSalaryResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) PersonalSalaryGet(c *gin.Context) {
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

	response, err := h.serviceManager.UserService().PersonalSalaryGet(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get personal_salary", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// PersonalSalaryList ...
// @Router /salary/personal/list [get]
// @Summary PersonalSalaryList
// @Description This API for getting list of personalsSalary
// @Tags personalSalary
// @Accept  json
// @Produce  json
// @Param id query string false "ID"
// @Success 200 {object} models.PersonalsSalaryList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) PersonalSalaryList(c *gin.Context) {
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

	response, err := h.serviceManager.UserService().PersonalSalaryList(
		ctx, &pb.SalaryReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list personal_salarys", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// PersonalSalaryUpdate ...
// @Router /salary/personal/update/{id} [put]
// @Summary PersonalSalaryUpdate
// @Description This API for updating personalSalary
// @Tags personalSalary
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param PersonalSalary request body models.PersonalSalaryUpdate true "personalSalaryUpdateRequest"
// @Success 200 {object} models.PersonalSalaryResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) PersonalSalaryUpdate(c *gin.Context) {
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
		body        pb.PersonalSalary
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

	response, err := h.serviceManager.UserService().PersonalSalaryUpdate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update personal_salary", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// PersonalSalaryDelete ...
// @Router /salary/personal/delete/:id [delete]
// @Summary PersonalSalaryDelete
// @Description This API for deleting personalSalary
// @Tags personalSalary
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Empty
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) PersonalSalaryDelete(c *gin.Context) {
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

	response, err := h.serviceManager.UserService().PersonalSalaryDelete(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete personal_salary", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
