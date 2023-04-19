package product

import (
	"angkringan/pkg/entity"
	"context"
)

type Repository interface {
	Create(ctx context.Context, product entity.Product) (*entity.Product, error)
	Update(ctx context.Context, product entity.Product) (*entity.Product, error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context, pageNumber int, pageSize int) ([]entity.Product, error)
	FindById(ctx context.Context, id int) (*entity.Product, error)
	FindByName(ctx context.Context, name string) (*entity.Product, error)
}
