package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofrs/uuid"
)

// Artist represents an artist record.
type Artist struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
}

// Validate validates the Artist fields.
func (m Artist) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 120)),
	)
}
