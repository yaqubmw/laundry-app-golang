package cli

import (
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/model/dto"
	"enigma-laundry-apps/usecase"
	"enigma-laundry-apps/utils/exceptions"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type EmployeeController struct {
	employeeUC usecase.EmployeeUseCase
}

func (c *EmployeeController) HandlerMainForm() {
	fmt.Println(`
| ++++ Master Employee ++++ |
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
			employee := c.createHandlerForm()
			err := c.employeeUC.RegisterNewEmployee(employee)
			exceptions.CheckErr(err)
			return
		case "2":
			requestPaging := dto.PaginationParam{
				Page: 1,
			}
			employees, paging, err := c.employeeUC.FindAllEmployee(requestPaging)
			exceptions.CheckErr(err)
			c.findAllHandlerForm(employees, paging)
			return
		case "3":
			c.getHandlerForm()
			return
		case "4":
			employee := c.updateHandlerForm()
			err := c.employeeUC.UpdateEmployee(employee)
			exceptions.CheckErr(err)
			return
		case "5":
			id := c.deleteHandlerForm()
			err := c.employeeUC.DeleteEmployee(id)
			exceptions.CheckErr(err)
			return
		case "6":
			return
		default:
			fmt.Println("Menu tidak ditemukan")
		}
	}
}

func (c *EmployeeController) createHandlerForm() model.Employee {
	var (
		employeeName, employeePhone, employeeAddress, saveConfirmation string
	)

	fmt.Print("Employee Name: ")
	fmt.Scanln(&employeeName)
	fmt.Print("Phone Number: ")
	fmt.Scanln(&employeePhone)
	fmt.Print("Employee Address: ")
	fmt.Scanln(&employeeAddress)
	fmt.Printf("Employee Name: %s, Phone Number: %s, Address: %s akan disimpan? (y/t)", employeeName, employeePhone, employeeAddress)
	fmt.Scanln(&saveConfirmation)

	if saveConfirmation == "y" {
		employee := model.Employee{
			Id:          uuid.New().String(),
			Name:        employeeName,
			PhoneNumber: employeePhone,
			Address:     employeeAddress,
		}
		return employee
	}
	return model.Employee{}
}

func (c *EmployeeController) findAllHandlerForm(employees []model.Employee, paging dto.Paging) {
	for _, employee := range employees {
		fmt.Println("Employee List")
		fmt.Printf("Employee ID: %s \n", employee.Id)
		fmt.Printf("Employee Name: %s \n", employee.Name)
		fmt.Printf("Phone Number: %s \n", employee.PhoneNumber)
		fmt.Printf("Address: %s \n", employee.Address)
		fmt.Println()
		fmt.Println("Paging: ")
		fmt.Printf("Page: %d \n", paging.Page)
		fmt.Printf("RowsPerPage: %d \n", paging.RowsPerPage)
		fmt.Printf("TotalPages: %d \n", paging.TotalPages)
		fmt.Printf("TotalRows: %d \n", paging.TotalRows)
		fmt.Println()
	}
}

func (c *EmployeeController) getHandlerForm() {
	var id string
	fmt.Print("Employee ID: ")
	fmt.Scanln(&id)
	employee, err := c.employeeUC.FindByIdEmployee(id)
	exceptions.CheckErr(err)
	fmt.Printf("Employee ID %s \n", id)
	fmt.Println(strings.Repeat("=", 15))
	fmt.Printf("Employee ID: %s \n", employee.Id)
	fmt.Printf("Employee Name: %s \n", employee.Name)
	fmt.Printf("Phone Number: %s \n", employee.PhoneNumber)
	fmt.Printf("Address: %s \n", employee.Address)
	fmt.Println()
}

func (c *EmployeeController) updateHandlerForm() model.Employee {
	var (
		employeeId, employeeName, employeePhone, employeeAddress, saveConfirmation string
	)

	fmt.Print("Employee ID: ")
	fmt.Scanln(&employeeId)
	fmt.Print("Employee Name: ")
	fmt.Scanln(&employeeName)
	fmt.Print("Phone Number: ")
	fmt.Scanln(&employeePhone)
	fmt.Print("Employee Address: ")
	fmt.Scanln(&employeeAddress)
	fmt.Printf("Employee Id: %s, Name: %s, Phone Number: %s, Address: %s akan disimpan? (y/t)", employeeId, employeeName, employeePhone, employeeAddress)
	fmt.Scanln(&saveConfirmation)

	if saveConfirmation == "y" {
		employee := model.Employee{
			Id:          employeeId,
			Name:        employeeName,
			PhoneNumber: employeePhone,
			Address:     employeeAddress,
		}
		return employee
	}
	return model.Employee{}
}

func (c *EmployeeController) deleteHandlerForm() string {
	var id string
	fmt.Print("Employee ID: ")
	fmt.Scanln(&id)
	return id
}

func NewEmployeeController(usecase usecase.EmployeeUseCase) *EmployeeController {
	return &EmployeeController{employeeUC: usecase}
}
