package repo

import (
	"database/sql"
	"fmt"
	"goclean/model"
	"goclean/utils"
)

type ServiceRepo interface {
	GetServiceById(int) (*model.ServiceModel, error)
	GetServiceByName(string) (*model.ServiceModel, error)
	InsertService(*model.ServiceModel) error
}

type serviceRepoImpl struct {
	db *sql.DB
}

func (svcRepo *serviceRepoImpl) InsertService(svc *model.ServiceModel) error {
	qry := utils.INSERT_SERVICE

	_, err := svcRepo.db.Exec(qry, svc.Name, svc.Uom, svc.Price)
	if err != nil {
		return fmt.Errorf("error on serviceRepoImpl.InsertService() : %w", err)
	}
	return nil
}

func (svcRepo *serviceRepoImpl) GetServiceByName(name string) (*model.ServiceModel, error) {
	qry := utils.GET_SERVICE_BY_NAME
	svc := &model.ServiceModel{}
	err := svcRepo.db.QueryRow(qry, name).Scan(&svc.Id, &svc.Name, &svc.Uom, &svc.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on serviceRepoImpl.GetServiceByName() : %w", err)
	}
	return svc, nil
}

func (svcRepo *serviceRepoImpl) GetServiceById(id int) (*model.ServiceModel, error) {
	qry := utils.GET_SERVICE_BY_ID

	svc := &model.ServiceModel{}
	err := svcRepo.db.QueryRow(qry, id).Scan(&svc.Id, &svc.Name, &svc.Uom, &svc.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on serviceRepoImpl.getServiceById() : %w", err)
	}
	return svc, nil
}

func NewServiceRepo(db *sql.DB) ServiceRepo {
	return &serviceRepoImpl{
		db: db,
	}
}
