package service

import (
	"context"
	"lms_back/api/models"
	"lms_back/pkg/logger"
	"lms_back/storage"
)

type taskService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewTaskService(storage storage.IStorage, logger logger.ILogger) taskService {
	return taskService{
		storage: storage,
		logger:  logger,
	}
}

func (u taskService) Create(ctx context.Context, task models.Task) (models.Task, error) {

	pKey, err := u.storage.Task().Create(ctx, task)
	if err != nil {
		u.logger.Error("ERROR in service layer while creating task", logger.Error(err))
		return models.Task{}, err
	}

	return pKey, nil
}

func (u taskService) Update(ctx context.Context, task models.Task) (models.Task, error) {

	pKey, err := u.storage.Task().Update(ctx, task)
	if err != nil {
		u.logger.Error("ERROR in service layer while updating task", logger.Error(err))
		return models.Task{}, err
	}
	return pKey, nil
}

func (u taskService) GetByID(ctx context.Context, id string) (models.Task, error) {

	pKey, err := u.storage.Task().GetByID(ctx, id)
	if err != nil {
		u.logger.Error("ERROR in service layer while getbyid task", logger.Error(err))
		return models.Task{}, err
	}

	return pKey, nil
}

func (u taskService) GetAll(ctx context.Context, req models.GetAllTasksRequest) (models.GetAllTasksResponse, error) {

	pKey, err := u.storage.Task().GetAll(ctx, req)
	if err != nil {
		u.logger.Error("ERROR in service layer while GetAll task", logger.Error(err))
		return models.GetAllTasksResponse{}, err
	}

	return pKey, nil
}

func (u taskService) Delete(ctx context.Context, id string) error {

	err := u.storage.Task().Delete(ctx, id)
	if err != nil {
		u.logger.Error("ERROR in service layer while deleting task", logger.Error(err))
		return err
	}

	return nil
}
