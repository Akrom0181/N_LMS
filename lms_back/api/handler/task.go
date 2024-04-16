package handler

import (
	"context"
	"fmt"
	_ "lms_back/api/docs"
	"lms_back/api/models"
	"lms_back/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateTask godoc
// @Router 		/task [POST]
// @Summary 	create a task
// @Description This api is creates a new task and returns its id
// @Tags 		task
// @Accept		json
// @Produce		json
// @Param		car body models.CreateTask true "task"
// @Success		200  {object}  models.CreateTask
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateTask(c *gin.Context) {
	task := models.Task{}

	if err := c.ShouldBindJSON(&task); err != nil {

		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	id, err := h.Service.Task().Create(ctx, task)
	if err != nil {
		handleResponseLog(c, h.Log, "error while creating task", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "created successfully", http.StatusOK, id)
}

// UpdateTask godoc
// @Router                /task/{id} [PUT]
// @Summary 			  update a task
// @Description:          this api updates task information
// @Tags 			      task
// @Accept 			      json
// @Produce 		      json
// @Param 			      id path string true "Task ID"
// @Param       		  task body models.UpdateTask true "task"
// @Success 		      200 {object} models.Task
// @Failure 		      400 {object} models.Response
// @Failure               404 {object} models.Response
// @Failure 		      500 {object} models.Response
func (h Handler) UpdateTask(c *gin.Context) {

	task := models.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	task.Id = c.Query("id")
	err := uuid.Validate(task.Id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	id, err := h.Service.Task().Update(ctx, task)
	if err != nil {
		handleResponseLog(c, h.Log, "error while updating task", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "updated successfully", http.StatusOK, id)
}

// GetAlltasks godoc
// @Router 			/task [GET]
// @Summary 		get all tasks
// @Description 	This API returns task list
// @Tags 			task
// Accept			json
// @Produce 		json
// @Param 			page query int false "page number"
// @Param 			limit query int false "limit per page"
// @Param 			search query string false "search keyword"
// @Success 		200 {object} models.GetAllTasksResponse
// @Failure 		400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure 		500 {object} models.Response
func (h Handler) GetAllTask(c *gin.Context) {
	var (
		request = models.GetAllTasksRequest{}
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

	task, err := h.Service.Task().GetAll(ctx, request)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting task", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, task)
}

// GetByIDTask godoc
// @Router       /task/{id} [GET]
// @Summary      return a task by ID
// @Description  Retrieves a task by its ID
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        id path string true "Task ID"
// @Success      200 {object} models.GetTask
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDtask(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	task, err := h.Service.Task().GetByID(ctx, id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting task by id", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, task)
}

// DeleteTask godoc
// @Router          /task/{id} [DELETE]
// @Summary         delete a task by ID
// @Description     Deletes a task by its ID
// @Tags            task
// @Accept          json
// @Produce         json
// @Param           id path string true "Task ID"
// @Success         200 {string} models.Response
// @Failure         400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure         500 {object} models.Response
func (h Handler) DeleteTask(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	err = h.Service.Task().Delete(ctx, id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while deleting task", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "deleted task", http.StatusOK, id)
}
