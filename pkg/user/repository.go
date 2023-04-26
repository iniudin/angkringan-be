package user

import (
	"angkringan/pkg/entity"
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, user entity.User) (*entity.User, error)
	UpdatePassword(ctx context.Context, user entity.User) (*entity.User, error)
	UpdatePhone(ctx context.Context, user entity.User) (*entity.User, error)
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context, pageNumber int, pageSize int) (*[]entity.User, error)
	FindById(ctx context.Context, id string) (*entity.User, error)
	FindByPhone(ctx context.Context, phone string) (*entity.User, error)
}

type RepositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

// Create a new user
func (r *RepositoryImpl) Create(ctx context.Context, user entity.User) (*entity.User, error) {

	if _, err := r.db.ExecContext(
		ctx, "INSERT INTO user (id, phone, password) VALUES (?,?,?)",
		user.ID, user.Phone, user.Password,
	); err != nil {
		return nil, err
	}

	newProduct, err := r.FindById(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return newProduct, nil
}

// UpdatePassword an updating user password
func (r *RepositoryImpl) UpdatePassword(ctx context.Context, user entity.User) (*entity.User, error) {
	_, err := r.db.ExecContext(
		ctx, "UPDATE user SET password = ? WHERE id = ?",
		user.Password, user.ID,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdatePhone an updating user phone number
func (r *RepositoryImpl) UpdatePhone(ctx context.Context, user entity.User) (*entity.User, error) {
	_, err := r.db.ExecContext(
		ctx, "UPDATE user SET phone = ? WHERE id = ?",
		user.Phone, user.ID,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Delete a Product
func (r *RepositoryImpl) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM user WHERE id = ?", id)
	return err
}

// FindAll Retrieve all Product
func (r *RepositoryImpl) FindAll(ctx context.Context, pageNumber int, pageSize int) (*[]entity.User, error) {
	var users []entity.User
	offset := (pageNumber - 1) * pageSize

	result, err := r.db.QueryContext(
		ctx,
		"SELECT id, phone, password, created_at, updated_at FROM user LIMIT ? OFFSET ?",
		pageSize,
		offset,
	)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		user := entity.User{}
		if err := result.Scan(
			&user.ID,
			&user.Phone,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := result.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}

// FindById find a Product by id
func (r *RepositoryImpl) FindById(ctx context.Context, id string) (*entity.User, error) {
	user := entity.User{}
	result, err := r.db.QueryContext(
		ctx,
		"SELECT id, phone, password, created_at, updated_at FROM user WHERE id = ?",
		id,
	)
	if err != nil {
		return nil, err
	}

	if result.Next() {
		if err := result.Scan(
			&user.ID,
			&user.Phone,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}
	}

	if err := result.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByPhone find a Product by name
func (r *RepositoryImpl) FindByPhone(ctx context.Context, phone string) (*entity.User, error) {
	user := entity.User{}
	result, err := r.db.QueryContext(
		ctx,
		"SELECT id, phone, password, created_at, updated_at FROM user WHERE phone = ?",
		phone,
	)
	if err != nil {
		return nil, err
	}

	if result.Next() {
		if err := result.Scan(
			&user.ID,
			&user.Phone,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}
	}

	if err := result.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}
