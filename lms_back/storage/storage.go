package storage

import (
	"lms_back/api/models"
)

type IStorage interface {
	CloseDB()
	Admin()    IAdminStorage
	Branches() IBranchStorage
	Group()    IGroupStorage
	Payment()  IPaymentStorage
	Student()  IStudentStorage
	Teacher()  ITeacherStorage
	Schedule() IScheduleStorage
	Task()     ITaskStorage
	Lesson()   ILessonStorage
}

type IAdminStorage interface {
	Create(models.Admin) (models.Admin, error)
	GetAll(request models.GetAllAdminsRequest) (models.GetAllAdminsResponse, error)
	GetByID(id string) (models.Admin, error)
	Update(models.Admin) (models.Admin, error)
	Delete(string) error
}

type IBranchStorage interface {
	Create(models.Branch) (models.Branch, error)
	GetAll(request models.GetAllBranchesRequest) (models.GetAllBranchesResponse, error)
	GetByID(id string) (models.Branch, error)
	Update(models.Branch) (models.Branch, error)
	Delete(string) error
}

type IGroupStorage interface {
	Create(models.Group) (models.Group, error)
	GetAll(request models.GetAllGroupsRequest) (models.GetAllGroupsResponse, error)
	GetByID(id string) (models.Group, error)
	Update(models.Group) (models.Group, error)
	Delete(string) error
}

type IPaymentStorage interface {
	Create(models.Payment) (models.Payment, error)
	GetAll(request models.GetAllPaymentsRequest) (models.GetAllPaymentsResponse, error)
	GetByID(id string) (models.Payment, error)
	Update(models.Payment) (models.Payment, error)
	Delete(string) error
}

type IStudentStorage interface {
	Create(models.Student) (models.Student, error)
	GetAll(request models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error)
	GetByID(id string) (models.Student, error)
	Update(models.Student) (models.Student, error)
	Delete(string) error
}

type ITeacherStorage interface {
	Create(models.Teacher) (models.Teacher, error)
	GetAll(request models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error)
	GetByID(id string) (models.Teacher, error)
	Update(models.Teacher) (models.Teacher, error)
	Delete(string) error
}

type IScheduleStorage interface {
	Create(models.Schedule) (models.Schedule, error)
	GetAll(request models.GetAllSchedulesRequest) (models.GetAllSchedulesResponse, error)
	GetByID(id string) (models.Schedule, error)
	Update(models.Schedule) (models.Schedule, error)
	Delete(string) error
}

type ITaskStorage interface {
	Create(models.Task) (models.Task, error)
	GetAll(req models.GetAllTasksRequest) (models.GetAllTasksResponse, error)
	GetByID(id string) (models.Task, error)
	Update(models.Task) (models.Task, error)
	Delete(string) error
}

type ILessonStorage interface {
	Create(models.Lesson) (models.Lesson, error)
	GetAll(req models.GetAllLessonsRequest) (models.GetAllLessonsResponse, error)
	GetByID(id string) (models.Lesson, error)
	Update(models.Lesson) (models.Lesson, error)
	Delete(string) error
}
