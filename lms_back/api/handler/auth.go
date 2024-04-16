package handler

import (
	"fmt"
	"lms_back/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// StudentLogin godoc
// @Router       /student/login [POST]
// @Summary      Customer login
// @Description  Customer login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body     models.StudentLoginRequest true "login"
// @Success      201  {object}  models.StudentLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) StudentLogin(c *gin.Context) {
	loginReq := models.StudentLoginRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponseLog(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}
	fmt.Println("loginReq: ", loginReq)

	loginResp, err := h.Service.Auth().StudentLogin(c.Request.Context(), loginReq)
	if err != nil {
		handleResponseLog(c, h.Log, "unauthorized", http.StatusUnauthorized, err)
		return
	}

	handleResponseLog(c, h.Log, "Succes", http.StatusOK, loginResp)

}
