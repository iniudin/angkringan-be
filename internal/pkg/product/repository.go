package product

import (
	"context"
	"database/sql"
	"github.com/rs/xid"
)

type Repository interface {
	Create(ctx context.Context, product Product) (*Product, error)
	Update(ctx context.Context, product Product) (*Product, error)
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context, pageNumber int, pageSize int) (*[]Product, error)
	FindById(ctx context.Context, id string) (*Product, error)
	FindByName(ctx context.Context, name string) (*Product, error)
}

type RepositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

// Create a Product
func (r *RepositoryImpl) Create(ctx context.Context, product Product) (*Product, error) {
	guid := xid.New()

	if _, err := r.db.ExecContext(
		ctx, "INSERT INTO product (id, name, description, price) VALUES (?,?,?,?)",
		guid.String(), product.Name, product.Description, product.Price,
	); err != nil {
		return nil, err
	}

	newProduct, err := r.FindById(ctx, guid.String())
	if err != nil {
		return nil, err
	}

	return newProduct, nil
}

// Update a Product
func (r *RepositoryImpl) Update(ctx context.Context, product Product) (*Product, error) {
	_, err := r.db.ExecContext(
		ctx, "UPDATE product SET name = ?, description = ?, price = ? WHERE id = ?",
		product.Name, product.Description, product.Price, product.ID,
	)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// Delete a Product
func (r *RepositoryImpl) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM product WHERE id = ?", id)
	return err
}

// FindAll Retrieve all Product
func (r *RepositoryImpl) FindAll(ctx context.Context, pageNumber int, pageSize int) (*[]Product, error) {
	var products []Product
	offset := (pageNumber - 1) * pageSize

	result, err := r.db.QueryContext(
		ctx,
		"SELECT id, name, description, price, created_at, updated_at FROM product LIMIT ? OFFSET ?",
		pageSize,
		offset,
	)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		product := Product{}
		if err := result.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := result.Err(); err != nil {
		return nil, err
	}

	return &products, nil
}

// FindById find a Product by id
func (r *RepositoryImpl) FindById(ctx context.Context, id string) (*Product, error) {
	product := Product{}
	result, err := r.db.QueryContext(
		ctx,
		"SELECT id, name, description, price, created_at, updated_at FROM product WHERE id = ?",
		id,
	)
	if err != nil {
		return nil, err
	}

	if result.Next() {
		if err := result.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, err
		}
	}

	if err := result.Err(); err != nil {
		return nil, err
	}

	return &product, nil
}

// FindByName find a Product by name
func (r *RepositoryImpl) FindByName(ctx context.Context, name string) (*Product, error) {
	product := Product{}
	result, err := r.db.QueryContext(
		ctx,
		"SELECT id, name, description, price, created_at, updated_at FROM product WHERE name = ?",
		name,
	)
	if err != nil {
		return nil, err
	}

	if result.Next() {
		if err := result.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, err
		}
	}

	if err := result.Err(); err != nil {
		return nil, err
	}

	return &product, nil
}
