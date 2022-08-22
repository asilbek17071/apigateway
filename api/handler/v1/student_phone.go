package v1

import (
	"context"
	"net/http"
	"time"

	pb "github.com/asilbek17071/apigateway/genproto/student_service"
	l "github.com/asilbek17071/apigateway/pkg/logger"
	"github.com/asilbek17071/apigateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// StudentPhoneCreate ...
// @Summary StudentPhoneCreate
// @Router /student/phonecreate/ [post]
// @Description This API for creating a new student
// @Tags student
// @Accept  json
// @Produce  json
// @Param Student request body models.SPhon true "studentCreateRequest"
// @Success 200 {object} models.SPhone
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) StudentPhoneCreate(c *gin.Context) {
	var (
		body        pb.SPhone
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

	response, err := h.serviceManager.StudentService().StudentPhoneCreate(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create student_phone", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// StudentPhoneGet ...
// @Router /student/listphone/{id} [get]
// @Summary StudentPhoneGet
// @Description This API for getting student StudentPhoneList
// @Tags student
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.SPhone
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) StudentPhoneGet(c *gin.Context) {
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

	response, err := h.serviceManager.StudentService().StudentPhoneGet(
		ctx, &pb.ByIdReq{
			Id: params.Id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get student_phone", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
