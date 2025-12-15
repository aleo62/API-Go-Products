package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repository repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repository: repository}
}

func (pu ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	id, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	
	product.ID = id
	return product, nil
}

func (pu ProductUsecase) DeleteProduct(id int) (bool, error) {
	ok, err := pu.repository.DeleteProduct(id)
	if err != nil {
		return false, err
	}
	
	return ok, nil
}
