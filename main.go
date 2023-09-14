package main

import (
	"log"
	"university-api/handler"
	"university-api/initializers"
	"university-api/student"

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

	router := gin.Default()
	routerV1 := router.Group("/v1")

	/*
		STUDENT ROUTER
	*/
	routerV1Students := routerV1.Group("/students")
	routerV1Students.GET("", studentHandler.GetAllStudentsHandler)
	routerV1Students.POST("", studentHandler.CreateStudentHandler)
	routerV1Students.GET("/:id", studentHandler.GetStudentHandler)
	routerV1Students.PUT("/:id", studentHandler.UpdateStudentHandler)
	routerV1Students.DELETE("/:id", studentHandler.DeleteStudentHandler)

	router.Run(":3030")
}
