package model

import "time"

type Employee struct {
	Domain       string    `json:"domain"`
	Machine      string    `json:"machine"`
	User         string    `json:"user"`
	IP           string    `json:"IP"`
	LastActivity time.Time `json:"lastActivity"`
}
