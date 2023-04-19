package product

import (
	"angkringan/api/model/request"
	"angkringan/api/model/response"
	"angkringan/pkg/entity"
	"context"
)

type ServiceImpl struct {
	repository Repository
}

func NewService(repository Repository) *ServiceImpl {
	return &ServiceImpl{repository: repository}
}

func (s *ServiceImpl) Create(ctx context.Context, product request.CreateProduct) (*response.ProductResponse, error) {
	newProduct, err := s.repository.Create(ctx, entity.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	})

	return response.ToProductResponse(newProduct), err
}

func (s *ServiceImpl) Update(ctx context.Context, product request.UpdateProduct) (*response.ProductResponse, error) {
	updateProduct, err := s.repository.Create(ctx, entity.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	})

	return response.ToProductResponse(updateProduct), err
}

func (s *ServiceImpl) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func (s *ServiceImpl) FindAll(ctx context.Context, pageNumber int, pageSize int) (*[]response.ProductResponse, error) {
	products, err := s.repository.FindAll(ctx, pageNumber, pageSize)
	return response.ToProductResponses(products), err
}

func (s *ServiceImpl) FindById(ctx context.Context, id int) (*response.ProductResponse, error) {
	product, err := s.repository.FindById(ctx, id)
	return response.ToProductResponse(product), err
}

func (s *ServiceImpl) FindByName(ctx context.Context, name string) (*response.ProductResponse, error) {
	product, err := s.repository.FindByName(ctx, name)
	return response.ToProductResponse(product), err
}
