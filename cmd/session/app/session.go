package app

import (
	"github.com/dudakp/input-server/cmd/session/app/model"
	"github.com/dudakp/input-server/cmd/session/app/repository"
	pb "github.com/dudakp/input-server/cmd/session/app/server"
	"github.com/dudakp/input-server/cmd/session/app/validation"
	"github.com/dudakp/input-server/internal/config"
	"github.com/dudakp/input-server/internal/logging"
	"github.com/google/uuid"
)

/**
* TODO: get rid of dependency on grpc
* TODO: create functions for updating session users
* TODO: create functions for streaming session changes to joined user
 */

var (
	logger = logging.GetLoggerFor("session", config.IsDevelopment())
)

type SessionService struct {
	sessionRepository repository.SessionRepository
}

func (r *SessionService) CreateSession(request *pb.CreateSessionRequest, owner string) (*model.Session, error) {
	session, err := r.sessionRepository.UpsertSession(
		model.Session{
			Id:     uuid.New(),
			Name:   request.Name,
			Region: request.Region,
			Users:  []string{owner},
		})
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (r *SessionService) JoinSession(joinEvent *pb.JoinEvent, user string) (*model.Session, error) {
	logger.Info().Msgf("user %s is joining session %s", user, joinEvent.SessionId)

	validator := validation.NewJoinValidator(joinEvent.SessionId, user)
	err := validator.Validate()
	if err != nil {
		logger.Error().Msgf("validation failed: %v", err)
		return nil, validator.ToResponse()
	}

	id, _ := uuid.Parse(joinEvent.SessionId)
	return r.sessionRepository.UpsertSession(model.Session{Id: id, Users: []string{user}})
}

// TODO: add parameters
func (r *SessionService) GetUpdates() (*model.Session, error) {
	panic("implement me")
}

func (r *SessionService) FindSession(id uuid.UUID) (*model.Session, error) {
	return r.sessionRepository.FindSession(id)
}

func (r *SessionService) FindSessionsByRegion(name string) ([]*model.Session, error) {
	return r.sessionRepository.FindSessionByRegion(name)
}
