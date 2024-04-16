package handler

import (
	"context"
	"fmt"
	_ "lms_back/api/docs"
	"lms_back/api/models"
	"lms_back/config"
	"lms_back/pkg/password"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateTeacher godoc
// @Router 		/teacher [POST]
// @Summary 	Create a teacher
// @Description This api is creates a new teacher and returns its id
// @Tags 		teacher
// @Accept		json
// @Produce		json
// @Param		car  body      models.CreateTeacher true "car"
// @Success		200  {object}  models.Teacher
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateTeacher(c *gin.Context) {
	teacher := models.Teacher{}

	if err := c.ShouldBindJSON(&teacher); err != nil {
		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	hashedPass, err := password.HashPassword(teacher.Password)
	if err != nil {
		handleResponseLog(c, h.Log, "error while generating customer password", http.StatusInternalServerError, err.Error())
		return
	}
	teacher.Password = string(hashedPass)

	id, err := h.Service.Teacher().Create(ctx, teacher)
	if err != nil {
		handleResponseLog(c, h.Log, "error while creating teacher", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "teacher created", http.StatusOK, id)
}

// UpdateTeacher godoc
// @Router                /teacher/{id} [PUT]
// @Summary 			  update a teacher
// @Description:          this api updates teacher information
// @Tags 			      teacher
// @Accept 			      json
// @Produce 		      json
// @Param 			      id path string true "Teacher ID"
// @Param       		  car body models.UpdateTeacher true "teacher"
// @Success 		      200 {object} models.Teacher
// @Failure 		      400 {object} models.Response
// @Failure               404 {object} models.Response
// @Failure 		      500 {object} models.Response
func (h Handler) UpdateTeacher(c *gin.Context) {
	teacher := models.Teacher{}
	if err := c.ShouldBindJSON(&teacher); err != nil {
		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}
	teacher.Id = c.Query("id")
	err := uuid.Validate(teacher.Id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	hashedPass, err := password.HashPassword(teacher.Password)
	if err != nil {
		handleResponseLog(c, h.Log, "error while generating customer password", http.StatusInternalServerError, err.Error())
		return
	}
	teacher.Password = string(hashedPass)

	id, err := h.Service.Teacher().Update(ctx, teacher)
	if err != nil {
		handleResponseLog(c, h.Log, "error while updating teacher", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "updated teacher", http.StatusOK, id)
}

// GetAllTeachers godoc
// @Router 			/teacher [GET]
// @Summary 		get all teachers
// @Description 	This API returns teacher list
// @Tags 			teacher
// Accept			json
// @Produce 		json
// @Param 			page query int false "page number"
// @Param 			limit query int false "limit per page"
// @Param 			search query string false "search keyword"
// @Success 		200 {object} models.GetAllTeachersResponse
// @Failure 		400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure 		500 {object} models.Response
func (h Handler) GetAllTeacher(c *gin.Context) {
	var (
		request = models.GetAllTeachersRequest{}
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

	teachers, err := h.Service.Teacher().GetAll(ctx, request)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting teachers", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, teachers)
}

// GetByIDTEacher godoc
// @Router       /teacher/{id} [GET]
// @Summary      return a teacher by ID
// @Description  Retrieves a teacher by its ID
// @Tags         teacher
// @Accept       json
// @Produce      json
// @Param        id path string true "Teacher ID"
// @Success      200 {object} models.GetTeacher
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDTeacher(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	teacher, err := h.Service.Teacher().GetByID(ctx, id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting teacher by id", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, teacher)
}

// DeleteTeacher godoc
// @Router          /teacher/{id} [DELETE]
// @Summary         delete a teacher by ID
// @Description     Deletes a teacher by its ID
// @Tags            teacher
// @Accept          json
// @Produce         json
// @Param           id path string true "Teacher ID"
// @Success         200 {string} models.Response
// @Failure         400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure         500 {object} models.Response
func (h Handler) DeleteTeacher(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	err = h.Service.Teacher().Delete(ctx, id)
	if err != nil {
		fmt.Println("error while deleting teacher, err:", err)
		handleResponseLog(c, h.Log, "error while deleting teacher", http.StatusInternalServerError, err)
		return
	}
	handleResponseLog(c, h.Log, "teacher deleted", http.StatusOK, id)
}
