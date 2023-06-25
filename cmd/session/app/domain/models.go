package session

import (
	"github.com/dudakp/input-server/cmd/level/app/domain"
	"github.com/google/uuid"
	"time"
)

type Session struct {
	Id      uuid.UUID
	Name    string
	Players []Player
	Level   level.EmLevel
}

type Player struct {
	Id         uuid.UUID
	Name       string
	LastUpdate time.Time
}
