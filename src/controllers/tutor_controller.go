package controllers

import (
	"giasuaeapi/src/date_picker"
	"giasuaeapi/src/entities"
	"giasuaeapi/src/helper"
	"giasuaeapi/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TutorController interface {
	InsertTutor(context *gin.Context)
	UpdateTutor(context *gin.Context)
	DeleteTutor(context *gin.Context)
	FindAllTutor(context *gin.Context)
	FindByID(context *gin.Context)
}

type tutorController struct {
	TutorService services.TutorService
}

// DeleteTutor implements TutorController
func (*tutorController) DeleteTutor(context *gin.Context) {
	panic("unimplemented")
}

// UpdateTutor implements TutorController
func (*tutorController) UpdateTutor(context *gin.Context) {
	panic("unimplemented")
}

// FindAllTutor implements TutorController
func (ctrl *tutorController) FindAllTutor(context *gin.Context) {
	var tutors []entities.TutorSet = ctrl.TutorService.FindAllTutor()
	res := helper.BuildResponse(true, "OK", tutors)
	context.JSON(http.StatusOK, res)
}

// FindByID implements TutorController
func (ctrl *tutorController) FindByID(context *gin.Context) {
	var tutor entities.TutorDetail
	id, err := strconv.ParseInt(context.Query("id"), 0, 0)
	if err != nil {
		res := helper.BuildResponseError("Không có id được tìm thấy", err.Error(), helper.EmptyObjec{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	tutor = ctrl.TutorService.FindByID(int(id))
	res := helper.BuildResponse(true, "OK", tutor)
	context.JSON(http.StatusOK, res)
}

// InsertTutor implements TutorController
func (ctrl *tutorController) InsertTutor(context *gin.Context) {
	var tutor entities.TutorReq = entities.TutorReq{
		Created_at: date_picker.FormatDataNow(),
	}

	err := context.ShouldBind(&tutor)
	if err != nil {
		res := helper.BuildResponseError("Không thể bind JSON", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	err2 := ctrl.TutorService.InsertTutor(&tutor)
	if err2 != nil {
		res := helper.BuildResponseError("Thêm gia sư thất bại", err2.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", nil)
	context.JSON(http.StatusOK, res)
}

func NewTutorController(svc services.TutorService) TutorController {
	return &tutorController{
		TutorService: svc,
	}
}
