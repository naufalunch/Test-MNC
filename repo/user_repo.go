package repo

import (
	"database/sql"
	"fmt"
	"goclean/model"
	"goclean/utils"
)

type UserRepo interface {
	GetUserById(int) (*model.UserModel, error)
	GetUserByName(string) (*model.UserModel, error)
	InsertUser(*model.UserModel) error
	GetAllUser() ([]model.UserModel, error)
}

type userRepoImpl struct {
	db *sql.DB
}

func (usrRepo *userRepoImpl) InsertUser(usr *model.UserModel) error {
	qry := utils.INSERT_USER
	_, err := usrRepo.db.Exec(qry, usr.Id, usr.Username, usr.Password, usr.Active)
	if err != nil {
		return fmt.Errorf("error on userRepoImpl.Insertuser() : %w", err)
	}
	return nil
}

func (usrRepo *userRepoImpl) GetUserById(id int) (*model.UserModel, error) {
	qry := utils.GET_USER_BY_ID

	usr := &model.UserModel{}
	err := usrRepo.db.QueryRow(qry, id).Scan(&usr.Id, &usr.Username, &usr.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on userRepoImpl.getUserById() : %w", err)
	}
	return usr, nil
}

func (usrRepo *userRepoImpl) GetAllUser() ([]model.UserModel, error) {
	qry := utils.GET_ALL_USER

	usr := &model.UserModel{}
	rows, err := usrRepo.db.Query(qry)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on userRepoImpl.GetAllUser() : %w", err)
	}

	arrSvc := []model.UserModel{}
	for rows.Next() {
		rows.Scan(&usr.Id, &usr.Username, &usr.Password, &usr.Active)
		arrSvc = append(arrSvc, *usr)
	}
	return arrSvc, nil
}

func (usrRepo *userRepoImpl) GetUserByName(name string) (*model.UserModel, error) {
	qry := utils.GET_USER_BY_NAME

	usr := &model.UserModel{}
	err := usrRepo.db.QueryRow(qry, name).Scan(&usr.Id, &usr.Username, &usr.Password, &usr.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on userRepoImpl.GetUserByName() : %w", err)
	}
	return usr, nil
}
func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepoImpl{
		db: db,
	}
}
