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

// AttendanceCreate ...
// @Summary AttendanceCreate
// @Router /attendance/create/ [post]
// @Description This API for creating a new attendance
// @Tags attendance
// @Accept  json
// @Produce  json
// @Param Attendance request body models.Attendance true "attendanceCreateRequest"
// @Success 200 {object} models.AttendanceResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) AttendanceCreate(c *gin.Context) {
	var (
		body        pb.Attendance
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

	response, err := h.serviceManager.CourseService().AttendanceCreate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create attendance", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// AttendanceGet ...
// @Router /attendance/byid/{id} [get]
// @Summary AttendanceGet
// @Description This API for getting attendance AttendanceList
// @Tags attendance
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.AttendanceResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) AttendanceGet(c *gin.Context) {
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

	response, err := h.serviceManager.CourseService().AttendanceGet(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get attendance", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// AttendanceList ...
// @Router /attendance/list/ [get]
// @Summary AttendanceList
// @Description This API for getting list of attendances
// @Tags attendance
// @Accept  json
// @Produce  json
// @Param id query string false "ID"
// @Success 200 {object} models.AttendancesList
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) AttendanceList(c *gin.Context) {
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

	response, err := h.serviceManager.CourseService().AttendanceList(
		ctx, &pb.GetParticipate{
			GroupId: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list attendances", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ParticipateResp ...
// @Router /attendance/list/participate/ [get]
// @Summary ParticipateResp
// @Description Bugungi kundagi yo'qlama qilinib bo'linganlarini ko'rish, birinchi count nechta kelganligi, ikkinchisi esa asli nechta kelishi kereligi
// @Tags attendance
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Participate
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) ParticipateResp(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CourseService().ParticipateResp(
		ctx, &pb.EmptyRes{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list attendances", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// AttendanceUpdate ...
// @Router /attendance/update/{id} [put]
// @Summary AttendanceUpdate
// @Description This API for updating attendance
// @Tags attendance
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param Attendance request body models.AttendanceUpdate true "attendanceUpdateRequest"
// @Success 200 {object} models.AttendanceResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) AttendanceUpdate(c *gin.Context) {
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
		body        pb.Attendance
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

	response, err := h.serviceManager.CourseService().AttendanceUpdate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update attendance", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// AttendanceDelete ...
// @Router /attendance/delete/{id} [delete]
// @Summary AttendanceDelete
// @Description This API for deleting attendance
// @Tags attendance
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Empty
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) AttendanceDelete(c *gin.Context) {
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

	response, err := h.serviceManager.CourseService().AttendanceDelete(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete attendance", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
