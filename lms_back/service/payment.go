package service

import (
	"context"
	"lms_back/api/models"
	"lms_back/pkg/logger"
	"lms_back/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type paymentService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewPaymentService(storage storage.IStorage, logger logger.ILogger) paymentService {
	return paymentService{
		storage: storage,
		logger:  logger,
	}
}

func (u paymentService) Create(ctx context.Context, payment models.CreatePayment) (resp models.Payment, err error) {

	pKey, err := u.storage.Payment().Create(ctx, payment)

	if err != nil {

		return models.Payment{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if payment.Student_id != "" && payment.Price > 0 {
		student, err := u.storage.Student().GetByID(ctx, pKey.Student_id)
		if err != nil {
			return models.Payment{}, status.Error(codes.InvalidArgument, err.Error())
		}

		student.PaidSum += payment.Price

		_, err = u.storage.Student().Update(ctx, models.Student{
			ID:      payment.Student_id,
			PaidSum: student.PaidSum,
			GroupID: student.GroupID,
			Status:  student.Status,
		})

		if err != nil {
			return models.Payment{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	resp, err = u.storage.Payment().GetByID(ctx, pKey.Id)
	if err != nil {
		return models.Payment{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return
}

func (u paymentService) Update(ctx context.Context, payment models.Payment) (models.Payment, error) {

	pKey, err := u.storage.Payment().Update(ctx, payment)
	if err != nil {
		u.logger.Error("ERROR in service layer while updating payment", logger.Error(err))
		return models.Payment{}, err
	}

	return pKey, nil
}

func (u paymentService) GetByID(ctx context.Context, id string) (models.Payment, error) {

	pKey, err := u.storage.Payment().GetByID(ctx, id)
	if err != nil {
		u.logger.Error("ERROR in service layer while getbyid payment", logger.Error(err))
		return models.Payment{}, err
	}

	return pKey, nil
}

func (u paymentService) GetAll(ctx context.Context, req models.GetAllPaymentsRequest) (models.GetAllPaymentsResponse, error) {

	pKey, err := u.storage.Payment().GetAll(ctx, req)
	if err != nil {
		u.logger.Error("ERROR in service layer while GetAll payment", logger.Error(err))
		return models.GetAllPaymentsResponse{}, err
	}

	return pKey, nil
}

func (u paymentService) Delete(ctx context.Context, id string) error {

	err := u.storage.Payment().Delete(ctx, id)
	if err != nil {
		u.logger.Error("ERROR in service layer while deleting payment", logger.Error(err))
		return err
	}

	return nil
}
