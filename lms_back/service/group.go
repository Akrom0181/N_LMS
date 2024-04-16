package service

import (
	"context"
	"lms_back/api/models"
	"lms_back/pkg/logger"
	"lms_back/storage"
)

type groupService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewGroupService(storage storage.IStorage, logger logger.ILogger) groupService {
	return groupService{
		storage: storage,
		logger:  logger,
	}
}

func (u groupService) Create(ctx context.Context, group models.Group) (models.Group, error) {

	pKey, err := u.storage.Group().Create(ctx, group)
	if err != nil {
		u.logger.Error("ERROR in service layer while creating car", logger.Error(err))
		return models.Group{}, err
	}

	return pKey, nil
}

func (u groupService) Update(ctx context.Context, group models.Group) (models.Group, error) {

	pKey, err := u.storage.Group().Update(ctx, group)
	if err != nil {
		u.logger.Error("ERROR in service layer while updating group", logger.Error(err))
		return models.Group{}, err
	}

	return pKey, nil
}

func (u groupService) GetByID(ctx context.Context, id string) (models.Group, error) {

	pKey, err := u.storage.Group().GetByID(ctx, id)
	if err != nil {
		u.logger.Error("ERROR in service layer while getbyid group", logger.Error(err))
		return models.Group{}, err
	}

	return pKey, nil
}

func (u groupService) GetAll(ctx context.Context, req models.GetAllGroupsRequest) (models.GetAllGroupsResponse, error) {

	pKey, err := u.storage.Group().GetAll(ctx, req)
	if err != nil {
		u.logger.Error("ERROR in service layer while GetAll group", logger.Error(err))
		return models.GetAllGroupsResponse{}, err
	}

	return pKey, nil
}

func (u groupService) Delete(ctx context.Context, id string) error {

	err := u.storage.Group().Delete(ctx, id)
	if err != nil {
		u.logger.Error("ERROR in service layer while deleting group", logger.Error(err))
		return err
	}

	return nil
}
