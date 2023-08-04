package handler

import (
	"goclean/manager"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	usecaseManager manager.UsecaseManager
	srv            *gin.Engine
}

func (s *server) Run() {
	s.srv.Use(LoggerMiddleware())
	NewServiceHandler(s.srv, s.usecaseManager.GetServiceUsecase())
	NewUserHandler(s.srv, s.usecaseManager.GetUserUsecase())
	NewLoginHandler(s.srv, s.usecaseManager.GetUserUsecase())
	s.srv.Run()
}

func NewServer() Server {
	infra := manager.NewInfraManager()
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)

	srv := gin.Default()

	return &server{
		usecaseManager: usecase,
		srv:            srv,
	}

}
