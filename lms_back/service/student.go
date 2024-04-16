package service

import (
	"context"
	"lms_back/api/models"
	"lms_back/pkg/logger"
	"lms_back/storage"
	"lms_back/pkg/password"


)

type studentService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewStudentService(storage storage.IStorage, logger logger.ILogger) studentService {
	return studentService{
		storage: storage,
		logger:  logger,
	}
}

func (s studentService) Login(ctx context.Context, req models.StudentLoginRequest) (string, error) {

	hashedPswd, err := s.storage.Student().GetPassword(ctx, req.Login)
	if err != nil {
		s.logger.Error("error while getting customer password", logger.Error(err))
		return "", err
	}

	err = password.CompareHashAndPassword(hashedPswd, req.Password)
	if err != nil {
		s.logger.Error("incorrect password", logger.Error(err))
		return "", err
	}
	return "Login successfully", nil
}

func (u studentService) Create(ctx context.Context, student models.Student) (models.Student, error) {

	pKey, err := u.storage.Student().Create(ctx, student)
	if err != nil {
		u.logger.Error("ERROR in service layer while creating student", logger.Error(err))
		return models.Student{}, err
	}

	return pKey, nil
}

func (u studentService) Update(ctx context.Context, student models.Student) (models.Student, error) {

	pKey, err := u.storage.Student().Update(ctx, student)
	if err != nil {
		u.logger.Error("ERROR in service layer while updating student", logger.Error(err))
		return models.Student{}, err
	}
	return pKey, nil
}

func (u studentService) GetByID(ctx context.Context, id string) (models.Student, error) {

	pKey, err := u.storage.Student().GetByID(ctx, id)
	if err != nil {
		u.logger.Error("ERROR in service layer while getbyid student", logger.Error(err))
		return models.Student{}, err
	}

	return pKey, nil
}

func (u studentService) GetAll(ctx context.Context, req models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error) {

	pKey, err := u.storage.Student().GetAll(ctx, req)
	if err != nil {
		u.logger.Error("ERROR in service layer while GetAll student", logger.Error(err))
		return models.GetAllStudentsResponse{}, err
	}

	return pKey, nil
}

func (u studentService) Delete(ctx context.Context, id string) error {

	err := u.storage.Student().Delete(ctx, id)
	if err != nil {
		u.logger.Error("ERROR in service layer while deleting student", logger.Error(err))
		return err
	}

	return nil
}
