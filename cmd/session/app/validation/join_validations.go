package validation

// TODO: add validations:
//  1. session exists
//  2. user is not already in another session
// TODO: add tests?

import (
	"errors"
	"fmt"
	"github.com/dudakp/input-server/cmd/session/app/repository"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type joinValidationError error

var (
	invalidUUID      joinValidationError = errors.New("INVALID_UUID")
	sessionNotExists joinValidationError = errors.New("SESSION_NOT_EXIST")
	userInSession    joinValidationError = errors.New("USER_IN_SESSION")
)

type JoinValidator struct {
	repository.SessionRepository
	errors error

	sessionId string
	user      string
}

func NewJoinValidator(sessionId string, user string) *JoinValidator {
	return &JoinValidator{
		sessionId: sessionId,
		user:      user,
	}
}

func (r *JoinValidator) Validate() (err error) {
	r.errors = errors.Join(r.errors, uuidValid(r.sessionId))
	if r.errors != nil {
		return r.errors
	}
	sessionUUID, _ := uuid.Parse(r.sessionId)
	r.errors = errors.Join(r.errors, sessionExists(sessionUUID))
	r.errors = errors.Join(r.errors, userNotInSession(sessionUUID, r.user))
	return
}

func (r *JoinValidator) ToResponse() error {
	return status.Errorf(codes.InvalidArgument, fmt.Sprintf("%v", r.errors))
}

func uuidValid(id string) joinValidationError {
	_, err := uuid.Parse(id)
	if err != nil {
		return invalidUUID
	}
	return nil
}

func sessionExists(id uuid.UUID) joinValidationError {
	return sessionNotExists
}

func userNotInSession(id uuid.UUID, user string) joinValidationError {
	return userInSession
}
