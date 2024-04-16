package handler

import (
	"context"
	"fmt"
	_ "lms_back/api/docs"
	"lms_back/api/models"
	"lms_back/pkg/password"
	"lms_back/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateAdmin godoc
// @Router     /admin [POST]
// @Summary    create a admin
// Description This api creates a new admin return its id
// @Tags       admin
// @Accept     json
// @Produce    json
// @Param      admin body   models.CreateAdmin true "admin"
// @Success    200 {object} models.Admin
// @Failure    400 {object} models.Response
// @Failure    404 {object} models.Response
// @Failure    500 {object} models.Response
func (h Handler) CreateAdmin(c *gin.Context) {
	admin := models.Admin{}

	if err := c.ShouldBindJSON(&admin); err != nil {
		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	hashedPass, err := password.HashPassword(admin.Password)
	if err != nil {
		handleResponseLog(c, h.Log, "error while generating customer password", http.StatusInternalServerError, err.Error())
		return
	}
	admin.Password = string(hashedPass)

	id, err := h.Service.Admin().Create(ctx, admin)
	if err != nil {
		handleResponseLog(c, h.Log, "error while creating admin", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "created admin", http.StatusOK, id)
}

// UpdateAdmin godoc
// @Router                /admin/{id} [PUT]
// @Summary               Update Admin
// @Description           This API Updates admin Information
// @Tags   	  			  admin
// @Accept     	          json
// @Produce               json
// @Param				  id path string true "Admin Id"
// @Param                 admin body models.UpdateAdmin true "admin"
// @Success 		      200 {object} models.Admin
// @Failure 		      400 {object} models.Response
// @Failure               404 {object} models.Response
// @Failure 		      500 {object} models.Response
func (h Handler) UpdateAdmin(c *gin.Context) {
	admin := models.Admin{}
	if err := c.ShouldBindJSON(&admin); err != nil {
		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}
	admin.Id = c.Param("id")
	err := uuid.Validate(admin.Id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating", http.StatusBadRequest, err.Error())
		return
	}

	hashedPass, err := password.HashPassword(admin.Password)
	if err != nil {
		handleResponseLog(c, h.Log, "error while generating customer password", http.StatusInternalServerError, err.Error())
		return
	}
	admin.Password = string(hashedPass)

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	id, err := h.Service.Admin().Update(ctx, admin)
	if err != nil {
		handleResponseLog(c, h.Log, "error while updating admin", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "updated successfully", http.StatusOK, id)
}

// GetAllAdmin godoc
// @Router 			/admin [GET]
// @Summary 		get all admin
// @Description 	This API returns admin list
// @Tags 			admin
// Accept			json
// @Produce 		json
// @Param 			page query int false "page number"
// @Param 			limit query int false "limit per page"
// @Param 			search query string false "search keyword"
// @Success 		200 {object} models.GetAllAdminsResponse
// @Failure 		400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure 		500 {object} models.Response
func (h Handler) GetAllAdmins(c *gin.Context) {
	var (
		request = models.GetAllAdminsRequest{}
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

	admins, err := h.Service.Admin().GetAll(ctx, request)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting admins", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, admins)
}

// GetByIDAdmin godoc
// @Router       /admin/{id} [GET]
// @Summary      return a admin by ID
// @Description  Retrieves a admin by its ID
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        id path string true "Admin ID"
// @Success      200 {object} models.GetAdmin
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDAdmin(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	admin, err := h.Service.Admin().GetByID(ctx, id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting admin by id", http.StatusInternalServerError, err)
		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, admin)
}

// DeleteAdmin godoc
// @Router          /admin/{id} [DELETE]
// @Summary         delete a admin by ID
// @Description     Deletes a admin by its ID
// @Tags            admin
// @Accept          json
// @Produce         json
// @Param           id path string true "Admin ID"
// @Success         200 {string} models.Response
// @Failure         400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure         500 {object} models.Response
func (h Handler) DeleteAdmin(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	err = h.Service.Admin().Delete(ctx, id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while deleting admin", http.StatusInternalServerError, err)
		return
	}
	handleResponseLog(c, h.Log, "deleted admin", http.StatusOK, id)
}

// GetById AdminPayment godoc
// @Router       /adminPay/{idAdmin} [GET]
// @Summary      return a admin by payments
// @Description  Retrieves a admin by its payments
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        id path string true "Admin ID"
// @Success      200 {object} models.GetAdmin
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIdAdminReport(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)
	adminPay := models.AdminKey{
		Id: id,
	}

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	admin, err := h.Service.Admin().GetByIdAdminReport(ctx, adminPay)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting admin by id", http.StatusInternalServerError, err)
		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, admin)
}
