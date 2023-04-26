package user

import (
	"angkringan/api/model/request"
	"angkringan/pkg/entity"
	"angkringan/pkg/password"
	"context"
	"fmt"
	"github.com/rs/xid"
)

type Service interface {
	Create(ctx context.Context, request request.RegisterUser) (*entity.User, error)
	UpdatePhone(ctx context.Context, request request.UpdatePhoneUser) (*entity.User, error)
	UpdatePassword(ctx context.Context, request request.UpdatePasswordUser) (*entity.User, error)
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context, pageNumber int, pageSize int) (*[]entity.User, error)
	FindById(ctx context.Context, id string) (*entity.User, error)
	FindByPhone(ctx context.Context, name string) (*entity.User, error)
}

type ServiceImpl struct {
	repository Repository
}

func NewService(repository Repository) *ServiceImpl {
	return &ServiceImpl{repository: repository}
}

func (s *ServiceImpl) Create(ctx context.Context, request request.RegisterUser) (*entity.User, error) {
	guid := xid.New()
	hashPassword := password.HashPassword(request.Password)
	return s.repository.Create(ctx, entity.User{
		ID:       guid.String(),
		Phone:    request.Phone,
		Password: hashPassword,
	})

}

func (s *ServiceImpl) UpdatePhone(ctx context.Context, request request.UpdatePhoneUser) (*entity.User, error) {
	user, err := s.repository.FindById(ctx, request.ID)
	if err != nil {
		return nil, fmt.Errorf("user id not found")
	}
	if err := password.ComparePassword(user.Password, request.Password); err != nil {
		return nil, fmt.Errorf("user password not match")
	}

	return s.repository.UpdatePhone(ctx, entity.User{
		ID:    user.ID,
		Phone: user.Phone,
	})
}

func (s *ServiceImpl) UpdatePassword(ctx context.Context, request request.UpdatePasswordUser) (*entity.User, error) {
	user, err := s.repository.FindById(ctx, request.ID)
	if err != nil {
		return nil, fmt.Errorf("user id not found")
	}

	if err := password.ComparePassword(user.Password, request.Password); err != nil {
		return nil, fmt.Errorf("user password not match")
	}

	newPassword := password.HashPassword(request.NewPassword)

	return s.repository.UpdatePhone(ctx, entity.User{
		ID:       user.ID,
		Password: newPassword,
	})
}

func (s *ServiceImpl) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}

func (s *ServiceImpl) FindAll(ctx context.Context, pageNumber int, pageSize int) (*[]entity.User, error) {
	return s.repository.FindAll(ctx, pageNumber, pageSize)
}

func (s *ServiceImpl) FindById(ctx context.Context, id string) (*entity.User, error) {
	return s.repository.FindById(ctx, id)
}

func (s *ServiceImpl) FindByPhone(ctx context.Context, phone string) (*entity.User, error) {
	return s.repository.FindByPhone(ctx, phone)
}
