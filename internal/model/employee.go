package model

import (
	"encoding/json"
	"time"
)

type Employee struct {
	Domain       string    `json:"domain"`
	Machine      string    `json:"machine"`
	User         string    `json:"user"`
	IP           string    `json:"IP"`
	LastActivity time.Time `json:"lastActivity"`
}

func (e *Employee) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Employee) Unmarshal(bytes []byte) error {
	return json.Unmarshal(bytes, e)
}
