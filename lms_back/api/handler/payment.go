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

// CreatePayment godoc
// @Router 		/payment [POST]
// @Summary 	create a payment
// @Description This api is creates a new payment and returns its id
// @Tags 		payment
// @Accept		json
// @Produce		json
// @Param		payment body models.CreatePayment true "payment"
// @Success		200  {object}  models.Payment
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreatePayment(c *gin.Context) {
	payment := models.CreatePayment{}

	if err := c.ShouldBindJSON(&payment); err != nil {
		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	id, err := h.Service.Payment().Create(ctx, payment)
	if err != nil {
		handleResponseLog(c, h.Log, "error while creating payment", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponseLog(c, h.Log, "created successfully", http.StatusOK, id)
}

// UpdatePayment godoc
// @Router                /payment/{id} [PUT]
// @Summary 			  update a payment
// @Description:          this api updates payment information
// @Tags 			      payment
// @Accept 			      json
// @Produce 		      json
// @Param 			      id path string true "Payment ID"
// @Param       		  payment body models.UpdatePayment true "payment"
// @Success 		      200 {object} models.Payment
// @Failure 		      400 {object} models.Response
// @Failure               404 {object} models.Response
// @Failure 		      500 {object} models.Response
func (h Handler) UpdatePayment(c *gin.Context) {
	payment := models.Payment{}
	if err := c.ShouldBindJSON(&payment); err != nil {
		handleResponseLog(c, h.Log, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	payment.Id = c.Param("id")
	err := uuid.Validate(payment.Id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	id, err := h.Service.Payment().Update(ctx, payment)
	if err != nil {
		handleResponseLog(c, h.Log, "error while updating payment", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "updated payment", http.StatusOK, id)
}

// GetAllPayment godoc
// @Router 			/payment [GET]
// @Summary 		get all payment
// @Description 	This API returns payment list
// @Tags 			payment
// Accept			json
// @Produce 		json
// @Param 			page query int false "page number"
// @Param 			limit query int false "limit per page"
// @Param 			search query string false "search keyword"
// @Success 		200 {object} models.GetAllPaymentsResponse
// @Failure 		400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure 		500 {object} models.Response
func (h Handler) GetAllPayment(c *gin.Context) {
	var (
		request = models.GetAllPaymentsRequest{}
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

	payment, err := h.Service.Payment().GetAll(ctx, request)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting payment", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, payment)
}

// GetByIDpayment godoc
// @Router       /payment/{id} [GET]
// @Summary      return a payment by ID
// @Description  Retrieves a payment by its ID
// @Tags         payment
// @Accept       json
// @Produce      json
// @Param        id path string true "payment ID"
// @Success      200 {object} models.GetPayment
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDPayment(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	payment, err := h.Service.Payment().GetByID(ctx, id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while getting payment by id", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "", http.StatusOK, payment)
}

// DeletePayment godoc
// @Router          /payment/{id} [DELETE]
// @Summary         delete a payment by ID
// @Description     Deletes a payment by its ID
// @Tags            payment
// @Accept          json
// @Produce         json
// @Param           id path string true "Payment ID"
// @Success         200 {string} models.Response
// @Failure         400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure         500 {object} models.Response
func (h Handler) DeletePayment(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, config.Timeout)
	defer cancel()

	err = h.Service.Payment().Delete(ctx, id)
	if err != nil {
		handleResponseLog(c, h.Log, "error while deleting payment", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponseLog(c, h.Log, "successfully deletes", http.StatusOK, id)
}
