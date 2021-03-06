// default data model

package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// basic default data for a "note"

type Task struct {
	gorm.Model			// Adds some metadata fields automatically
	ID					uuid.UUID `gorm:"type:uuid"` // Explicitly specify the type to be uuid
	Title				*string
	Subtitle			string
	Text				string
	CompletedOnDate		time.Time
	// other possible types: int, int64, bool, float32, float64
	// FIELD 		DATATYPE `StructTag:"Tag Example"`
}