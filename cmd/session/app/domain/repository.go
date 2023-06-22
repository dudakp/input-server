package domain

/**
 * TODO: define set of repository methods for managing session users
 */

import (
	"errors"
	"github.com/google/uuid"
)

var (
	NotFound = errors.New("session not found")
)

type SessionRepository interface {
	CreateSession(session Session) (*Session, error)
	AddPlayerToSession(sessionId uuid.UUID, playerId uuid.UUID) (*Session, error)
	RemovePlayerFromSession(sessionId uuid.UUID, playerId uuid.UUID) (*Session, error)
	FindSession(id uuid.UUID) (*Session, error)
	FindAllSessions() ([]*Session, error)
}

// TODO: implement
type LevelRepository interface {
	SaveLevel(level Level) (*Level, error)
	FindLevelById(id uuid.UUID) (*Level, error)
	FindLevelProjectionById(id uuid.UUID) (*EmLevel, error)
}
