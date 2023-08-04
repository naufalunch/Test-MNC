package usecase

import (
	"fmt"
	"goclean/apperror"
	"goclean/model"
	"goclean/repo"
)

type ServiceUsecase interface {
	GetServiceById(int) (*model.ServiceModel, error)
	InsertService(*model.ServiceModel) error
}

type serviceUsecaseImpl struct {
	svcRepo repo.ServiceRepo
}

func (svcUsecase *serviceUsecaseImpl) GetServiceById(id int) (*model.ServiceModel, error) {
	return svcUsecase.svcRepo.GetServiceById(id)
}

func (svcUsecase *serviceUsecaseImpl) InsertService(svc *model.ServiceModel) error {
	svcDB, err := svcUsecase.svcRepo.GetServiceByName(svc.Name)
	if err != nil {
		return fmt.Errorf("serviceUsecaseImpl.InsertService() : %w", err)
	}

	if svcDB != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data service dengan nama %v sudah ada", svc.Name),
		}
	}

	return svcUsecase.svcRepo.InsertService(svc)
}

func NewServiceUseCase(svcRepo repo.ServiceRepo) ServiceUsecase {
	return &serviceUsecaseImpl{
		svcRepo: svcRepo,
	}
}
