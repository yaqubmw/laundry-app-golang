package delivery

import (
	"enigma-laundry-apps/config"
	"enigma-laundry-apps/delivery/controller/api"
	"enigma-laundry-apps/delivery/middleware"
	"enigma-laundry-apps/manager"
	"enigma-laundry-apps/utils/exceptions"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
	log            *logrus.Logger
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initController() {
	s.engine.Use(middleware.LogRequestMiddleware(s.log))
	// semua controller disini
	api.NewUomController(s.useCaseManager.UomUseCase(), s.engine)
	api.NewProductController(s.engine, s.useCaseManager.ProductUseCase())
	api.NewCustomerController(s.engine, s.useCaseManager.CustomerUseCase())
	api.NewEmployeeController(s.engine, s.useCaseManager.EmployeeUseCase())
	api.NewBillController(s.engine, s.useCaseManager.BillUseCase())
	api.NewUserController(s.engine, s.useCaseManager.UserUseCase())
	api.NewAuthController(s.engine, s.useCaseManager.AuthUseCase())
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	exceptions.CheckErr(err)
	infraManager, _ := manager.NewInfraManager(cfg)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		useCaseManager: useCaseManager,
		engine:         engine,
		host:           host,
		log:            logrus.New(),
	}
}
