package context

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
)

func GetUserFromContext(ctx context.Context) (string, error) {
	incomingContext, _ := metadata.FromIncomingContext(ctx)
	user := incomingContext.Get("username")
	if user == nil {
		return "", errors.New("context not found in context")
	}
	return user[0], nil
}
