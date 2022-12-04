package controllers

import (
	datepicker "giasuaeapi/src/date_picker"
	"giasuaeapi/src/dto"
	"giasuaeapi/src/helper"
	"giasuaeapi/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewClassController interface {
	InsertNewClass(context *gin.Context)
	UpdateNewClass(context *gin.Context)
	DeleteNewClass(context *gin.Context)
	FindAllNewClass(context *gin.Context)
	FindByID(context *gin.Context)
}

type newClassController struct {
	NewClassService services.NewClassService
}

// DeleteNewClass implements NewClassController
func (*newClassController) DeleteNewClass(context *gin.Context) {
	panic("unimplemented")
}

// FindAllNewClass implements NewClassController
func (*newClassController) FindAllNewClass(context *gin.Context) {
	panic("unimplemented")
}

// FindByID implements NewClassController
func (*newClassController) FindByID(context *gin.Context) {
	panic("unimplemented")
}

// InsertNewClass implements NewClassController
func (ctrl *newClassController) InsertNewClass(context *gin.Context) {
	var nc dto.Newclasses = dto.Newclasses{
		Created_at: datepicker.FormatDataNow(),
	}
	err := context.ShouldBind(&nc)
	if err != nil {
		res := helper.BuildResponseError("Thêm lớp họp mới thất bại", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	err2 := ctrl.NewClassService.InsertNewClass(&nc)
	if err2 != nil {
		res := helper.BuildResponseError("Thêm lớp họp mới thất bại", err2.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK", nil)
	context.JSON(http.StatusOK, res)
}

// UpdateNewClass implements NewClassController
func (*newClassController) UpdateNewClass(context *gin.Context) {
	panic("unimplemented")
}

func NewNewClassController(ncSvc services.NewClassService) NewClassController {
	return &newClassController{
		NewClassService: ncSvc,
	}
}
