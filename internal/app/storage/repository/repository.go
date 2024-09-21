package repository

import (
	"context"

	"github.com/allnightmarel0Ng/employee-controller/internal/infrastructure/redis"
	"github.com/allnightmarel0Ng/employee-controller/internal/model"
)

type StorageRepository interface {
	SetEmployee(ctx context.Context, employee model.Employee) error
	GetEmployee(ctx context.Context, IP string) (*model.Employee, error)
	DeleteEmployee(ctx context.Context, IP string) error
}

type storageRepository struct {
	client *redis.Client
}

func NewStorageRepository(client *redis.Client) StorageRepository {
	return &storageRepository{
		client: client,
	}
}

func (s *storageRepository) SetEmployee(ctx context.Context, employee model.Employee) error {
	raw, err := employee.Marshal()
	if err != nil {
		return err
	}

	return s.client.Set(ctx, employee.IP, raw, 0)
}

func (s *storageRepository) GetEmployee(ctx context.Context, IP string) (*model.Employee, error) {
	raw, err := s.client.Get(ctx, IP)
	if err != nil {
		return nil, err
	}

	var employee model.Employee
	err = employee.Unmarshal([]byte(raw))
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (s *storageRepository) DeleteEmployee(ctx context.Context, IP string) error {
	_, err := s.client.Del(ctx, IP)
	return err
}
