package session

import (
	level "github.com/dudakp/input-server/cmd/level/app/domain"
	"github.com/dudakp/input-server/internal/config"
	"github.com/dudakp/input-server/internal/logging"
	"github.com/google/uuid"
)

// TODO: create unit tests

var (
	logger = logging.GetLoggerFor("session", config.IsDevelopment())
)

type Service struct {
	sessionRepository Repository
	levelRepository   level.Repository
}

func NewSessionService(sessionRepository Repository) *Service {
	return &Service{sessionRepository: sessionRepository}
}

func (r *Service) CreateSession(name string, levelId uuid.UUID) (*uuid.UUID, error) {
	logger.Info().Msgf("creating session with name %s", name)
	l, err := r.levelRepository.FindLevelProjectionById(levelId)
	if err != nil {
		logger.Error().Msgf("unable to find level: %v", err)
		return nil, err
	}
	session, err := r.sessionRepository.CreateSession(
		Session{
			Id:      uuid.New(),
			Name:    name,
			Players: []Player{},
			Level:   *l,
		})
	if err != nil {
		logger.Error().Msgf("unable to create session: %v", err)
		return nil, err
	}
	logger.Info().Msgf("session created with id %s", session.Id)
	return &session.Id, nil
}

func (r *Service) JoinSession(sessionId, playerId uuid.UUID) (*Session, error) {
	logger.Info().Msgf("playerId %s is joining session %s", playerId, sessionId)

	validator := NewJoinValidator(r.sessionRepository)
	err := validator.Validate(JoinValidatorData{
		SessionId: sessionId,
		PlayerId:  playerId,
	})
	if err != nil {
		logger.Warn().Msgf("validation failed: %v", err)
		return nil, err
	}

	logger.Info().Msgf("playerId %s joined session %s", playerId, sessionId)
	return r.sessionRepository.AddPlayerToSession(sessionId, playerId)
}

func (r *Service) LeaveSession(sessionId, playerId uuid.UUID) error {
	logger.Info().Msgf("playerId %s is leaving session %s", playerId, sessionId)
	_, err := r.sessionRepository.RemovePlayerFromSession(sessionId, playerId)
	if err != nil {
		logger.Error().Msgf("unable to remove playerId %s from session %s: %v", playerId, sessionId, err)
		return err
	}

	if err != nil {
		logger.Error().Msgf("unable to update session: %v", err)
		return err
	}
	return nil
}

// TODO: add some validations
// TODO: this method is incomplete, think about what it needs to do
func (r *Service) GetUpdates(sessionId, playerId uuid.UUID) (*Session, error) {
	session, err := r.FindSession(sessionId)
	if err != nil {
		logger.Error().Msgf("unable to find session with id %s: %v", sessionId, err)
		return nil, err
	}
	return session, nil
}

func (r *Service) FindSession(id uuid.UUID) (*Session, error) {
	return r.sessionRepository.FindSession(id)
}

func (r *Service) FindAllSessions() ([]*Session, error) {
	return r.sessionRepository.FindAllSessions()
}
