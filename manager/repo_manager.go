package manager

import (
	"goclean/repo"
	"sync"
)

type RepoManager interface {
	GetServiceRepo() repo.ServiceRepo
	GetUserRepo() repo.UserRepo
}

type repoManager struct {
	infraManager InfraManager

	svcRepo repo.ServiceRepo
	usrRepo repo.UserRepo
}

var onceLoadServiceRepo sync.Once
var onceLoadUserRepo sync.Once

func (rm *repoManager) GetServiceRepo() repo.ServiceRepo {
	onceLoadServiceRepo.Do(func() {
		rm.svcRepo = repo.NewServiceRepo(rm.infraManager.GetDB())
	})
	return rm.svcRepo
}

func (rm *repoManager) GetUserRepo() repo.UserRepo {
	onceLoadUserRepo.Do(func() {
		rm.usrRepo = repo.NewUserRepo(rm.infraManager.GetDB())
	})
	return rm.usrRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
