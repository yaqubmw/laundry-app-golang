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
	// semua usecase di taruh disini (interface)
	uomUC      usecase.UomUseCase
	productUC  usecase.ProductUseCase
	customerUC usecase.CustomerUseCase
	employeeUC usecase.EmployeeUseCase
	billUC     usecase.BillUseCase
	engine     *gin.Engine
	host       string
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
	api.NewCustomerController(s.engine, s.customerUC)
	api.NewEmployeeController(s.engine, s.employeeUC)
	api.NewBillController(s.engine, s.billUC)
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	exceptions.CheckErr(err)
	dbConn, _ := config.NewDbConnection(cfg)
	db := dbConn.Conn()
	uomRepo := repository.NewUomRepository(db)
	productRepo := repository.NewProductRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)
	billRepo := repository.NewBillRepository(db)
	uomUseCase := usecase.NewUomUseCase(uomRepo)
	productUseCase := usecase.NewProductUseCase(productRepo, uomUseCase)
	customerUseCase := usecase.NewCustomerUseCase(customerRepo)
	employeeUseCase := usecase.NewEmployeeUseCase(employeeRepo)
	billUseCase := usecase.NewBillUseCase(billRepo, employeeUseCase, customerUseCase, productUseCase)
	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		uomUC:      uomUseCase,
		productUC:  productUseCase,
		customerUC: customerUseCase,
		employeeUC: employeeUseCase,
		billUC:     billUseCase,
		engine:     engine,
		host:       host,
	}
}
