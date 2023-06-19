package rpc

import (
	"context"
	"errors"
	"github.com/dudakp/input-server/internal/config"
	"github.com/dudakp/input-server/internal/logging"
	"google.golang.org/grpc/metadata"
)

var (
	logger = logging.GetLoggerFor("context", config.IsDevelopment())
)

func GetUserFromContext(ctx context.Context) (string, error) {
	logger.Debug().Msg("getting user from context")
	incomingContext, _ := metadata.FromIncomingContext(ctx)
	user := incomingContext.Get("username")
	if user == nil {
		return "", errors.New("user not found in context")
	}
	return user[0], nil
}
