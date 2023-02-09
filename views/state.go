package views

import "time"

type PayloadState struct {
	State   bool    `json:"state"`
	Battery float32 `json:"battery"`
}

type LastState struct {
	State     bool      `json:"state"`
	Battery   float32   `json:"battery"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type FinalState struct {
	State     bool      `json:"state"`
	Battery   float32   `json:"battery"`
	UpdatedAt time.Time `json:"updated_at" bson:"created_at"`
}
