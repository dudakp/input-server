package app

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/dudakp/input-server/cmd/session/app/grpc"
	"github.com/dudakp/input-server/cmd/session/app/model"
	"github.com/dudakp/input-server/internal/user"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type typingServer struct {
	pb.UnimplementedSessionServer
}

func (r *typingServer) CreateSession(ctx context.Context, req *pb.CreateSessionRequest) (*pb.Session, error) {
	u, _ := user.GetUserFromContext(ctx)
	session, err := CreateSession(req.Name, req.Region, u)
	if err != nil {
		return nil, err
	}
	return &pb.Session{
		Id:     session.Id.String(),
		Name:   session.Name,
		Users:  []*pb.User{{Name: u}},
		Region: session.Region,
	}, nil
}

// TODO: handle correct closing of the stream
func (r *typingServer) JoinSession(req *pb.JoinSessionRequest, stream pb.Session_JoinSessionServer) error {
	id, err := uuid.Parse(req.SessionId)
	if err != nil {
		return err
	}

	u, _ := user.GetUserFromContext(stream.Context())
	session, err := JoinSession(id, u)
	if err != nil {
		return err
	}
	users := r.convertUsers(session)

	for {
		session, err = FindSession(id)
		err := stream.Send(&pb.Session{
			Id:     session.Id.String(),
			Name:   session.Name,
			Users:  users,
			Region: session.Region,
		})
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

}

func (r *typingServer) ListSessions(req *pb.ListSessionsRequest, stream pb.Session_ListSessionsServer) error {
	sessions, err := FindSessionsByRegion(req.Region)
	if err != nil {
		return err
	}
	for _, session := range sessions {
		err = stream.Send(&pb.Session{
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

func (r *typingServer) convertUsers(session *model.Session) []*pb.User {
	var users []*pb.User
	for _, user := range session.Users {
		users = append(users, &pb.User{Name: user})
	}
	return users
}

// TODO: create common grpc server bootstraping function and move it to interanl package
func StartTypingServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSessionServer(grpcServer, &typingServer{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %w", err)
	}
}
