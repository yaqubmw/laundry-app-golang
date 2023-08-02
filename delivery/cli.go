package delivery

// import (
// 	"enigma-laundry-apps/config"
// 	"enigma-laundry-apps/delivery/controller/cli"
// 	"enigma-laundry-apps/repository"
// 	"enigma-laundry-apps/usecase"
// 	"enigma-laundry-apps/utils/exceptions"
// 	"fmt"
// 	"os"
// )

// type Console struct {
// 	// semua usecase taruh disini
// 	uomUC      usecase.UomUseCase
// 	productUC  usecase.ProductUseCase
// 	customerUC usecase.CustomerUseCase
// 	employeeUC usecase.EmployeeUseCase
// }

// func (c *Console) mainMenuForm() {
// 	fmt.Println(`
// |++++ Enigma Laundry Menu ++++|
// | 1. Master UOM               |
// | 2. Master Product           |
// | 3. Master Customer          |
// | 4. Master Employee          |
// | 5. Transaksi                |
// | 6. Keluar                   |
// 		     `)
// 	fmt.Print("Pilih Menu (1-6): ")
// }

// func (c *Console) Run() {
// 	for {
// 		c.mainMenuForm()
// 		var selectedMenu string
// 		fmt.Scanln(&selectedMenu)
// 		switch selectedMenu {
// 		case "1":
// 			cli.NewUomController(c.uomUC).UomMenuForm()
// 		case "2":
// 			cli.NewProductController(c.productUC).HandlerMainForm()
// 		case "3":
// 			cli.NewCustomerController(c.customerUC).HandlerMainForm()
// 		case "4":
// 			cli.NewEmployeeController(c.employeeUC).HandlerMainForm()
// 		case "5":
// 			fmt.Println("Transaksi")
// 		case "6":
// 			os.Exit(0)
// 		default:
// 			fmt.Println("Menu tidak ditemukan")
// 		}
// 	}
// }

// func NewConsole() *Console {
// 	cfg, err := config.NewConfig()
// 	exceptions.CheckErr(err)
// 	dbConn, _ := config.NewDbConnection(cfg)
// 	db := dbConn.Conn()
// 	uomRepo := repository.NewUomRepository(db)
// 	productRepo := repository.NewProductRepository(db)
// 	customerRepo := repository.NewCustomerRepository(db)
// 	employeeRepo := repository.NewEmployeeRepository(db)
// 	uomUseCase := usecase.NewUomUseCase(uomRepo)
// 	productUseCase := usecase.NewProductUseCase(productRepo, uomUseCase)
// 	customerUseCase := usecase.NewCustomerUseCase(customerRepo)
// 	employeeUseCase := usecase.NewEmployeeUseCase(employeeRepo)
// 	return &Console{
// 		uomUC:      uomUseCase,
// 		productUC:  productUseCase,
// 		customerUC: customerUseCase,
// 		employeeUC: employeeUseCase,
// 	}
// }
