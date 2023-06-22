package infrastructure

import (
	"github.com/dudakp/input-server/cmd/session/app/domain"
	"github.com/google/uuid"
)

var (
	sessions = make([]*domain.Session, 0)
)

type MockSessionRepository struct {
}

func NewMockSessionRepository() *MockSessionRepository {
	return &MockSessionRepository{}
}

func (m *MockSessionRepository) CreateSession(session domain.Session) (*domain.Session, error) {
	sessions = append(sessions, &session)
	return &session, nil
}

func (m *MockSessionRepository) AddPlayerToSession(sessionId uuid.UUID, playerId uuid.UUID) (*domain.Session, error) {
	session, err := m.FindSession(sessionId)
	if err != nil {
		return nil, err
	}
	session.Players = append(session.Players, domain.Player{
		Id: playerId,
	})
	return session, nil
}

func (m *MockSessionRepository) RemovePlayerFromSession(sessionId uuid.UUID, playerId uuid.UUID) (*domain.Session, error) {
	session, err := m.FindSession(sessionId)
	if err != nil {
		return nil, err
	}
	for i, player := range session.Players {
		if player.Id == playerId {
			session.Players = append(session.Players[:i], session.Players[i+1:]...)
			return session, nil
		}
	}
	return nil, domain.NotFound
}

func (m *MockSessionRepository) FindSession(id uuid.UUID) (*domain.Session, error) {
	for _, session := range sessions {
		if session.Id == id {
			return session, nil
		}
	}
	return nil, domain.NotFound
}

func (m *MockSessionRepository) FindAllSessions() ([]*domain.Session, error) {
	return sessions, nil
}
