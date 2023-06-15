package repository

/**
 * TODO: define set of repository methods for managing session users
 */

import (
	"errors"
	"github.com/dudakp/input-server/cmd/session/app/model"
	"github.com/google/uuid"
)

var (
	SessionNotFound = errors.New("session not found")
)

type SessionRepository interface {
	UpsertSession(session model.Session) (*model.Session, error)
	FindSession(id uuid.UUID) (*model.Session, error)
	FindSessionByRegion(name string) ([]*model.Session, error)
}
