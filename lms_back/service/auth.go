package service

import (
	"context"
	"fmt"
	"lms_back/api/models"
	"lms_back/pkg/jwt"
	"lms_back/pkg/logger"
	"lms_back/pkg/password"
	"lms_back/config"
	"lms_back/storage"
)

type authService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewAuthService(storage storage.IStorage, log logger.ILogger) authService {
	return authService{
		storage: storage,
		log:     log,
	}
}

func (a authService) StudentLogin(ctx context.Context, loginRequest models.StudentLoginRequest) (models.StudentLoginResponse, error) {
	fmt.Println(" loginRequest.Login: ", loginRequest.Login)
	student, err := a.storage.Student().GetByLogin(ctx, loginRequest.Login)
	if err != nil {
		a.log.Error("error while getting student credentials by login", logger.Error(err))
		return models.StudentLoginResponse{}, err
	}

	if err = password.CompareHashAndPassword(student.Password, loginRequest.Password); err != nil {
		a.log.Error("error while comparing password", logger.Error(err))
		return models.StudentLoginResponse{}, err
	}

	m := make(map[interface{}]interface{})

	m["user_id"] = student.ID
	m["user_role"] = config.STUDENT_ROLE

	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		a.log.Error("error while generating tokens for student login", logger.Error(err))
		return models.StudentLoginResponse{}, err
	}

	return models.StudentLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
