package infrastructure

import (
	"context"
	"flag"
	"fmt"
	"github.com/dudakp/input-server/cmd/session/app/domain"
	"github.com/dudakp/input-server/cmd/session/app/infrastructure/internal"
	"github.com/dudakp/input-server/internal/config"
	"github.com/dudakp/input-server/internal/logging"
	grpcCtx "github.com/dudakp/input-server/internal/rpc"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
)

// TODO: create integration tests

var (
	logger = logging.GetLoggerFor("session", config.IsDevelopment())
)

type typingServer struct {
	infrastructure.UnimplementedSessionServer
	sessionService *session.Service
}

func (r *typingServer) CreateSession(ctx context.Context, req *infrastructure.CreateSessionRequest) (*infrastructure.Session, error) {
	levelId, err := uuid.Parse(req.LevelId)
	if err != nil {
		logger.Error().Msgf("unable to parse levelID: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("unable to parse levelID: %v", err))
	}
	sessionID, err := r.sessionService.CreateSession(req.Name, levelId)
	if err != nil {
		logger.Error().Msgf("unable to create sessionID: %v", err)
		return nil, err
	}
	return &infrastructure.Session{
		Id: sessionID.String(),
	}, nil
}

func (r *typingServer) JoinSession(stream infrastructure.Session_JoinSessionServer) error {
	event, err := stream.Recv()
	_, uId, err := grpcCtx.GetUserFromContext(stream.Context())
	if err != nil {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("unable to get user from context: %v", err))
	}
	sessionId, err := uuid.Parse(event.GetJoin().SessionId)
	if err != nil {
		logger.Error().Msgf("unable to parse sessionID: %v", err)
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("unable to parse sessionID: %v", err))
	}
	// when the client disconnects, leave the session
	defer func(sessionService *session.Service, sessionId, playerId uuid.UUID) {
		err := sessionService.LeaveSession(sessionId, playerId)
		if err != nil {
			logger.Error().Msgf("unable to leave session: %v", err)
		}
	}(r.sessionService, sessionId, uId)

	for {
		event, err = stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			logger.Error().Msgf("unable to receive event: %v", err)
			return err
		}

		var s *session.Session
		if event.GetJoin() != nil {
			s, err = r.sessionService.JoinSession(sessionId, uId)
			if err != nil {
				logger.Error().Msgf("unable to join session: %v", err)
				return status.Errorf(codes.InvalidArgument, fmt.Sprintf("unable to join sessuib: %v", err))
			}
		} else if event.GetPing() != nil {
			s, err = r.sessionService.GetUpdates(sessionId, uId)
			if err != nil {
				return err
			}
		} else {
			return status.Errorf(codes.InvalidArgument, fmt.Sprintf("non-supported event type"))
		}

		err = stream.SendMsg(modelToRpcResponse(s))
		if err != nil {
			return err
		}
	}
}

func (r *typingServer) ListSessions(req *infrastructure.ListSessionsRequest, stream infrastructure.Session_ListSessionsServer) error {
	sessions, err := r.sessionService.FindAllSessions()
	if err != nil {
		return err
	}
	for _, session := range sessions {
		err = stream.Send(&infrastructure.Session{
			Id:         session.Id.String(),
			Name:       session.Name,
			NumPlayers: int32(len(session.Players)),
			Level: &infrastructure.Level{
				Id:         session.Level.Id.String(),
				Name:       session.Level.Name,
				Difficulty: int32(session.Level.Difficulty),
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
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
	infrastructure.RegisterSessionServer(grpcServer,
		&typingServer{
			sessionService: session.NewSessionService(infrastructure.NewMockSessionRepository()),
		})
	reflection.Register(grpcServer)
	logger.Info().Msgf("started grpc server on port: %d", grpcPort)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %w", err)
	}
}

func modelToRpcResponse(session *session.Session) *infrastructure.Session {
	return &infrastructure.Session{
		Id:         session.Id.String(),
		Name:       session.Name,
		Level:      &infrastructure.Level{Id: session.Level.Id.String(), Name: session.Level.Name, Difficulty: int32(session.Level.Difficulty)},
		NumPlayers: int32(len(session.Players)),
	}
}
