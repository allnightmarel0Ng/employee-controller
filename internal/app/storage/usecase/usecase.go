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
	if msg[:4] == "INFO" {
		tokens := strings.Split(msg[5:], " ")
		if len(tokens) != 4 {
			return errors.New("not enough information in message")
		}

		employee = &model.Employee{
			Domain:       tokens[0],
			Machine:      tokens[1],
			User:         tokens[2],
			IP:           tokens[3],
			LastActivity: time.Now(),
		}
	} else if msg[:8] == "ACTIVITY" {
		var err error
		employee, err = s.repo.GetEmployee(ctx, msg[10:])
		if err != nil {
			return err
		}
		employee.LastActivity = time.Now()
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
