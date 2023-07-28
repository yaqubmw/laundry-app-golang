package usecase

import (
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/model/dto"
	"enigma-laundry-apps/repository"
	"fmt"
)

type CustomerUseCase interface {
	RegisterNewCustomer(payload model.Customer) error
	FindAllCustomer(requestPaging dto.PaginationParam) ([]model.Customer, dto.Paging, error)
	FindByIdCustomer(id string) (model.Customer, error)
	UpdateCustomer(payload model.Customer) error
	DeleteCustomer(id string) error
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

func (c *customerUseCase) RegisterNewCustomer(payload model.Customer) error {
	if payload.Name == "" || payload.PhoneNumber == "" || payload.Address == "" {
		return fmt.Errorf("name, phone number, and address required field")
	}

	// cek phone_number
	isExistCustomer, _ := c.repo.GetByPhone(payload.PhoneNumber)
	if isExistCustomer.PhoneNumber == payload.PhoneNumber {
		return fmt.Errorf("customer with phone number: %s already exists", payload.PhoneNumber)
	}

	err := c.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create new customer: %v", err)
	}
	return nil

}

func (c *customerUseCase) FindAllCustomer(requestPaging dto.PaginationParam) ([]model.Customer, dto.Paging, error) {
	return c.repo.Paging(requestPaging)
}

func (c *customerUseCase) FindByIdCustomer(id string) (model.Customer, error) {
	return c.repo.Get(id)
}

func (c *customerUseCase) UpdateCustomer(payload model.Customer) error {
	if payload.Name == "" || payload.PhoneNumber == "" || payload.Address == "" {
		return fmt.Errorf("name, phone number, and address required field")
	}
	isExistCustomer, _ := c.repo.GetByPhone(payload.PhoneNumber)
	if isExistCustomer.PhoneNumber == payload.PhoneNumber {
		return fmt.Errorf("customer with phone number: %s already exists", payload.PhoneNumber)
	}

	err := c.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update customer: %v", err)
	}
	return nil

}

func (c *customerUseCase) DeleteCustomer(id string) error {
	customer, err := c.FindByIdCustomer(id)
	if err != nil {
		return fmt.Errorf("data with ID %s not found", id)
	}

	err = c.repo.Delete(customer.Id)
	if err != nil {
		return fmt.Errorf("failed to delete customer: %v", err)
	}
	return nil
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{repo: repo}
}
