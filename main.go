package main

import (
	"log"
	"university-api/course"
	"university-api/handler"
	"university-api/initializers"
	"university-api/middleware"
	"university-api/student"
	"university-api/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	initializers.LoadEnvVariables()
	db, err = initializers.ConnectToDatabase()
	err = initializers.SyncDatabase(db)

	if err != nil {
		log.Fatal("Connection to database failed")
	}
}

func main() {
	studentRepository := student.NewRepository(db)
	studentService := student.NewService(studentRepository)
	studentHandler := handler.NewStudentHandler(studentService)

	courseRepository := course.NewRepository(db)
	courseService := course.NewService(courseRepository)
	courseHandler := handler.NewCourseHandler(courseService)

	//inisiasi userhandler
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	routerV1 := router.Group("/v1")

	/*
		USER ROUTER
	*/
	routerV1.POST("/signup", userHandler.Signup)
	routerV1.POST("/login", userHandler.Login)

	/*
		STUDENT ROUTER
	*/
	routerV1Students := routerV1.Group("/students", middleware.RequireAuth)
	routerV1Students.GET("", studentHandler.GetAllStudentsHandler)
	routerV1Students.POST("", studentHandler.CreateStudentHandler)
	routerV1Students.GET("/:id", studentHandler.GetStudentHandler)
	routerV1Students.PUT("/:id", studentHandler.UpdateStudentHandler)
	routerV1Students.DELETE("/:id", studentHandler.DeleteStudentHandler)
	routerV1Students.GET("/add-course", studentHandler.AddCourseToStudentHandler)

	/*
		COURSE ROUTER
	*/
	routerV1Courses := routerV1.Group("/courses", middleware.RequireAuth)
	routerV1Courses.GET("", courseHandler.GetAllCoursesHandler)
	routerV1Courses.POST("", courseHandler.CreateCourseHandler)
	routerV1Courses.GET("/:id", courseHandler.GetCourseHandler)
	routerV1Courses.PUT("/:id", courseHandler.UpdateCourseHandler)
	routerV1Courses.DELETE("/:id", courseHandler.DeleteCourseHandler)

	router.Run(":3030")
}
