package models

import (
	"time"
)

type Handwash struct {
	HandwashName string    `json:"handwash_name,omitempty" bson:"handwash_name,omitempty" validate:"required,min=0"`
	Password     string    `json:"password,omitempty" validate:"required,min=3,max=255"`
	State        bool      `json:"state,omitempty" validate:"required,min=0,max=5"`
	Battery      float32   `json:"battery,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
