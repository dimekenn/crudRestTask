package service

import (
	"context"
	"crudRestTask/internal/app/mocks"
	pb "crudRestTask/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func setupRepo() *mocks.CRUDRepository {
	return new(mocks.CRUDRepository)
}


func TestCRUDServiceImpl_CreateUser(t *testing.T) {
	serviceMock := setupRepo()

	serviceMock.On("CreateUser", mock.Anything, mock.Anything).Return(nil)

	srv := NewCRUDService(serviceMock)

	res, err := srv.CreateUser(context.Background(), &pb.CreateUserReq{})
	assert.NotNil(t, res)
	assert.Nil(t, err)
	serviceMock.AssertExpectations(t)
}

func TestCRUDServiceImpl_GetUserByUUID(t *testing.T) {
	repoMock := setupRepo()

	repoMock.On("GetUserByUUID", mock.Anything, mock.Anything).Return(&pb.GetUserByUUIDRes{}, nil)

	srv := NewCRUDService(repoMock)
	res, err := srv.GetUserByUUID(context.Background(), &pb.GetUserByUUIDReq{})
	assert.NotNil(t, res)
	assert.Nil(t, err)
	repoMock.AssertExpectations(t)
}

func TestCRUDServiceImpl_UpdateUserByUUID(t *testing.T) {
	repoMock := setupRepo()

	repoMock.On("UpdateUserByUUID", mock.Anything, mock.Anything).Return(nil, nil)

	srv := NewCRUDService(repoMock)
	res, err := srv.UpdateUserByUUID(context.Background(), &pb.UpdateUserByUUIDReq{})
	assert.Nil(t, err)
	assert.NotNil(t, res)
	repoMock.AssertExpectations(t)
}
