package views

import "time"

type PayloadState struct {
	State   bool    `json:"state" validate:"required"`
	Battery float32 `json:"battery" validate:"required"`
}

type LastState struct {
	State     bool      `json:"state" validate:"required"`
	Battery   float32   `json:"battery" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type FinalState struct {
	State     bool      `json:"state" validate:"required"`
	Battery   float32   `json:"battery" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" bson:"created_at"`
}
