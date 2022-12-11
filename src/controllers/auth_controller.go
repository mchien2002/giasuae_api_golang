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

type AuthController interface {
	Login(cx *gin.Context)
	Register(cx *gin.Context)
}

type authController struct {
	AuthService services.AuthService
	JWTService  services.JWTService
}

// Login implements AuthController
func (ctr *authController) Login(cx *gin.Context) {
	username := cx.Query("username")
	password := cx.Query("password")
	value, err := ctr.AuthService.VerifyCredential(username, password)

	if err != nil {
		res := helper.BuildResponseError(err.Error(), err.Error(), helper.EmptyObjec{})
		cx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else if acc, ok := value.(entities.AccountWithToken); ok {
		acc.Token = ctr.JWTService.GenerateToken(strconv.Itoa(acc.ID))
		res := helper.BuildResponse(true, "Đăng nhập thành công", acc)
		cx.JSON(http.StatusOK, res)
		return
	}
}

// Register implements AuthController
func (ctrl *authController) Register(cx *gin.Context) {
	var acc entities.Account = entities.Account{
		Created_at: date_picker.FormatDataNow(),
	}
	err := cx.ShouldBind(&acc)
	if err != nil {
		res := helper.BuildResponseError("Đăng ký thất bại", err.Error(), helper.EmptyObjec{})
		cx.JSON(http.StatusBadRequest, res)
		return
	}
	ctrl.AuthService.CreateUser(&acc)
	res := helper.BuildResponse(true, "OK", nil)
	cx.JSON(http.StatusOK, res)
}

func NewAuthController(authSvc services.AuthService, jwtSvc services.JWTService) AuthController {
	return &authController{
		AuthService: authSvc,
		JWTService:  jwtSvc,
	}
}
