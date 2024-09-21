package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/allnightmarel0Ng/employee-controller/internal/app/storage/repository"
	"github.com/allnightmarel0Ng/employee-controller/internal/model"
)

type StorageUseCase interface {
	ProcessMessage(ctx context.Context, msg string) error
	SetEmployee(ctx context.Context, employee model.Employee) error
	GetEmployee(ctx context.Context, IP string) (*model.Employee, error)
	DeleteEmployee(ctx context.Context, IP string) error
}

type storageUseCase struct {
	repo repository.StorageRepository
}

func NewStorageUseCase(repo repository.StorageRepository) StorageUseCase {
	return &storageUseCase{
		repo: repo,
	}
}

func (s *storageUseCase) ProcessMessage(ctx context.Context, msg string) error {
	var employee *model.Employee
	msgLen := len(msg)
	if msgLen >= 4 && msg[:4] == "INFO" {
		tokens := strings.Split(msg[5:], " ")
		if len(tokens) != 3 {
			return errors.New("not enough information in message")
		}

		employee = &model.Employee{
			Host:         tokens[0],
			User:         tokens[1],
			IP:           tokens[2],
			LastActivity: time.Now(),
		}
	} else if msgLen >= 8 && msg[:8] == "ACTIVITY" {
		var err error
		employee, err = s.repo.GetEmployee(ctx, msg[10:])
		if err != nil {
			return err
		}
		employee.LastActivity = time.Now()
	} else if msgLen >= 12 && msg[:12] == "DISCONNECTED" {
		err := s.repo.DeleteEmployee(ctx, msg[13:])
		return err
	}

	if employee == nil {
		return errors.New("unable to process such message")
	}

	err := s.repo.SetEmployee(ctx, *employee)
	if err != nil {
		return err
	}

	return nil
}

func (s *storageUseCase) SetEmployee(ctx context.Context, employee model.Employee) error {
	return s.repo.SetEmployee(ctx, employee)
}

func (s *storageUseCase) GetEmployee(ctx context.Context, IP string) (*model.Employee, error) {
	return s.repo.GetEmployee(ctx, IP)
}

func (s *storageUseCase) DeleteEmployee(ctx context.Context, IP string) error {
	return s.repo.DeleteEmployee(ctx, IP)
}
