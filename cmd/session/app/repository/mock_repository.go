package repository

import (
	"github.com/dudakp/input-server/cmd/session/app/model"
	"github.com/google/uuid"
)

var (
	sessions = make([]*model.Session, 0)
)

type MockSessionRepository struct {
}

func NewMockSessionRepository() *MockSessionRepository {
	return &MockSessionRepository{}
}

func (m *MockSessionRepository) InsertSession(session model.Session) (*model.Session, error) {
	sessions = append(sessions, &session)
	return &session, nil
}

func (m *MockSessionRepository) FindSession(id uuid.UUID) (*model.Session, error) {
	for _, session := range sessions {
		if session.Id == id {
			return session, nil
		}
	}
	return nil, SessionNotFound
}

func (m *MockSessionRepository) FindSessionByRegion(region string) ([]*model.Session, error) {
	result := make([]*model.Session, 0)
	for _, session := range sessions {
		if session.Region == region {
			result = append(result, session)
		}
	}
	return result, nil
}
