package usecase

import "goclean/repo"

type LoginUsecase interface {
}

type loginUsecaseImpl struct {
	userRepo repo.UserRepo
}

func NewLoginUsecase(userrepo repo.UserRepo) LoginUsecase {
	return &loginUsecaseImpl{
		userRepo: userrepo,
	}

}
