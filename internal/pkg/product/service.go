package product

import (
	"context"
)

type Service interface {
	Create(ctx context.Context, product CreateProduct) (*Response, error)
	Update(ctx context.Context, product UpdateProduct) (*Response, error)
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context, pageNumber int, pageSize int) (*[]Response, error)
	FindById(ctx context.Context, id string) (*Response, error)
	FindByName(ctx context.Context, name string) (*Response, error)
}

type ServiceImpl struct {
	repository Repository
}

func NewService(repository Repository) *ServiceImpl {
	return &ServiceImpl{repository: repository}
}

func (s *ServiceImpl) Create(ctx context.Context, product CreateProduct) (*Product, error) {
	return s.repository.Create(ctx, Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	})
}

func (s *ServiceImpl) Update(ctx context.Context, product UpdateProduct) (*Product, error) {
	return s.repository.Update(ctx, Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	})
}

func (s *ServiceImpl) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}

func (s *ServiceImpl) FindAll(ctx context.Context, pageNumber int, pageSize int) (*[]Product, error) {
	return s.repository.FindAll(ctx, pageNumber, pageSize)
}

func (s *ServiceImpl) FindById(ctx context.Context, id string) (*Product, error) {
	return s.repository.FindById(ctx, id)
}

func (s *ServiceImpl) FindByName(ctx context.Context, name string) (*Product, error) {
	return s.repository.FindByName(ctx, name)
}
