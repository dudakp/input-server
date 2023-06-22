package domain

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	Id      uuid.UUID
	Name    string
	Players []Player
	Level   EmLevel
}

type EmLevel struct {
	Id         uuid.UUID
	Name       string
	Difficulty int
}

type Player struct {
	Id         uuid.UUID
	Name       string
	LastUpdate time.Time
}

type Level struct {
	Id         uuid.UUID
	Name       string
	Difficulty int
	Text       string
}
