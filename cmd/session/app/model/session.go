package model

import "github.com/google/uuid"

/**
* TODO: define final model
 */

type Session struct {
	Id     uuid.UUID
	Name   string
	Region string
	Users  []string
}
