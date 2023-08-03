package manager

import "enigma-laundry-apps/repository"

type RepoManager interface {
	UomRepo() repository.UomRepository
	ProductRepo() repository.ProductRepository
	CustomerRepo() repository.CustomerRepository
	EmployeeRepo() repository.EmployeeRepository
	BillRepo() repository.BillRepository
	UserRepo() repository.UserRepository
}

type repoManager struct {
	infra InfraManager
}

// UserRepo implements RepoManager.
func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func (r *repoManager) UomRepo() repository.UomRepository {
	return repository.NewUomRepository(r.infra.Conn())
}

func (r *repoManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infra.Conn())
}

func (r *repoManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.Conn())
}

func (r *repoManager) EmployeeRepo() repository.EmployeeRepository {
	return repository.NewEmployeeRepository(r.infra.Conn())
}

func (r *repoManager) BillRepo() repository.BillRepository {
	return repository.NewBillRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
