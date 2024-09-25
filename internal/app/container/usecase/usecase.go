package usecase

import (
	"errors"
	"github.com/allnightmarel0Ng/employee-controller/internal/app/container/repository"
	"github.com/allnightmarel0Ng/employee-controller/internal/model"
	"strings"
	"time"
)

type ContainerUseCase interface {
	ProcessKafkaMessage(msg string) error
}

type containerUseCase struct {
	repo repository.ContainerRepository
}

func (c *containerUseCase) ProcessKafkaMessage(msg string) error {
	if len(msg) > 4 && msg[:4] == "INFO" {
		tokens := strings.Split(msg[5:], " ")
		if len(tokens) != 3 {
			return errors.New("not enough information in message")
		}

		return c.repo.InsertEmployee(&model.Employee{
			Host:         tokens[0],
			User:         tokens[1],
			IP:           tokens[2],
			LastActivity: time.Now(),
			OnDuty:       true,
		})
	} else if len(msg) > 8 && msg[:8] == "ACTIVITY" {
		return c.repo.SetCurrentTime(msg[9:])
	} else if len(msg) > 12 && msg[:12] == "DISCONNECTED" {
		return c.repo.SetWorkingStatus(msg[13:], false)
	}

	return errors.New("unknown message type")
}

func (c *containerUseCase) ProcessGetEmployeeByTemplate(template *model.Employee) ([]*model.Employee, error) {
	return c.repo.GetEmployeeByTemplate(template)
}
