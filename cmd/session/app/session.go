package app

import (
	"github.com/dudakp/input-server/cmd/session/app/model"
	"github.com/dudakp/input-server/cmd/session/app/repository"
	"github.com/google/uuid"
)

/**
* TODO: create functions for updating session users
* TODO: create functions for streaming session changes to joined user
 */

var (
	sessionRepository repository.SessionRepository = repository.NewMockSessionRepository()
)

func CreateSession(name string, region string) (*model.Session, error) {
	session, err := sessionRepository.InsertSession(model.Session{Id: uuid.New(), Name: name, Region: region})
	if err != nil {
		return nil, err
	}
	return session, nil
}

func FindSessionsByRegion(name string) ([]*model.Session, error) {
	return sessionRepository.FindSessionByRegion(name)
}
