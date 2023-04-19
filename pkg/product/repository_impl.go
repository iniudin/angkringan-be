package product

import (
	"angkringan/pkg/entity"
	"context"
	"database/sql"
)

type RepositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

// Create a Product
func (r *RepositoryImpl) Create(ctx context.Context, product entity.Product) (*entity.Product, error) {
	result, err := r.db.ExecContext(
		ctx,
		"INSERT INTO product(name, description, price) VALUES (?.?.?)",
		product.Name, product.Description, product.Price,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	product.ID = int(id)
	return &product, nil
}

// Update a Product
func (r *RepositoryImpl) Update(ctx context.Context, product entity.Product) (*entity.Product, error) {
	_, err := r.db.ExecContext(
		ctx,
		"UPDATE product SET name = ?, description = ?, price = ? WHERE id = ?",
		product.Name, product.Description, product.Price, product.ID,
	)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// Delete a Product
func (r *RepositoryImpl) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM product WHERE id = ?", id)
	return err
}

// FindAll Retrieve all Product
func (r *RepositoryImpl) FindAll(ctx context.Context, pageNumber int, pageSize int) ([]entity.Product, error) {
	var products []entity.Product
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

	defer func(result *sql.Rows) {
		err = result.Close()
	}(result)

	for result.Next() {
		product := entity.Product{}
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

	err = result.Close()
	if err != nil {
		return nil, err
	}

	if err := result.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// FindById find a Product by id
func (r *RepositoryImpl) FindById(ctx context.Context, id int) (*entity.Product, error) {
	product := entity.Product{}
	result, err := r.db.QueryContext(
		ctx,
		"SELECT id, name, description, price, created_at, updated_at FROM product WHERE id = ?",
		id,
	)
	if err != nil {
		return nil, err
	}

	defer func(result *sql.Rows) {
		err = result.Close()
	}(result)

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

	err = result.Close()
	if err != nil {
		return nil, err
	}

	if err := result.Err(); err != nil {
		return nil, err
	}

	return &product, nil
}

// FindByName find a Product by name
func (r *RepositoryImpl) FindByName(ctx context.Context, name string) (*entity.Product, error) {
	product := entity.Product{}
	result, err := r.db.QueryContext(
		ctx,
		"SELECT id, name, description, price, created_at, updated_at FROM product WHERE name = ?",
		name,
	)
	if err != nil {
		return nil, err
	}

	defer func(result *sql.Rows) {
		err = result.Close()
	}(result)

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

	err = result.Close()
	if err != nil {
		return nil, err
	}

	if err := result.Err(); err != nil {
		return nil, err
	}

	return &product, nil
}
