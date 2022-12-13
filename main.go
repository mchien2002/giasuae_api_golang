package main

import (
	"giasuaeapi/src/config"
	"giasuaeapi/src/controllers"
	"giasuaeapi/src/middleware"
	"giasuaeapi/src/repositories"
	"giasuaeapi/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	accountReponsitory   repositories.AccountReponsitory   = repositories.NewAccountReponsitory(db)
	sugbjectRepository   repositories.SubjectRepository    = repositories.NewSubjectRepository(db)
	newClassRepository   repositories.NewClassRepository   = repositories.NewNewClassRepository(db)
	classRepository      repositories.ClassRepository      = repositories.NewClassITRepository(db)
	categoryRepository   repositories.CategoryRepository   = repositories.NewCategoryRepository(db)
	postRepository       repositories.PostRepository       = repositories.NewPostRepository(db)
	transRepository      repositories.TransRepository      = repositories.NewTransRepository(db)
	salaryinfoRepository repositories.SalaryinfoRepository = repositories.NewSalaryinfoRepository(db)
	tutorRepository      repositories.TutorRepository      = repositories.NewTutorRepository(db)

	jwtService        services.JWTService        = services.NewJWTService()
	accountService    services.AccountService    = services.NewAccountService(accountReponsitory)
	subjectService    services.SubjectService    = services.NewSubjectService(sugbjectRepository)
	authService       services.AuthService       = services.NewAuthService(accountReponsitory)
	newClassService   services.NewClassService   = services.NewNewClassService(newClassRepository)
	classService      services.ClassService      = services.NewClassITService(classRepository)
	categoryService   services.CategoryService   = services.NewCategoryService(categoryRepository)
	postService       services.PostService       = services.NewPostService(postRepository)
	transService      services.TransService      = services.NewTransService(transRepository)
	salaryinfoService services.SalaryinfoService = services.NewSalaryinfoService(salaryinfoRepository)
	tutorService      services.TutorService      = services.NewTutorService(tutorRepository)

	accountController    controllers.AccountController    = controllers.NewAccountController(accountService)
	authCtrl             controllers.AuthController       = controllers.NewAuthController(authService, jwtService)
	subjectController    controllers.SubjectController    = controllers.NewSubjectController(subjectService)
	newClassController   controllers.NewClassController   = controllers.NewNewClassController(newClassService)
	classController      controllers.ClassController      = controllers.NewClassITController(classService)
	categoryController   controllers.CategoryController   = controllers.NewCategoryController(categoryService)
	postController       controllers.PostController       = controllers.NewPostController(postService)
	transController      controllers.TransController      = controllers.NewTransController(transService)
	salaryinfoController controllers.SalaryinfoController = controllers.NewSalaryinfoController(salaryinfoService)
	tutorController      controllers.TutorController      = controllers.NewTutorController(tutorService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "WELCOME TO GIASUANHEM ><"})
	})

	authRoutes := r.Group("v1/auth")
	{
		authRoutes.POST("/login", authCtrl.Login)
		authRoutes.POST("/register", authCtrl.Register)
	}

	subjectRoutes := r.Group("v1/subject")
	{
		subjectRoutes.GET("/index", subjectController.FindAllSubject)
		subjectRoutes.POST("/index", middleware.AuthorJWT(jwtService), subjectController.InsertSubject)
		subjectRoutes.GET("/id", subjectController.FindByID)
		subjectRoutes.POST("/edit", middleware.AuthorJWT(jwtService), subjectController.UpdateSubject)
		subjectRoutes.POST("/remove", subjectController.DeleteSubject)
	}

	classRoutes := r.Group("v1/class")
	{
		classRoutes.GET("/index", classController.FindAllClass)
		classRoutes.POST("/index", classController.InsertClass)
		classRoutes.POST("/remove", classController.DeleteClass)
		classRoutes.POST("/edit", classController.UpdateClass)
		classRoutes.GET("/id", classController.FindByID)
	}

	categoryRoutes := r.Group("v1/category")
	{
		categoryRoutes.GET("/index", categoryController.FindAllCategory)
		categoryRoutes.GET("/id", categoryController.FindByID)
		categoryRoutes.POST("/index", categoryController.InsertCategory)
		categoryRoutes.POST("/edit", categoryController.UpdateCategory)
		categoryRoutes.POST("/remove", categoryController.DeleteCategory)
		categoryRoutes.GET("/filter", categoryController.FilterCategorry)
	}
	accountRoutes := r.Group("v1/account")
	{
		accountRoutes.GET("/index", accountController.FindAllAccount)
		accountRoutes.GET("/id", accountController.FindByID)
		accountRoutes.POST("/remove", accountController.DeleteAccount)
		accountRoutes.POST("/edit", accountController.UpdateAccount)
		accountRoutes.GET("/filter", accountController.FilterAccount)
	}

	newClassRoutes := r.Group("v1/new_class")
	{
		newClassRoutes.GET("/index", newClassController.FindAllNewClass)
		newClassRoutes.POST("/index", newClassController.InsertNewClass)
		newClassRoutes.POST("/edit", newClassController.UpdateNewClass)
		newClassRoutes.GET("/id", newClassController.FindByID)
		newClassRoutes.POST("/remove", newClassController.DeleteNewClass)
	}

	postRoutes := r.Group("v1/post")
	{
		postRoutes.GET("/index", postController.FindAllPost)
		postRoutes.POST("/index", postController.InsertPost)
		postRoutes.POST("/edit", postController.UpdatePost)
		postRoutes.POST("/remove", postController.DeletePost)
		postRoutes.GET("/id", postController.FindByID)
		postRoutes.GET("/filter", postController.FilterPost)
	}
	transRoutes := r.Group("v1/trans")
	{
		transRoutes.GET("/index", transController.FindAllTrans)
		transRoutes.POST("/index", transController.InsertTrans)
		transRoutes.POST("/id", transController.FindByIDAcc)
	}

	salRoutes := r.Group("v1/salaryinfo")
	{
		salRoutes.GET("/index", salaryinfoController.FindAllSalaryinfo)
		salRoutes.POST("/index", salaryinfoController.InsertSalaryinfo)
		salRoutes.POST("/remove", salaryinfoController.DeleteSalaryinfo)
		salRoutes.POST("/edit", salaryinfoController.UpdateSalaryinfo)
		salRoutes.GET("/id", salaryinfoController.FindByID)
		salRoutes.GET("/filter", salaryinfoController.FindByType)
	}

	tutorRoutes := r.Group("v1/tutor")
	{
		tutorRoutes.GET("/index", tutorController.FindAllTutor)
		tutorRoutes.POST("/index", tutorController.InsertTutor)
		tutorRoutes.GET("/id", tutorController.FindByID)
		tutorRoutes.POST("/remove", tutorController.DeleteTutor)
		tutorRoutes.POST("/edit", tutorController.UpdateTutor)
	}

	r.Run()
}
