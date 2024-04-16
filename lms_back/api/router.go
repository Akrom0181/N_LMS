package api

import (
	"errors"
	"lms_back/api/handler"
	"lms_back/pkg/logger"
	"lms_back/service"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @name Authorization
func New(service service.IServiceManager, log logger.ILogger) *gin.Engine {
	h := handler.NewStrg(service, log)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(authMiddleware)

	r.GET("/admin", h.GetAllAdmins)
	r.GET("/admin/:id", h.GetByIDAdmin)
	r.POST("/admin", h.CreateAdmin)
	r.PUT("/admin/:id", h.UpdateAdmin)
	r.DELETE("/admin/:id", h.DeleteAdmin)
	r.GET("adminPay/:id", h.GetByIdAdminReport)

	r.GET("/branch", h.GetAllBranches)
	r.GET("/branch/:id", h.GetByIDBranch)
	r.POST("/branch", h.CreateBranch)
	r.PUT("/branch/:id", h.UpdateBranch)
	r.DELETE("/branch/:id", h.DeleteBranch)

	r.GET("/group", h.GetAllGroups)
	r.GET("/group/:id", h.GetByIDGroup)
	r.POST("/group", h.CreateGroup)
	r.PUT("/group/:id", h.UpdateGroup)
	r.DELETE("/group/:id", h.DeleteGroup)

	r.GET("/lesson", h.GetAllLessons)
	r.GET("/lesson/:id", h.GetByIDLesson)
	r.POST("/lesson", h.CreateLesson)
	r.PUT("/lesson/:id", h.UpdateLesson)
	r.DELETE("/lesson/:id", h.DeleteLessson)

	r.GET("/payment", h.GetAllPayment)
	r.GET("/payment/:id", h.GetByIDPayment)
	r.POST("/payment", h.CreatePayment)
	r.PUT("/payment/:id", h.UpdatePayment)
	r.DELETE("/payment/:id", h.DeletePayment)

	r.GET("/schedule", h.GetAllSchedule)
	r.GET("/schedule/:id", h.GetByIDSchedule)
	r.POST("/schedule", h.CreateSchedule)
	r.PUT("/schedule/:id", h.UpdateSchedule)
	r.DELETE("/schedule/:id", h.DeleteSchedule)

	r.POST("/student/login", h.StudentLogin)

	r.GET("/student", h.GetAllStudent)
	r.GET("/student/:id", h.GetByIDStudent)
	r.POST("/student", h.CreateStudent)
	r.PUT("/student/:id", h.UpdateStudent)
	r.DELETE("/student/:id", h.DeleteStudent)

	r.GET("/task", h.GetAllTask)
	r.GET("/task/:id", h.GetByIDtask)
	r.POST("/task", h.CreateTask)
	r.PUT("/task/:id", h.UpdateTask)
	r.DELETE("/task/:id", h.DeleteTask)

	r.GET("/teacher", h.GetAllTeacher)
	r.GET("/teacher/:id", h.GetByIDTeacher)
	r.POST("/teacher", h.CreateTeacher)
	r.PUT("/teacher/:id", h.UpdateTeacher)
	r.DELETE("/teacher/:id", h.DeleteTeacher)

	return r
}

func authMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
	}
	c.Next()
}
