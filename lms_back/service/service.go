package service

import (
	"lms_back/pkg/logger"
	"lms_back/storage"
)

type IServiceManager interface {
	Admin() adminService
	Branch() branchService
	Group() groupService
	Lesson() lessonService
	Payment() paymentService
	Schedule() scheduleService
	Student() studentService
	Task() taskService
	Teacher() teacherService
	Auth() authService
}

type Service struct {
	adminService    adminService
	branchService   branchService
	groupService    groupService
	lessonService   lessonService
	paymentService  paymentService
	scheduleService scheduleService
	studentService  studentService
	taskService     taskService
	teacherService  teacherService
	authService     authService

	logger logger.ILogger
}

func New(storage storage.IStorage, log logger.ILogger) Service {
	return Service{
		adminService:    NewAdminService(storage, log),
		branchService:   NewBranchService(storage, log),
		groupService:    NewGroupService(storage, log),
		paymentService:  NewPaymentService(storage, log),
		scheduleService: NewScheduleService(storage, log),
		studentService:  NewStudentService(storage, log),
		taskService:     NewTaskService(storage, log),
		lessonService:   NewLessonService(storage, log),
		teacherService:  NewTeacherService(storage, log),

		authService:     NewAuthService(storage, log),
		logger:          log,
	}
}

func (s Service) Admin() adminService {
	return s.adminService
}

func (s Service) Branch() branchService {
	return s.branchService
}

func (s Service) Group() groupService {
	return s.groupService
}

func (s Service) Lesson() lessonService {
	return s.lessonService
}

func (s Service) Payment() paymentService {
	return s.paymentService
}

func (s Service) Schedule() scheduleService {
	return s.scheduleService
}

func (s Service) Student() studentService {
	return s.studentService
}

func (s Service) Task() taskService {
	return s.taskService
}

func (s Service) Teacher() teacherService {
	return s.teacherService
}

func (s Service) Auth() authService {
	return s.authService
}
