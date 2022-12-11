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

type AccountController interface {
	InsertAccount(context *gin.Context)
	UpdateAccount(context *gin.Context)
	DeleteAccount(context *gin.Context)
	FindAllAccount(context *gin.Context)
	FindByID(context *gin.Context)
}

type accountController struct {
	AccountService services.AccountService
}

// DeleteAccount implements AccountController
func (ctrl *accountController) DeleteAccount(context *gin.Context) {
	id, err := strconv.ParseInt(context.Query("id"), 0, 0)
	if err != nil {
		res := helper.BuildResponseError("Không có id được tìm thấy", err.Error(), helper.EmptyObjec{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var acc = entities.Account{
		ID: int(id),
	}
	ctrl.AccountService.DeleteAccount(&acc)
}

// FindAllAccount implements AccountController
func (ctrl *accountController) FindAllAccount(context *gin.Context) {
	var accs []entities.Account = ctrl.AccountService.FindAllAccount()
	res := helper.BuildResponse(true, "OK", accs)
	context.JSON(http.StatusOK, res)
}

// FindByID implements AccountController
func (ctrl *accountController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Query("id"), 0, 0)
	if err != nil {
		res := helper.BuildResponseError("Không có id được tìm thấy", err.Error(), helper.EmptyObjec{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var acc entities.Account = ctrl.AccountService.FindByID(int(id))
	res := helper.BuildResponse(true, "OK", acc)
	context.JSON(http.StatusOK, res)
}

// InsertAccount implements AccountController
func (ctrl *accountController) InsertAccount(context *gin.Context) {
	var acc entities.Account = entities.Account{
		Created_at: date_picker.FormatDataNow(),
	}

	err := context.ShouldBind(&acc)
	if err != nil {
		res := helper.BuildResponseError("Thêm tài khoản thất bại", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	ctrl.AccountService.InsertAccount(&acc)
	res := helper.BuildResponse(true, "OK", nil)
	context.JSON(http.StatusOK, res)
}

// UpdateAccount implements AccountController
func (ctrl *accountController) UpdateAccount(context *gin.Context) {
	var acc entities.Account
	err := context.ShouldBind(&acc)
	if err != nil {
		res := helper.BuildResponseError("Cập nhật tài khoản thất bại", err.Error(), helper.EmptyObjec{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	ctrl.AccountService.UpdateAccount(&acc)
	res := helper.BuildResponse(true, "OK", nil)
	context.JSON(http.StatusOK, res)
}

func NewAccountController(accsv services.AccountService) AccountController {
	return &accountController{
		AccountService: accsv,
	}
}
