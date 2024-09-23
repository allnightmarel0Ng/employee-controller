package model

import (
	"encoding/json"
	"time"
)

type Employee struct {
	Host         string    `json:"host"`
	User         string    `json:"user"`
	IP           string    `json:"IP"`
	LastActivity time.Time `json:"lastActivity"`
	OnDuty       bool      `json:"onDuty"`
}

func (e *Employee) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Employee) Unmarshal(bytes []byte) error {
	return json.Unmarshal(bytes, e)
}
