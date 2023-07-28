package controller

import (
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/model/dto"
	"enigma-laundry-apps/usecase"
	"enigma-laundry-apps/utils/exceptions"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type CustomerController struct {
	customerUC usecase.CustomerUseCase
}

func (c *CustomerController) HandlerMainForm() {
	fmt.Println(`
| ++++ Master Customer ++++ |
| 1. Tambah Data            |
| 2. Lihat Data             |
| 3. Detail Data            |
| 4. Update Data            |
| 5. Hapus Data             |
| 6. Kembali ke Menu        |
	`)
	fmt.Print("Pilih Menu (1-6): ")

	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			customer := c.createHandlerForm()
			err := c.customerUC.RegisterNewCustomer(customer)
			exceptions.CheckErr(err)
			return
		case "2":
			requestPaging := dto.PaginationParam{
				Page: 1,
			}
			customers, paging, err := c.customerUC.FindAllCustomer(requestPaging)
			exceptions.CheckErr(err)
			c.findAllHandlerForm(customers, paging)
			return
		case "3":
			c.getHandlerForm()
			return
		case "4":
			customer := c.updateHandlerForm()
			err := c.customerUC.UpdateCustomer(customer)
			exceptions.CheckErr(err)
			return
		case "5":
			id := c.deleteHandlerForm()
			err := c.customerUC.DeleteCustomer(id)
			exceptions.CheckErr(err)
			return
		case "6":
			return
		default:
			fmt.Println("Menu tidak ditemukan")
		}
	}
}

func (c *CustomerController) createHandlerForm() model.Customer {
	var (
		customerName, customerPhone, customerAddress, saveConfirmation string
	)

	fmt.Print("Customer Name: ")
	fmt.Scanln(&customerName)
	fmt.Print("Phone Number: ")
	fmt.Scanln(&customerPhone)
	fmt.Print("Customer Address: ")
	fmt.Scanln(&customerAddress)
	fmt.Printf("Customer Name: %s, Phone Number: %s, Address: %s akan disimpan? (y/t)", customerName, customerPhone, customerAddress)
	fmt.Scanln(&saveConfirmation)

	if saveConfirmation == "y" {
		customer := model.Customer{
			Id:          uuid.New().String(),
			Name:        customerName,
			PhoneNumber: customerPhone,
			Address:     customerAddress,
		}
		return customer
	}
	return model.Customer{}
}

func (c *CustomerController) findAllHandlerForm(customers []model.Customer, paging dto.Paging) {
	for _, customer := range customers {
		fmt.Println("Customer List")
		fmt.Printf("Customer ID: %s \n", customer.Id)
		fmt.Printf("Customer Name: %s \n", customer.Name)
		fmt.Printf("Phone Number: %s \n", customer.PhoneNumber)
		fmt.Printf("Address: %s \n", customer.Address)
		fmt.Println()
		fmt.Println("Paging: ")
		fmt.Printf("Page: %d \n", paging.Page)
		fmt.Printf("RowsPerPage: %d \n", paging.RowsPerPage)
		fmt.Printf("TotalPages: %d \n", paging.TotalPages)
		fmt.Printf("TotalRows: %d \n", paging.TotalRows)
		fmt.Println()
	}
}

func (c *CustomerController) getHandlerForm() {
	var id string
	fmt.Print("Customer ID: ")
	fmt.Scanln(&id)
	customer, err := c.customerUC.FindByIdCustomer(id)
	exceptions.CheckErr(err)
	fmt.Printf("Customer ID %s \n", id)
	fmt.Println(strings.Repeat("=", 15))
	fmt.Printf("Customer ID: %s \n", customer.Id)
	fmt.Printf("Customer Name: %s \n", customer.Name)
	fmt.Printf("Phone Number: %s \n", customer.PhoneNumber)
	fmt.Printf("Address: %s \n", customer.Address)
	fmt.Println()
}

func (c *CustomerController) updateHandlerForm() model.Customer {
	var (
		customerId, customerName, customerPhone, customerAddress, saveConfirmation string
	)

	fmt.Print("Customer ID: ")
	fmt.Scanln(&customerId)
	fmt.Print("Customer Name: ")
	fmt.Scanln(&customerName)
	fmt.Print("Phone Number: ")
	fmt.Scanln(&customerPhone)
	fmt.Print("Customer Address: ")
	fmt.Scanln(&customerAddress)
	fmt.Printf("Customer Id: %s, Name: %s, Phone Number: %s, Address: %s akan disimpan? (y/t)", customerId, customerName, customerPhone, customerAddress)
	fmt.Scanln(&saveConfirmation)

	if saveConfirmation == "y" {
		customer := model.Customer{
			Id:          customerId,
			Name:        customerName,
			PhoneNumber: customerPhone,
			Address:     customerAddress,
		}
		return customer
	}
	return model.Customer{}
}

func (c *CustomerController) deleteHandlerForm() string {
	var id string
	fmt.Print("Customer ID: ")
	fmt.Scanln(&id)
	return id
}

func NewCustomerController(usecase usecase.CustomerUseCase) *CustomerController {
	return &CustomerController{customerUC: usecase}
}
