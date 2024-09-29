package handler

import (
	"context"
	"github.com/allnightmarel0Ng/employee-controller/internal/model"
	"time"

	"github.com/allnightmarel0Ng/employee-controller/internal/app/container/usecase"
	pb "github.com/allnightmarel0Ng/employee-controller/internal/protos/container"
)

type ContainerGRPCHandler struct {
	useCase usecase.ContainerUseCase
	pb.UnimplementedContainerServer
}

func (c *ContainerGRPCHandler) Find(ctx context.Context, in *pb.TemplateRequest) (*pb.EmployeesResponse, error) {
	employee := in.GetEmployee()

	parsedTime, err := time.Parse("2023-10-05T14:48:00Z", employee.LastActivity)
	if err != nil {
		return nil, err
	}

	found, err := c.useCase.ProcessGetEmployeeByTemplate(&model.Employee{
		Host:         employee.HostName,
		User:         employee.UserName,
		IP:           employee.IP,
		LastActivity: parsedTime,
		OnDuty:       employee.OnDuty,
	})

	var result []*pb.Employee
	for _, value := range found {
		result = append(result, &pb.Employee{
			HostName:     value.Host,
			UserName:     value.User,
			IP:           value.IP,
			LastActivity: value.LastActivity.String(),
			OnDuty:       value.OnDuty,
		})
	}

	return &pb.EmployeesResponse{Employees: result}, err
}
