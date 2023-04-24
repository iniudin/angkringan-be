package product

import (
	"angkringan/api/model/request"
	"angkringan/pkg/entity"
	"context"
)

type Service interface {
	Create(ctx context.Context, product request.CreateProduct) (*entity.Product, error)
	Update(ctx context.Context, product request.UpdateProduct) (*entity.Product, error)
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context, pageNumber int, pageSize int) (*[]entity.Product, error)
	FindById(ctx context.Context, id string) (*entity.Product, error)
	FindByName(ctx context.Context, name string) (*entity.Product, error)
}

type ServiceImpl struct {
	repository Repository
}

func NewService(repository Repository) *ServiceImpl {
	return &ServiceImpl{repository: repository}
}

func (s *ServiceImpl) Create(ctx context.Context, product request.CreateProduct) (*entity.Product, error) {
	return s.repository.Create(ctx, entity.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	})

}

func (s *ServiceImpl) Update(ctx context.Context, product request.UpdateProduct) (*entity.Product, error) {
	return s.repository.Update(ctx, entity.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	})
}

func (s *ServiceImpl) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}

func (s *ServiceImpl) FindAll(ctx context.Context, pageNumber int, pageSize int) (*[]entity.Product, error) {
	return s.repository.FindAll(ctx, pageNumber, pageSize)
}

func (s *ServiceImpl) FindById(ctx context.Context, id string) (*entity.Product, error) {
	return s.repository.FindById(ctx, id)
}

func (s *ServiceImpl) FindByName(ctx context.Context, name string) (*entity.Product, error) {
	return s.repository.FindByName(ctx, name)
}
