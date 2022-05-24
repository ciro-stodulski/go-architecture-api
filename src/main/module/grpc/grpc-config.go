package grpc

import (
	"go-api/src/main/container"
	finduserservice "go-api/src/presentation/grpc/services/user/find-user"
	"go-api/src/presentation/grpc/services/user/find-user/pb"
	"log"
)

func (server *GRPCServer) LoadServices(container *container.Container) {
	pb.RegisterFindUserServiceServer(server.Engine, finduserservice.New(container))
	log.Default().Print("gRPC: Services registered")
}
