package service

import (
	"context"
	"lms_back/api/models"
	"lms_back/pkg/logger"
	"lms_back/storage"
)

type adminService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewAdminService(storage storage.IStorage, logger logger.ILogger) adminService {
	return adminService{
		storage: storage,
		logger:  logger,
	}
}

func (u adminService) Create(ctx context.Context, admin models.Admin) (models.Admin, error) {

	pKey, err := u.storage.Admin().Create(ctx, admin)
	if err != nil {
		u.logger.Error("failed to create admin", logger.Error(err))
		return models.Admin{}, err
	}

	return pKey, nil
}

func (u adminService) Update(ctx context.Context, admin models.Admin) (models.Admin, error) {

	pKey, err := u.storage.Admin().Update(ctx, admin)
	if err != nil {
		u.logger.Error("failed to update admin", logger.Error(err))

		return models.Admin{}, err
	}

	return pKey, nil
}

func (u adminService) GetByID(ctx context.Context, id string) (models.Admin, error) {

	pKey, err := u.storage.Admin().GetByID(ctx, id)
	if err != nil {
		u.logger.Error("failed to get admin by ID", logger.Error(err))
		return models.Admin{}, err
	}

	return pKey, nil
}

func (u adminService) GetAll(ctx context.Context, req models.GetAllAdminsRequest) (models.GetAllAdminsResponse, error) {

	pKey, err := u.storage.Admin().GetAll(ctx, req)
	if err != nil {
		u.logger.Error("failed to get all cars", logger.Error(err))
		return models.GetAllAdminsResponse{}, err
	}

	return pKey, nil
}

func (u adminService) Delete(ctx context.Context, id string) error {

	err := u.storage.Admin().Delete(ctx, id)
	if err != nil {
		u.logger.Error("failed to delete car", logger.Error(err))
		return err
	}

	return nil
}

func (u adminService) GetByIdAdminReport(ctx context.Context, id models.AdminKey) ([]models.AdminPayment, error) {

	pKey, err := u.storage.AdminReport().GetByIDAdminPayment(ctx, id)
	if err != nil {
		u.logger.Error("ERROR in service layer while getting adminPayment", logger.Error(err))
		return []models.AdminPayment{}, err
	}

	return pKey, nil
}
