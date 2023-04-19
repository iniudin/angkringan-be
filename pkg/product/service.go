package product

import (
	"angkringan/api/model/request"
	"angkringan/api/model/response"
	"context"
)

type Service interface {
	Create(ctx context.Context, product request.CreateProduct) (*response.ProductResponse, error)
	Update(ctx context.Context, product request.UpdateProduct) (*response.ProductResponse, error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context, pageNumber int, pageSize int) (*[]response.ProductResponse, error)
	FindById(ctx context.Context, id int) (*response.ProductResponse, error)
	FindByName(ctx context.Context, name string) (*response.ProductResponse, error)
}
