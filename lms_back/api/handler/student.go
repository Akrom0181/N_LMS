package handler

import (
	"context"
	"fmt"
	"lms_back/api/models"
	"lms_back/config"
	"lms_back/pkg/password"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateStudent godoc
// @Security ApiKeyAuth
// @Router 		/student/login [POST]
// @Summary     login
// @Description This api is to login
// @Tags 		auth
// @Accept		json
// @Produce		json
// @Param		student  body  models.CreateStudent true "student"
// @Success		200  {object}  models.Student
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateStudent(c *gin.Context) {
	student := models.Student{}

	if err := c.ShouldBindJSON(&student); err != nil {

		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	hashedPass, err := password.HashPassword(student.Password)
	if err != nil {
		handleResponseLog(c, h.Log, "error while generating customer password", http.StatusInternalServerError, err.Error())
		return
	}
	student.Password = string(hashedPass)

	id, err := h.Service.Student().Create(ctx, student)
	if err != nil {
		handleResponseLog(c, h.Log, "error while creating student", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "created successfully", http.StatusOK, id)
}

// UpdateStudent godoc
// @Security ApiKeyAuth
// @Router                /student/{id} [PUT]
// @Summary 			  update a student
// @Description:          this api updates student information
// @Tags 			      student
// @Accept 			      json
// @Produce 		      json
// @Param 			      id path string true "Student ID"
// @Param       		  car body models.UpdateStudent true "student"
// @Success 		      200 {object} models.Student
// @Failure 		      400 {object} models.Response
// @Failure               404 {object} models.Response
// @Failure 		      500 {object} models.Response
func (h Handler) UpdateStudent(c *gin.Context) {

	student := models.Student{}
	if err := c.ShouldBindJSON(&student); err != nil {
		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	student.ID = c.Query("id")

	err := uuid.Validate(student.ID)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()


	hashedPass, err := password.HashPassword(student.Password)
	if err != nil {
		handleResponseLog(c, h.Log, "error while generating customer password", http.StatusInternalServerError, err.Error())
		return
	}
	
	student.Password = string(hashedPass)
	id, err := h.Service.Student().Update(ctx, student)
	if err != nil {
		handleResponseLog(c, h.Log, "error while updating student", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "updated successfully", http.StatusOK, id)
}

// GetAllStudents godoc
// @Router 			/student [GET]
// @Summary 		get all students
// @Description 	This API returns student list
// @Tags 			student
// Accept			json
// @Produce 		json
// @Param 			page query int false "page number"
// @Param 			limit query int false "limit per page"
// @Param 			search query string false "search keyword"
// @Success 		200 {object} models.GetAllStudentsResponse
// @Failure 		400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure 		500 {object} models.Response
func (h Handler) GetAllStudent(c *gin.Context) {
	var (
		request = models.GetAllStudentsRequest{}
	)

	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing page", http.StatusInternalServerError, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponseLog(c, h.Log, "error while parsing limit", http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	student, err := h.Service.Student().GetAll(ctx, request)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting student", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, student)
}

// GetByIDStudent godoc
// @Router       /student/{id} [GET]
// @Summary      return a student by ID
// @Description  Retrieves a student by its ID
// @Tags         student
// @Accept       json
// @Produce      json
// @Param        id path string true "Student ID"
// @Success      200 {object} models.GetStudent
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDStudent(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	student, err := h.Service.Student().GetByID(ctx, id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting student by id", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, student)
}

// DeleteStudent godoc
// @Router          /student/{id} [DELETE]
// @Summary         delete a student by ID
// @Description     Deletes a student by its ID
// @Tags            student
// @Accept          json
// @Produce         json
// @Param           id path string true "Student ID"
// @Success         200 {string} models.Response
// @Failure         400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure         500 {object} models.Response
func (h Handler) DeleteStudent(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	err = h.Service.Student().Delete(ctx, id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while deleting student", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "deleted student", http.StatusOK, id)
}
