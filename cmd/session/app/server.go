package app

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/dudakp/input-server/cmd/session/app/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type typingServer struct {
	pb.UnimplementedTypingServer
}

func (r *typingServer) CreateSession(ctx context.Context, req *pb.CreateSessionRequest) (*pb.Session, error) {
	session, err := CreateSession(req.Name, req.Region)
	if err != nil {
		return nil, err
	}
	return &pb.Session{
		Id:     session.Id.String(),
		Name:   session.Name,
		Users:  nil,
		Region: session.Region,
	}, nil
}

func (r *typingServer) JoinSession(req *pb.JoinSessionRequest, stream pb.Typing_JoinSessionServer) error {
	panic("implement me")
}

func (r *typingServer) ListSessions(req *pb.ListSessionsRequest, stream pb.Typing_ListSessionsServer) error {
	sessions, err := FindSessionsByRegion(req.Region)
	if err != nil {
		return err
	}
	for _, session := range sessions {
		err = stream.Send(&pb.Session{Id: session.Id.String(), Name: session.Name, Users: nil, Region: session.Region})
		if err != nil {
			return err
		}
	}
	return nil
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
	pb.RegisterTypingServer(grpcServer, &typingServer{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %w", err)
	}
}
