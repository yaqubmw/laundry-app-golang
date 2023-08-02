package manager

import "enigma-laundry-apps/usecase"

type UseCaseManager interface {
	UomUseCase() usecase.UomUseCase
	ProductUseCase() usecase.ProductUseCase
	CustomerUseCase() usecase.CustomerUseCase
	EmployeeUseCase() usecase.EmployeeUseCase
	BillUseCase() usecase.BillUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

// BillUseCase implements UseCaseManager.
func (u *useCaseManager) BillUseCase() usecase.BillUseCase {
	return usecase.NewBillUseCase(u.repoManager.BillRepo(), u.EmployeeUseCase(), u.CustomerUseCase(), u.ProductUseCase())
}

// CustomerUseCase implements UseCaseManager.
func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repoManager.CustomerRepo())
}

// EmployeeUseCase implements UseCaseManager.
func (u *useCaseManager) EmployeeUseCase() usecase.EmployeeUseCase {
	return usecase.NewEmployeeUseCase(u.repoManager.EmployeeRepo())
}

// ProductUseCase implements UseCaseManager.
func (u *useCaseManager) ProductUseCase() usecase.ProductUseCase {
	return usecase.NewProductUseCase(u.repoManager.ProductRepo(), u.UomUseCase())
}

// UomUseCase implements UseCaseManager.
func (u *useCaseManager) UomUseCase() usecase.UomUseCase {
	return usecase.NewUomUseCase(u.repoManager.UomRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}
