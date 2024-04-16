package service

import (
	"context"
	"lms_back/api/models"
	"lms_back/pkg/logger"
	"lms_back/storage"
)

type teacherService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewTeacherService(storage storage.IStorage, logger logger.ILogger) teacherService {
	return teacherService{
		storage: storage,
		logger:  logger,
	}
}

func (u teacherService) Create(ctx context.Context, teacher models.Teacher) (models.Teacher, error) {

	pKey, err := u.storage.Teacher().Create(ctx, teacher)
	if err != nil {
		u.logger.Error("ERROR in service layer while creating teacher", logger.Error(err))
		return models.Teacher{}, err
	}
	return pKey, nil
}

func (u teacherService) Update(ctx context.Context, teacher models.Teacher) (models.Teacher, error) {

	pKey, err := u.storage.Teacher().Update(ctx, teacher)
	if err != nil {
		u.logger.Error("ERROR in service layer while updating teacher", logger.Error(err))
		return models.Teacher{}, err
	}
	return pKey, nil
}

func (u teacherService) GetByID(ctx context.Context, id string) (models.Teacher, error) {

	pKey, err := u.storage.Teacher().GetByID(ctx, id)
	if err != nil {
		u.logger.Error("ERROR in service layer while getbyid teacher", logger.Error(err))
		return models.Teacher{}, err
	}

	return pKey, nil
}

func (u teacherService) GetAll(ctx context.Context, req models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error) {

	pKey, err := u.storage.Teacher().GetAll(ctx, req)
	if err != nil {
		u.logger.Error("ERROR in service layer while GetAll teacher", logger.Error(err))
		return models.GetAllTeachersResponse{}, err
	}

	return pKey, nil
}

func (u teacherService) Delete(ctx context.Context, id string) error {

	err := u.storage.Teacher().Delete(ctx, id)
	if err != nil {
		u.logger.Error("ERROR in service layer while deleting teacher", logger.Error(err))
		return err
	}

	return nil
}
