package session

// TODO: add validations:
//  1. session exists
//  2. PlayerId is not already in another session
// TODO: add tests?

import (
	"errors"
	"github.com/google/uuid"
)

type joinValidationError error

var (
	invalidUUID      joinValidationError = errors.New("INVALID_UUID")
	sessionNotExists joinValidationError = errors.New("SESSION_NOT_EXIST")
	userInSession    joinValidationError = errors.New("USER_IN_SESSION")
)

type JoinValidatorData struct {
	SessionId uuid.UUID
	PlayerId  uuid.UUID
}

type JoinValidator struct {
	sessionRepository Repository
	errors            error
}

func NewJoinValidator(sessionRepository Repository) *JoinValidator {
	return &JoinValidator{sessionRepository: sessionRepository}
}

func (r *JoinValidator) Validate(data JoinValidatorData) (err error) {
	r.errors = errors.Join(r.errors, r.sessionExists(data.SessionId))
	r.errors = errors.Join(r.errors, r.userNotInSession(data.SessionId, data.PlayerId))
	return
}

func (r *JoinValidator) sessionExists(id uuid.UUID) joinValidationError {
	_, err := r.sessionRepository.FindSession(id)
	if err != nil {
		return sessionNotExists
	}
	return nil
}

func (r *JoinValidator) userNotInSession(id uuid.UUID, playerId uuid.UUID) joinValidationError {
	return nil
}
