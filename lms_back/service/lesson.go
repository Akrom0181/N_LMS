package service

import (
	"context"
	"lms_back/api/models"
	"lms_back/pkg/logger"
	"lms_back/storage"
)

type lessonService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewLessonService(storage storage.IStorage, logger logger.ILogger) lessonService {
	return lessonService{
		storage: storage,
		logger:  logger,
	}
}

func (u lessonService) Create(ctx context.Context, lesson models.Lesson) (models.Lesson, error) {

	pKey, err := u.storage.Lesson().Create(ctx, lesson)
	if err != nil {
		u.logger.Error("ERROR in service layer while creating lesson", logger.Error(err))
		return models.Lesson{}, err
	}

	return pKey, nil
}

func (u lessonService) Update(ctx context.Context, lesson models.Lesson) (models.Lesson, error) {

	pKey, err := u.storage.Lesson().Update(ctx, lesson)
	if err != nil {
		u.logger.Error("ERROR in service layer while updating lesson", logger.Error(err))
		return models.Lesson{}, err
	}

	return pKey, nil
}

func (u lessonService) GetByID(ctx context.Context, id string) (models.Lesson, error) {

	pKey, err := u.storage.Lesson().GetByID(ctx, id)
	if err != nil {
		u.logger.Error("ERROR in service layer while getbyid lesson", logger.Error(err))
		return models.Lesson{}, err
	}

	return pKey, nil
}

func (u lessonService) GetAll(ctx context.Context, req models.GetAllLessonsRequest) (models.GetAllLessonsResponse, error) {

	pKey, err := u.storage.Lesson().GetAll(ctx, req)
	if err != nil {
		u.logger.Error("ERROR in service layer while GetAll lesson", logger.Error(err))
		return models.GetAllLessonsResponse{}, err
	}

	return pKey, nil
}

func (u lessonService) Delete(ctx context.Context, id string) error {

	err := u.storage.Lesson().Delete(ctx, id)
	if err != nil {
		u.logger.Error("ERROR in service layer while deleting lesson", logger.Error(err))
		return err
	}

	return nil
}
