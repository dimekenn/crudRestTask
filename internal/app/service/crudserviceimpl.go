package service

import (
	"context"
	"crudRestTask/internal/app/repository"
	pb "crudRestTask/proto"
	"github.com/micro/go-micro/v2/errors"
)


type CRUDServiceImpl struct {
	repo repository.CRUDRepository
}

//this function creates new user entity with repository
func (C CRUDServiceImpl) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	if req.FirstName == ""{
		return nil, errors.BadRequest("bad request", "fields are empty")
	}
	err := C.repo.CreateUser(ctx, req)
	if err != nil{
		return nil, err
	}
	res := &pb.CreateUserRes{}
	res.Msg = "user created"
	return res, nil
}

//GetUserByUUID takes user id and sends to CRUDRepository and returns user entity from CRUDRepository
func (C CRUDServiceImpl) GetUserByUUID(ctx context.Context, req *pb.GetUserByUUIDReq) (*pb.GetUserByUUIDRes, error) {
	return C.repo.GetUserByUUID(ctx, req)
}

//Function UpdateUserByUUID upgrades user data(username, lastname, gae, email) by UUID
func (C CRUDServiceImpl) UpdateUserByUUID(ctx context.Context, req *pb.UpdateUserByUUIDReq) (*pb.UpdateUserByUUIDRes, error) {
	err := C.repo.UpdateUserByUUID(ctx, req)
	if err != nil{
		return nil, err
	}
	res := &pb.UpdateUserByUUIDRes{}
	res.Msg = "user updated"
	return res, nil
}

//NewCRUDService creates new repository instance of CRUDService
func NewCRUDService(repo repository.CRUDRepository) CRUDService {
	return &CRUDServiceImpl{repo: repo}
}