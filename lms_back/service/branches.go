package service

import (
	"context"
	"lms_back/api/models"
	"lms_back/pkg/logger"
	"lms_back/storage"
)

type branchService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewBranchService(storage storage.IStorage, logger logger.ILogger) branchService {
	return branchService{
		storage: storage,
		logger:  logger,
	}
}

func (u branchService) Create(ctx context.Context, branch models.Branch) (models.Branch, error) {

	pKey, err := u.storage.Branch().Create(ctx, branch)
	if err != nil {
		u.logger.Error("error while creating branch in service layer", logger.Error(err))
		return models.Branch{}, err
	}

	return pKey, nil
}

func (u branchService) Update(ctx context.Context, branch models.Branch) (models.Branch, error) {

	pKey, err := u.storage.Branch().Update(ctx, branch)
	if err != nil {
		u.logger.Error("ERROR in service layer while updating branch", logger.Error(err))
		return models.Branch{}, err
	}

	return pKey, nil
}

func (u branchService) GetByID(ctx context.Context, id string) (models.Branch, error) {

	pKey, err := u.storage.Branch().GetByID(ctx, id)
	if err != nil {
		u.logger.Error("ERROR in service layer while getbyid branch", logger.Error(err))
		return models.Branch{}, err
	}

	return pKey, nil
}

func (u branchService) GetAll(ctx context.Context, req models.GetAllBranchesRequest) (models.GetAllBranchesResponse, error) {

	pKey, err := u.storage.Branch().GetAll(ctx, req)
	if err != nil {
		u.logger.Error("ERROR in service layer while GetAll branch", logger.Error(err))
		return models.GetAllBranchesResponse{}, err
	}

	return pKey, nil
}

func (u branchService) Delete(ctx context.Context, id string) error {

	err := u.storage.Branch().Delete(ctx, id)
	if err != nil {
		u.logger.Error("ERROR in service layer while deleting branch", logger.Error(err))
		return err
	}

	return nil
}
