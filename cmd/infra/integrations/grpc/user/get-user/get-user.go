package getuserservice

import (
	"context"
	"fmt"
	entity "go-api/cmd/core/entities"
	user "go-api/cmd/core/entities/user"
	"go-api/cmd/infra/integrations/grpc/user/get-user/pb"
	"time"

	"google.golang.org/grpc"
)

type (
	PbGetUserService interface {
		GetUser(context.Context, *pb.NewRequestGetUser, ...grpc.CallOption) (*pb.NewResponseGetUser, error)
	}

	GetUserService interface {
		GetUser(id string) (*user.User, error)
	}

	getUserService struct {
		service PbGetUserService
	}
)

func New(service PbGetUserService) GetUserService {

	return &getUserService{
		service: service,
	}
}

func (findUser *getUserService) GetUser(id string) (*user.User, error) {
	req := &pb.NewRequestGetUser{
		ID: id,
	}

	res, err := findUser.service.GetUser(context.Background(), req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(res.Customer.CreatedAt)

	data, _ := time.Parse("2006-01-02 15:00", res.Customer.CreatedAt)
	fmt.Println(data)

	return &user.User{ID: entity.ConvertId(res.Customer.ID), Name: res.Customer.Name, Email: res.Customer.Email, CreatedAt: data}, nil
}
