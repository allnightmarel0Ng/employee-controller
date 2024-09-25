package repository

import (
	"time"

	"github.com/allnightmarel0Ng/employee-controller/internal/infrastructure/postgres"
	"github.com/allnightmarel0Ng/employee-controller/internal/model"
)

type ContainerRepository interface {
	InsertEmployee(employee *model.Employee) error
	GetEmployeeByTemplate(template *model.Employee) ([]*model.Employee, error)
	SetWorkingStatus(IP string, status bool) error
	SetCurrentTime(IP string) error
}

type containerRepository struct {
	db *postgres.Database
}

func NewContainerRepository(db *postgres.Database) ContainerRepository {
	return &containerRepository{
		db: db,
	}
}

func (c *containerRepository) InsertEmployee(employee *model.Employee) error {
	_, err := c.db.Exec(
		"INSERT INTO public.employees (hostname, username, ip, last_activity, on_duty) VALUES ($1, $2, $3, $4, $5)",
		employee.Host, employee.User, employee.IP, employee.LastActivity, true)
	return err
}

func (c *containerRepository) GetEmployeeByTemplate(template *model.Employee) ([]*model.Employee, error) {
	query := "SELECT * FROM public.employees"
	if template.Host != "" {
		query += "WHERE hostname = " + template.Host
	}
	if template.User != "" {
		query += "AND username = " + template.User + " "
	}
	if template.IP != "" {
		query += "AND ip = " + template.IP + " "
	}
	query += ";"

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*model.Employee
	for rows.Next() {
		var hostName, userName, IP string
		var lastActivity time.Time
		var workingStatus bool
		err = rows.Scan(&hostName, &userName, &IP, &lastActivity, &workingStatus)
		if err != nil {
			return nil, err
		}

		result = append(result, &model.Employee{
			Host:         hostName,
			User:         userName,
			IP:           IP,
			LastActivity: lastActivity,
			OnDuty:       workingStatus,
		})
	}

	return result, nil
}

func (c *containerRepository) SetWorkingStatus(IP string, status bool) error {
	_, err := c.db.Exec("UPDATE public.employees SET on_duty = $1 WHERE ip = $1;", status, IP)
	return err
}

func (c *containerRepository) SetCurrentTime(IP string) error {
	_, err := c.db.Exec("UPDATE public.employees SET last_activity = $1 WHERE ip = $1;", time.Now(), IP)
	return err
}
