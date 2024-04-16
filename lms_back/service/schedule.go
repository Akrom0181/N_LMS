package service

import (
	"context"
	"lms_back/api/models"
	"lms_back/pkg/logger"
	"lms_back/storage"
)

type scheduleService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewScheduleService(storage storage.IStorage, logger logger.ILogger) scheduleService {
	return scheduleService{
		storage: storage,
		logger:  logger,
	}
}

func (u scheduleService) Create(ctx context.Context, schedule models.Schedule) (models.Schedule, error) {

	pKey, err := u.storage.Schedule().Create(ctx, schedule)
	if err != nil {
		u.logger.Error("ERROR in service layer while creating schedule", logger.Error(err))
		return models.Schedule{}, err
	}

	return pKey, nil
}

func (u scheduleService) Update(ctx context.Context, schedule models.Schedule) (models.Schedule, error) {

	pKey, err := u.storage.Schedule().Update(ctx, schedule)
	if err != nil {
		u.logger.Error("ERROR in service layer while updating schedule", logger.Error(err))
		return models.Schedule{}, err
	}
	return pKey, nil
}

func (u scheduleService) GetByID(ctx context.Context, id string) (models.Schedule, error) {

	pKey, err := u.storage.Schedule().GetByID(ctx, id)
	if err != nil {
		u.logger.Error("ERROR in service layer while getbyid schedule", logger.Error(err))
		return models.Schedule{}, err
	}

	return pKey, nil
}

func (u scheduleService) GetAll(ctx context.Context, req models.GetAllSchedulesRequest) (models.GetAllSchedulesResponse, error) {

	pKey, err := u.storage.Schedule().GetAll(ctx, req)
	if err != nil {
		u.logger.Error("ERROR in service layer while GetAll schedule", logger.Error(err))
		return models.GetAllSchedulesResponse{}, err
	}

	return pKey, nil
}

func (u scheduleService) Delete(ctx context.Context, id string) error {

	err := u.storage.Schedule().Delete(ctx, id)
	if err != nil {
		u.logger.Error("ERROR in service layer while deleting schedule", logger.Error(err))
		return err
	}

	return nil
}
