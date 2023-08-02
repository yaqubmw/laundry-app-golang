package usecase

import (
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/model/dto"
	"enigma-laundry-apps/repository"
	"enigma-laundry-apps/utils/common"
	"fmt"
	"time"
)

type BillUseCase interface {
	RegisterNewBill(payload model.Bill) error
	FindByIdBill(id string) (dto.BillResponseDto, error)
	FindAllBill(requestPaging dto.PaginationParam) ([]dto.BillResponseDto, dto.Paging, error)
}

type billUseCase struct {
	repo       repository.BillRepository
	empUseCase EmployeeUseCase
	cstUseCase CustomerUseCase
	prdUseCase ProductUseCase
}

func (b *billUseCase) RegisterNewBill(newBill model.Bill) error {
	// get customer
	customer, err := b.cstUseCase.FindByIdCustomer(newBill.CustomerId)
	if err != nil {
		return fmt.Errorf("customer with ID %s not found", newBill.CustomerId)
	}
	// get employee
	employee, err := b.empUseCase.FindByIdEmployee(newBill.EmployeeId)
	if err != nil {
		return fmt.Errorf("employee with ID %s not found", newBill.EmployeeId)
	}
	newBillDetail := make([]model.BillDetail, 0, len(newBill.BillDetails))
	for _, detail := range newBill.BillDetails {
		// get product
		product, err := b.prdUseCase.FindByIdProduct(detail.ProductId)
		if err != nil {
			return fmt.Errorf("product with ID %s not found", newBill.EmployeeId)
		}
		detail.Id = common.GenerateID()
		detail.BillId = newBill.Id
		detail.ProductId = product.Id
		detail.ProductPrice = product.Price
		newBillDetail = append(newBillDetail, detail)
	}
	newBill.BillDate = time.Now()
	newBill.EntryDate = time.Now()
	newBill.CustomerId = customer.Id
	newBill.EmployeeId = employee.Id
	newBill.BillDetails = newBillDetail

	err = b.repo.Create(newBill)
	if err != nil {
		return fmt.Errorf("failed to register new bill %v", err)
	}

	return nil
}

func (b *billUseCase) FindAllBill(requestPaging dto.PaginationParam) ([]dto.BillResponseDto, dto.Paging, error) {
	return b.repo.List(requestPaging)
}

func (b *billUseCase) FindByIdBill(id string) (dto.BillResponseDto, error) {
	// sub total
	var subTotal int
	var billResponseDto dto.BillResponseDto
	billResponse, err := b.repo.Get(id)
	if err != nil {
		return dto.BillResponseDto{}, fmt.Errorf("failed get by id bill: %v", err.Error())
	}

	for _, item := range billResponse.BillDetails {
		subTotal += item.ProductPrice * item.Qty
	}
	billResponseDto = billResponse
	billResponseDto.TotalBill = subTotal
	return billResponseDto, nil
}

func NewBillUseCase(repo repository.BillRepository, empUseCase EmployeeUseCase, cstUseCase CustomerUseCase, prdUseCase ProductUseCase) BillUseCase {
	return &billUseCase{
		repo:       repo,
		empUseCase: empUseCase,
		cstUseCase: cstUseCase,
		prdUseCase: prdUseCase,
	}
}
