package level

import (
	"github.com/google/uuid"
)

// TODO: implement
type Repository interface {
	SaveLevel(level Level) (*Level, error)
	FindLevelById(id uuid.UUID) (*Level, error)
	FindLevelProjectionById(id uuid.UUID) (*EmLevel, error)
}
