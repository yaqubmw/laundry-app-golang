package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Uom struct {
	Id   string
	Name string
}

type Customer struct {
	Id          string
	Name        string
	PhoneNumber string
	Address     string
}

type Employee struct {
	Id          string
	Name        string
	PhoneNumber string
	Address     string
}

type Product struct {
	Id    string
	Name  string
	Price int
	Uom   Uom
}

type Bill struct {
	Id          string
	BillDate    time.Time
	EntryDate   time.Time
	FinishDate  time.Time
	EmployeeId  string
	CustomerId  string
	BillDetails []BillDetail
}

type BillDetail struct {
	Id           string
	BillId       string
	ProductId    string
	ProductPrice int
	Qty          int
}

func connectDB() *sql.DB {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "1234"
	dbname := "enigmalaundry"
	driver := "postgres"
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open(driver, dsn)

	if err != nil {
		panic(err)
	}

	return db
}

func createUom(db *sql.DB, uom Uom) error {
	_, err := db.Exec("INSERT INTO uom (id, name) VALUES ($1, $2)", uom.Id, uom.Name)
	if err != nil {
		return err
	}
	fmt.Println("UOM created successfully")
	return nil

}

func main() {
	db := connectDB()
	var (
		uomId, uomName string
	)

	fmt.Print("UOM ID: ")
	fmt.Scanln(&uomId)

	fmt.Print("UOM Name: ")
	fmt.Scanln(&uomName)

	uom := Uom{
		Id:   uomId,
		Name: uomName,
	}

	err := createUom(db, uom)
	if err != nil {
		fmt.Println(err)
	}

}
