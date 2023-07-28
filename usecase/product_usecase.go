package usecase

import (
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/model/dto"
	"enigma-laundry-apps/repository"
	"fmt"
)

type ProductUseCase interface {
	RegisterNewProduct(payload model.Product) error
	FindAllProduct(requestPaging dto.PaginationParam) ([]model.Product, dto.Paging, error)
	FindByIdProduct(id string) (model.Product, error)
	UpdateProduct(payload model.Product) error
	DeleteProduct(id string) error
}

type productUseCase struct {
	repo  repository.ProductRepository
	uomUC UomUseCase
}

func (p *productUseCase) RegisterNewProduct(payload model.Product) error {
	if payload.Name == "" || payload.Price == 0 || payload.Uom.Id == "" {
		return fmt.Errorf("name, price, and uomID required field")
	}

	// cek uom ada atau tidak
	uom, err := p.uomUC.FindByIdUom(payload.Uom.Id)
	if err != nil {
		return fmt.Errorf("uom with ID %s not found", payload.Uom.Id)
	}

	payload.Uom = uom
	err = p.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to register new product: %v", err)
	}
	return nil

}

func (p *productUseCase) FindAllProduct(requestPaging dto.PaginationParam) ([]model.Product, dto.Paging, error) {
	return p.repo.Paging(requestPaging)
}

func (p *productUseCase) FindByIdProduct(id string) (model.Product, error) {
	return p.repo.Get(id)
}

func (p *productUseCase) UpdateProduct(payload model.Product) error {
	return p.repo.Update(payload)
}

func (p *productUseCase) DeleteProduct(id string) error {
	return p.repo.Delete(id)
}

func NewProductUseCase(repo repository.ProductRepository, uomUC UomUseCase) ProductUseCase {
	return &productUseCase{repo: repo, uomUC: uomUC}
}
