package server

import (
	"context"
	"flag"
	"fmt"
	"github.com/dudakp/input-server/cmd/session/app"
	"github.com/dudakp/input-server/cmd/session/app/model"
	"github.com/dudakp/input-server/internal/logging"
	grpcCtx "github.com/dudakp/input-server/internal/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

var (
	logger = logging.GetLoggerFor("session", false)
)

type typingServer struct {
	UnimplementedSessionServer
	sessionService app.SessionService
}

func (r *typingServer) CreateSession(ctx context.Context, req *CreateSessionRequest) (*Session, error) {
	u, err := grpcCtx.GetUserFromContext(ctx)
	if err != nil {
		logger.Error().Msgf("session can't be created - unable to get user from context: %v", err)
		return nil, err
	}
	session, err := r.sessionService.CreateSession(req, u)
	if err != nil {
		logger.Error().Msgf("unable to create session: %v", err)
		return nil, err
	}
	return &Session{
		Id:     session.Id.String(),
		Name:   session.Name,
		Users:  []*User{{Name: session.Users[0]}},
		Region: session.Region,
	}, nil
}

// TODO: handle correct closing of the stream
func (r *typingServer) JoinSession(stream Session_JoinSessionServer) error {
	event, err := stream.Recv()

	u, _ := grpcCtx.GetUserFromContext(stream.Context())
	for {
		event, err = stream.Recv()
		var session *model.Session
		if event.GetJoin() != nil {
			session, err = r.sessionService.JoinSession(event.GetJoin(), u)
			if err != nil {
				logger.Error().Msgf("unable to join session: %v", err)
				return status.Errorf(codes.InvalidArgument, fmt.Sprintf("unable to join session: %v", err))
			}
		} else if event.GetPing() != nil {
			session, err = r.sessionService.GetUpdates()
		} else {
			return status.Errorf(codes.InvalidArgument, fmt.Sprintf("non-supported event type"))
		}
		// TODO: create mapper/converter
		stream.Send(&Session{
			Id:   session.Id.String(),
			Name: session.Name,
			//Users:  []*User{{Name: session.Users[0]}},
			Region: session.Region,
		})
	}
	if err != nil {
		logger.Error().Msgf("unable to join session: %v", err)
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("unable to join session: %v", err))
	}
	return nil
}

func (r *typingServer) ListSessions(req *ListSessionsRequest, stream Session_ListSessionsServer) error {
	sessions, err := r.sessionService.FindSessionsByRegion(req.Region)
	if err != nil {
		return err
	}
	for _, session := range sessions {
		err = stream.Send(&Session{
			Id:     session.Id.String(),
			Name:   session.Name,
			Users:  r.convertUsers(session),
			Region: session.Region,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *typingServer) convertUsers(session *model.Session) []*User {
	var users []*User
	for _, user := range session.Users {
		users = append(users, &User{Name: user})
	}
	return users
}

// TODO(maybe): create common context server bootstraping function and move it to interanl package
func StartTypingServer(grpcPort int) {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterSessionServer(grpcServer, &typingServer{})
	logger.Info().Msgf("started grpc server on port: %d", grpcPort)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %w", err)
	}
}
