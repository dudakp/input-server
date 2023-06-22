package rpc

import (
	"context"
	"errors"
	"github.com/dudakp/input-server/internal/config"
	"github.com/dudakp/input-server/internal/logging"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

var (
	logger = logging.GetLoggerFor("context", config.IsDevelopment())
)

func GetUserFromContext(ctx context.Context) (string, uuid.UUID, error) {
	incomingContext, _ := metadata.FromIncomingContext(ctx)
	user := incomingContext.Get("username")
	userId := incomingContext.Get("userid")
	if user == nil || len(user) == 0 {
		return "", uuid.UUID{}, errors.New("user not found in context")
	}
	userName := user[0]
	id := userId[0]
	userUuid, err := uuid.Parse(id)
	if err != nil {
		logger.Error().Msgf("unable to parse user uuid: %v", err)
		return "", uuid.UUID{}, err
	}

	return userName, userUuid, nil
}
