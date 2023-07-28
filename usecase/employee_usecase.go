package usecase

import (
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/model/dto"
	"enigma-laundry-apps/repository"
	"fmt"
)

type EmployeeUseCase interface {
	RegisterNewEmployee(payload model.Employee) error
	FindAllEmployee(requestPaging dto.PaginationParam) ([]model.Employee, dto.Paging, error)
	FindByIdEmployee(id string) (model.Employee, error)
	UpdateEmployee(payload model.Employee) error
	DeleteEmployee(id string) error
}

type employeeUseCase struct {
	repo repository.EmployeeRepository
}

func (c *employeeUseCase) RegisterNewEmployee(payload model.Employee) error {
	if payload.Name == "" || payload.PhoneNumber == "" || payload.Address == "" {
		return fmt.Errorf("name, phone number, and address required field")
	}

	// cek phone_number
	isExistEmployee, _ := c.repo.GetByPhone(payload.PhoneNumber)
	if isExistEmployee.PhoneNumber == payload.PhoneNumber {
		return fmt.Errorf("employee with phone number: %s already exists", payload.PhoneNumber)
	}

	err := c.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create new employee: %v", err)
	}
	return nil

}

func (c *employeeUseCase) FindAllEmployee(requestPaging dto.PaginationParam) ([]model.Employee, dto.Paging, error) {
	return c.repo.Paging(requestPaging)
}

func (c *employeeUseCase) FindByIdEmployee(id string) (model.Employee, error) {
	return c.repo.Get(id)
}

func (c *employeeUseCase) UpdateEmployee(payload model.Employee) error {
	if payload.Name == "" || payload.PhoneNumber == "" || payload.Address == "" {
		return fmt.Errorf("name, phone number, and address required field")
	}
	isExistEmployee, _ := c.repo.GetByPhone(payload.PhoneNumber)
	if isExistEmployee.PhoneNumber == payload.PhoneNumber {
		return fmt.Errorf("employee with phone number: %s already exists", payload.PhoneNumber)
	}

	err := c.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update employee: %v", err)
	}
	return nil

}

func (c *employeeUseCase) DeleteEmployee(id string) error {
	employee, err := c.FindByIdEmployee(id)
	if err != nil {
		return fmt.Errorf("data with ID %s not found", id)
	}

	err = c.repo.Delete(employee.Id)
	if err != nil {
		return fmt.Errorf("failed to delete employee: %v", err)
	}
	return nil
}

func NewEmployeeUseCase(repo repository.EmployeeRepository) EmployeeUseCase {
	return &employeeUseCase{repo: repo}
}
