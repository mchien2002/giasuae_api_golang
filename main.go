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
	db *gorm.DB = config.SetupDatabaseConnection()

	accountReponsitory repositories.AccountReponsitory = repositories.NewAccountReponsitory(db)
	sugbjectRepository repositories.SubjectRepository  = repositories.NewSubjectRepository(db)
	newClassRepository repositories.NewClassRepository = repositories.NewNewClassRepository(db)
	classRepository    repositories.ClassRepository    = repositories.NewClassITRepository(db)

	jwtService      services.JWTService      = services.NewJWTService()
	accountService  services.AccountService  = services.NewAccountService(accountReponsitory)
	subjectService  services.SubjectService  = services.NewSubjectService(sugbjectRepository)
	authService     services.AuthService     = services.NewAuthService(accountReponsitory)
	newClassService services.NewClassService = services.NewNewClassService(newClassRepository)
	classService    services.ClassService    = services.NewClassITService(classRepository)

	accountController  controllers.AccountController  = controllers.NewAccountController(accountService)
	authCtrl           controllers.AuthController     = controllers.NewAuthController(authService, jwtService)
	subjectController  controllers.SubjectController  = controllers.NewSubjectController(subjectService)
	newClassController controllers.NewClassController = controllers.NewNewClassController(newClassService)
	classController    controllers.ClassController    = controllers.NewClassITController(classService)
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

	classRoutes := r.Group("giasuae/v1/class")
	{
		classRoutes.GET("/index", classController.FindAllClass)
		classRoutes.POST("/index", classController.InsertClass)
		classRoutes.POST("/remove", classController.InsertClass)
		classRoutes.POST("/edit", classController.InsertClass)
	}
	accountRoutes := r.Group("giasuae/v1/account")
	{
		accountRoutes.GET("/index", accountController.FindAllAccount)
		accountRoutes.GET("/filter", accountController.FindByID)
		accountRoutes.POST("/index", accountController.InsertAccount)
		accountRoutes.POST("/remove", accountController.DeleteAccount)
		accountRoutes.POST("/edit", accountController.UpdateAccount)
	}

	newClassRoutes := r.Group("giasuae/v1/new_class")
	{
		newClassRoutes.GET("/index", newClassController.FindAllNewClass)
		newClassRoutes.POST("/index", newClassController.InsertNewClass)
		newClassRoutes.POST("/edit", newClassController.UpdateNewClass)
		newClassRoutes.GET("/id", newClassController.FindByID)
	}
	r.Run()
}
