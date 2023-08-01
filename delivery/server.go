package delivery

import (
	"enigma-laundry-apps/config"
	"enigma-laundry-apps/delivery/controller/api"
	"enigma-laundry-apps/repository"
	"enigma-laundry-apps/usecase"
	"enigma-laundry-apps/utils/exceptions"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	uomUC     usecase.UomUseCase
	productUC usecase.ProductUseCase
	engine    *gin.Engine
	host      string
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}

}

func (s *Server) initController() {

	// semua controller disini
	api.NewUomController(s.uomUC, s.engine)
	api.NewProductController(s.engine, s.productUC)

}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	exceptions.CheckErr(err)
	dbConn, _ := config.NewDbConnection(cfg)
	db := dbConn.Conn()
	uomRepo := repository.NewUomRepository(db)
	uomUseCase := usecase.NewUomUseCase(uomRepo)
	productRepo := repository.NewProductRepository(db)
	productUseCase := usecase.NewProductUseCase(productRepo, uomUseCase)

	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		uomUC:     uomUseCase,
		productUC: productUseCase,
		engine:    engine,
		host:      host,
	}
}
