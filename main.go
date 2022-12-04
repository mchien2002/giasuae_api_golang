package main

import (
	"giasuaeapi/src/config"
	"giasuaeapi/src/controllers"
	"giasuaeapi/src/middleware"
	"giasuaeapi/src/repositories"
	"giasuaeapi/src/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                 *gorm.DB                        = config.SetupDatabaseConnection()
	jwtService         services.JWTService             = services.NewJWTService()
	accountReponsitory repositories.AccountReponsitory = repositories.NewAccountReponsitory(db)
	accountService     services.AccountService         = services.NewAccountService(accountReponsitory)
	accountController  controllers.AccountController   = controllers.NewAccountController(accountService)
	authService        services.AuthService            = services.NewAuthService(accountReponsitory)
	authCtrl           controllers.AuthController      = controllers.NewAuthController(authService, jwtService)
	sugbjectRepository repositories.SubjectRepository  = repositories.NewSubjectRepository(db)
	subjectService     services.SubjectService         = services.NewSubjectService(sugbjectRepository)
	subjectController  controllers.SubjectController   = controllers.NewSubjectController(subjectService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("giasuae/v1/auth")
	{
		authRoutes.POST("/login", authCtrl.Login)
		authRoutes.POST("/register", authCtrl.Register)
	}

	subjectRoutes := r.Group("giasuae/v1/subject")
	{
		subjectRoutes.GET("/index", subjectController.FindAllSubject)
		subjectRoutes.POST("/index", middleware.AuthorJWT(jwtService), subjectController.InsertSubject)
		subjectRoutes.GET("/id", subjectController.FindByID)
		subjectRoutes.POST("/edit", middleware.AuthorJWT(jwtService), subjectController.UpdateSubject)
	}
	accountRoutes := r.Group("giasuae/v1/account")
	{
		accountRoutes.GET("/index", accountController.FindAllAccount)
		accountRoutes.GET("/filter", accountController.FindByID)
		accountRoutes.POST("/index", accountController.InsertAccount)
		accountRoutes.POST("/remove", accountController.DeleteAccount)
		accountRoutes.POST("/edit", accountController.UpdateAccount)
	}
	r.Run()
}
