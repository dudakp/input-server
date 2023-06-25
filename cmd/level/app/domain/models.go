package level

import "github.com/google/uuid"

type Level struct {
	Id         uuid.UUID
	Name       string
	Difficulty int
	Text       string
}

// EmLevel - minimal level data projection
type EmLevel struct {
	Id         uuid.UUID
	Name       string
	Difficulty int
}
