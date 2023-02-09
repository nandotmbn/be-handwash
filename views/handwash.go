package views

type HandwashView struct {
	HandwashId   interface{} `json:"handwash_id,omitempty" validate:"required"`
	HandwashName string      `json:"handwash_name,omitempty" bson:"handwash_name,omitempty" validate:"required,min=0"`
}

type PayloadRetriveId struct {
	HandwashName string `json:"handwash_name,omitempty" bson:"handwash_name,omitempty" validate:"required,min=0"`
	Password     string `json:"password,omitempty" validate:"required,min=3,max=255"`
}

type FinalRetriveId struct {
	HandwashId   interface{} `json:"handwash_id,omitempty" bson:"_id,omitempty" validate:"required"`
	HandwashName string      `json:"handwash_name,omitempty" bson:"handwash_name,omitempty" validate:"required,min=0"`
}
