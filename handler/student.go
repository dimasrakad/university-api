package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"university-api/student"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type studentHandler struct {
	studentService student.Service
}

func NewStudentHandler(service student.Service) *studentHandler {
	return &studentHandler{service}
}

func (h *studentHandler) GetAllStudentsHandler(c *gin.Context) {

	students, err := h.studentService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var studentsResponse []student.StudentResponse

	for _, s := range students {
		studentResponse := student.ConvertToStudentResponse(s)
		studentsResponse = append(studentsResponse, studentResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": studentsResponse,
	})
}

func (h *studentHandler) CreateStudentHandler(c *gin.Context) {
	var studentRequest student.StudentRequest

	err := c.ShouldBindJSON(&studentRequest)

	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}
	}

	s, err := h.studentService.Create(studentRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	studentResponse := student.ConvertToStudentResponse(s)

	c.JSON(http.StatusOK, gin.H{
		"data": studentResponse,
	})
}

func (h *studentHandler) GetStudentHandler(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	b, err := h.studentService.FindByID(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	studentResponse := student.ConvertToStudentResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": studentResponse,
	})
}

func (h *studentHandler) UpdateStudentHandler(c *gin.Context) {
	var studentRequest student.StudentRequest

	err := c.ShouldBindJSON(&studentRequest)

	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}
	}

	ID, _ := strconv.Atoi(c.Param("id"))
	b, err := h.studentService.Update(ID, studentRequest)
	studentResponse := student.ConvertToStudentResponse(b)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": studentResponse,
	})
}

func (h *studentHandler) DeleteStudentHandler(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	b, err := h.studentService.Delete(ID)
	studentResponse := student.ConvertToStudentResponse(b)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": studentResponse,
	})
}

func (h *studentHandler) AddCourseToStudentHandler(c *gin.Context) {
	stringStudentID := c.Query("studentID")
	stringCourseID := c.Query("courseID")

	intStudentID, _ := strconv.Atoi(stringStudentID)
	intCourseID, _ := strconv.Atoi(stringCourseID)

	b, err := h.studentService.AddCourseToStudent(intStudentID, intCourseID)
	studentResponse := student.ConvertToStudentResponse(b)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": studentResponse,
	})
}
