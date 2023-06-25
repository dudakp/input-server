package session

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

type Repository interface {
	CreateSession(session Session) (*Session, error)
	AddPlayerToSession(sessionId uuid.UUID, playerId uuid.UUID) (*Session, error)
	RemovePlayerFromSession(sessionId uuid.UUID, playerId uuid.UUID) (*Session, error)
	FindSession(id uuid.UUID) (*Session, error)
	FindAllSessions() ([]*Session, error)
}
