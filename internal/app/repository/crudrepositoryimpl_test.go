package repository

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"log"
	pb "taskRestAPI/proto"
	"testing"
)

func setupDb() *sql.DB {
	db, err:= sql.Open("postgres", "postgresql://postgres:qwe@localhost:5433/cruddb?sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
	return db
}

func TestCRUDRepositoryImpl_GetUserByUUID(t *testing.T) {
	db := setupDb()

	repo := NewCRUDRepository(db)

	user , err:=repo.GetUserByUUID(context.Background(), &pb.GetUserByUUIDReq{Uuid: "1"})
	assert.NotNil(t, user)
	assert.Nil(t, err)
}

func TestCRUDRepositoryImpl_UpdateUserByUUID(t *testing.T) {
	db := setupDb()

	repo := NewCRUDRepository(db)

	err:=repo.UpdateUserByUUID(context.Background(), &pb.UpdateUserByUUIDReq{})

	assert.NotNil(t, err)
}

func TestCRUDRepositoryImpl_CreateUser(t *testing.T) {
	db := setupDb()

	repo := NewCRUDRepository(db)

	err:=repo.CreateUser(context.Background(), &pb.CreateUserReq{})

	assert.Nil(t, err)
}
