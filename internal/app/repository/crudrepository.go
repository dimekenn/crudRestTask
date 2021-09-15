package repository

import (
	"context"
	pb "crudRestTask/proto"
)

type CRUDRepository interface {
	CreateUser(ctx context.Context, req *pb.CreateUserReq) error
	GetUserByUUID(ctx context.Context, req *pb.GetUserByUUIDReq) (*pb.GetUserByUUIDRes, error)
	UpdateUserByUUID(ctx context.Context, req *pb.UpdateUserByUUIDReq) error
}
