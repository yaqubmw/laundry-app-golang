package main

import (
	"enigma-laundry-apps/config"
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/repository"
	"enigma-laundry-apps/usecase"
	"fmt"

	_ "github.com/lib/pq"
)

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	_ "github.com/lib/pq"
// )

// type Uom struct {
// 	Id   string
// 	Name string
// }

// type Customer struct {
// 	Id          string
// 	Name        string
// 	PhoneNumber string
// 	Address     string
// }

// type Employee struct {
// 	Id          string
// 	Name        string
// 	PhoneNumber string
// 	Address     string
// }

// type Product struct {
// 	Id    string
// 	Name  string
// 	Price int
// 	Uom   Uom
// }

// type Bill struct {
// 	Id          string
// 	BillDate    time.Time
// 	EntryDate   time.Time
// 	FinishDate  time.Time
// 	EmployeeId  string
// 	CustomerId  string
// 	BillDetails []BillDetail
// }

// type BillDetail struct {
// 	Id           string
// 	BillId       string
// 	ProductId    string
// 	ProductPrice int
// 	Qty          int
// }

// func connectDB() *sql.DB {
// 	host := "localhost"
// 	port := 5432
// 	user := "postgres"
// 	password := "1234"
// 	dbname := "enigmalaundry"
// 	driver := "postgres"
// 	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	db, err := sql.Open(driver, dsn)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return db
// }

// func createUom(db *sql.DB, uom Uom) error {
// _, err := db.Exec("INSERT INTO uom (id, name) VALUES ($1, $2)", uom.Id, uom.Name)
// if err != nil {
// 	return err
// }
// fmt.Println("UOM created successfully")
// return nil

// }

// func mainMenuForm() {
// 	fmt.Println(`
// | ++++ Enigma Laundry Menu ++++ |
// | 1. Master UOM                 |
// | 2. Master Product             |
// | 3. Master Customer            |
// | 4. Master Employee            |
// | 5. Transaksi                  |
// | 6. Keluar                     |
// 	`)
// 	fmt.Print("Pilih Menu (1-6): ")
// }

// // UOM Menu Form
// func uomMenuForm() {
// 	fmt.Println(`
// | ++++ Master UOM ++++ |
// | 1. Tambah Data       |
// | 2. Lihat Data        |
// | 3. Update Data       |
// | 4. Hapus Data        |
// | 5. Kembali ke Menu   |
// 	`)
// 	fmt.Print("Pilih Menu (1-5): ")

// 	db := connectDB()
// 	defer db.Close()

// 	for {
// 		var selectedMenu string
// 		fmt.Scanln(&selectedMenu)
// 		switch selectedMenu {
// 		case "1":
// 			uom := uomCreateForm()
// 			err := createUom(db, uom)
// 			checkErr(err)
// 			return
// 		case "2":
// 			fmt.Println("Lihat Data")
// 		case "3":
// 			fmt.Println("Update Data")
// 		case "4":
// 			fmt.Println("Hapus Data")
// 		case "5":
// 			return
// 		default:
// 			fmt.Println("Menu tidak ditemukan")
// 		}
// 	}
// }

// func uomCreateForm() Uom {
// 	var (
// 		uomId, uomName, saveConfirmation string
// 	)

// 	fmt.Print("UOM ID: ")
// 	fmt.Scanln(&uomId)
// 	fmt.Print("UOM Name: ")
// 	fmt.Scanln(&uomName)
// 	fmt.Printf("UOM Id: %s, Name: %s akan disimpan? (y/t)", uomId, uomName)
// 	fmt.Scanln(&saveConfirmation)

// 	if saveConfirmation == "y" {
// 		uom := Uom{
// 			Id:   uomId,
// 			Name: uomName,
// 		}
// 		return uom
// 	}
// 	return Uom{}
// }

// func runConsole() {
// 	for {
// 		mainMenuForm()
// 		var selectedMenu string
// 		fmt.Scanln(&selectedMenu)
// 		switch selectedMenu {
// 		case "1":
// 			uomMenuForm()
// 		case "6":
// 			os.Exit(0)
// 		default:
// 			fmt.Println("Menu tidak ditemukan")
// 		}
// 	}
// }

// func checkErr(err error) {
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	dbConn, _ := config.NewDbConnection(cfg)
	db := dbConn.Conn()
	uomRepo := repository.NewUomRepository(db)
	uomUseCase := usecase.NewUomUseCase(uomRepo)

	// repository
	uom := model.Uom{
		Id:   "7",
		Name: "Butir",
	}

	err = uomUseCase.RegisterNewUom(uom)
	if err != nil {
		fmt.Println(err)
	}
}
