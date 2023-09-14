package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"university-api/course"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type courseHandler struct {
	courseService course.Service
}

func NewCourseHandler(service course.Service) *courseHandler {
	return &courseHandler{service}
}

func (h *courseHandler) GetAllCoursesHandler(c *gin.Context) {
	courses, err := h.courseService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var coursesResponse []course.CourseResponse

	for _, s := range courses {
		courseResponse := course.ConvertToCourseResponse(s)
		coursesResponse = append(coursesResponse, courseResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": coursesResponse,
	})
}

func (h *courseHandler) CreateCourseHandler(c *gin.Context) {
	var courseRequest course.CourseRequest

	err := c.ShouldBindJSON(&courseRequest)

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

	course, err := h.courseService.Create(courseRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": course,
	})
}

func (h *courseHandler) GetCourseHandler(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	b, err := h.courseService.FindByID(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	courseResponse := course.ConvertToCourseResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": courseResponse,
	})
}

func (h *courseHandler) UpdateCourseHandler(c *gin.Context) {
	var courseRequest course.CourseRequest

	err := c.ShouldBindJSON(&courseRequest)

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
	b, err := h.courseService.Update(ID, courseRequest)
	courseResponse := course.ConvertToCourseResponse(b)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": courseResponse,
	})
}

func (h *courseHandler) DeleteCourseHandler(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	b, err := h.courseService.Delete(ID)
	courseResponse := course.ConvertToCourseResponse(b)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": courseResponse,
	})
}
