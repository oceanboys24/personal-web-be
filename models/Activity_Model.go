package model

import (
	"time"
)



type Activity struct {
	ID string `json:"id"`
	Action  string `json:"action"`
	Endpoint string `json:"endpoint"`
	CreatedAt time.Time `json:"created_at"`
}

type ActivityResponse struct {
	ID string `json:"id"`
	Action  string `json:"action"`
	Endpoint string `json:"endpoint"`
	CreatedAt string `json:"created_at"`
}