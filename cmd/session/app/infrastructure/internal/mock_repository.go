package infrastructure

import (
	session "github.com/dudakp/input-server/cmd/session/app/domain"
	"github.com/google/uuid"
)

var (
	sessions = make([]*session.Session, 0)
)

type MockSessionRepository struct {
}

func NewMockSessionRepository() *MockSessionRepository {
	return &MockSessionRepository{}
}

func (r *MockSessionRepository) CreateSession(session session.Session) (*session.Session, error) {
	sessions = append(sessions, &session)
	return &session, nil
}

func (r *MockSessionRepository) AddPlayerToSession(sessionId uuid.UUID, playerId uuid.UUID) (*session.Session, error) {
	s, err := r.FindSession(sessionId)
	if err != nil {
		return nil, err
	}
	s.Players = append(s.Players, session.Player{
		Id: playerId,
	})
	return s, nil
}

func (r *MockSessionRepository) RemovePlayerFromSession(sessionId uuid.UUID, playerId uuid.UUID) (*session.Session, error) {
	s, err := r.FindSession(sessionId)
	if err != nil {
		return nil, err
	}
	for i, player := range s.Players {
		if player.Id == playerId {
			s.Players = append(s.Players[:i], s.Players[i+1:]...)
			return s, nil
		}
	}
	return nil, session.NotFound
}

func (r *MockSessionRepository) FindSession(id uuid.UUID) (*session.Session, error) {
	for _, s := range sessions {
		if s.Id == id {
			return s, nil
		}
	}
	return nil, session.NotFound
}

func (r *MockSessionRepository) FindAllSessions() ([]*session.Session, error) {
	return sessions, nil
}
