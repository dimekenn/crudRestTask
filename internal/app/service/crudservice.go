package service

import (
	"context"
	pb "crudRestTask/proto"
)

type CRUDService interface{
	CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error)
	GetUserByUUID(ctx context.Context, req *pb.GetUserByUUIDReq) (*pb.GetUserByUUIDRes, error)
	UpdateUserByUUID(ctx context.Context, req *pb.UpdateUserByUUIDReq) (*pb.UpdateUserByUUIDRes, error)
}