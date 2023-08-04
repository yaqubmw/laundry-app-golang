package manager

import "enigma-laundry-apps/usecase"

type UseCaseManager interface {
	UomUseCase() usecase.UomUseCase
	ProductUseCase() usecase.ProductUseCase
	CustomerUseCase() usecase.CustomerUseCase
	EmployeeUseCase() usecase.EmployeeUseCase
	BillUseCase() usecase.BillUseCase
	UserUseCase() usecase.UserUseCase
	AuthUseCase() usecase.AuthUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

// AuthUseCase implements UseCaseManager.
func (u *useCaseManager) AuthUseCase() usecase.AuthUseCase {
	return usecase.NewAuthUseCase(u.UserUseCase())
}

// UserUseCase implements UseCaseManager.
func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.UserRepo())
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
