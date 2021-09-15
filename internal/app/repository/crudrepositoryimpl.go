package repository

import (
	"context"
	pb "crudRestTask/proto"
	"database/sql"
	"github.com/micro/go-micro/v2/errors"
)

//CRUDRepositoryImpl implements the CRUDRepository interface
type CRUDRepositoryImpl struct {
	db *sql.DB
}

//NewCRUDRepository creates  new instance of CRUDRepository API using sql.db
func NewCRUDRepository(db *sql.DB) CRUDRepository {
	return &CRUDRepositoryImpl{db: db}
}

//CreateUser function creates new user in database with fields username, lastname, age, email
func (c CRUDRepositoryImpl) CreateUser(ctx context.Context, req *pb.CreateUserReq)  error {
	tx, err := c.db.Begin()
	if err != nil{
		return errors.InternalServerError("internalserver", err.Error())
	}

	_, err = tx.ExecContext(
		ctx,
		"insert into users (firstname, lastname, email, age) values ($1, $2, $3, $4)",
		req.FirstName, req.LastName, req.Email, req.Age,
	)
	if err != nil{
		tx.Rollback()
		return errors.InternalServerError("internalserver", err.Error())
	}
	tx.Commit()
	return nil
}

//GetUserByUUID takes id of user and returns user entity
func (c CRUDRepositoryImpl) GetUserByUUID(ctx context.Context, req *pb.GetUserByUUIDReq) (*pb.GetUserByUUIDRes, error) {
	tx, err := c.db.Begin()
	if err != nil{
		return nil, errors.InternalServerError("internalserver", err.Error())
	}

	res := &pb.GetUserByUUIDRes{}

	row := tx.QueryRowContext(
		ctx,
		"select * from users where id = $1",
		req.Uuid,
	)

	if row.Err() != nil{
		tx.Rollback()
		return nil, errors.InternalServerError("internalserver", row.Err().Error())
	}
	scanErr := row.Scan(&res.Uuid, &res.FirstName, &res.LastName, &res.Email, &res.Age, &res.CreatedDate)
	if scanErr != nil{
		tx.Rollback()
		return nil, errors.NotFound("notfound", scanErr.Error())
	}

	if res.FirstName == ""{
		tx.Rollback()
		return nil, errors.BadRequest("badrequest", "user not found")
	}

	tx.Commit()
	return res, nil
}

//UpdateUserByUUID takes all fields of user entity and updates all columns in users table by user id
func (c CRUDRepositoryImpl) UpdateUserByUUID(ctx context.Context, req *pb.UpdateUserByUUIDReq) error {
	tx, err := c.db.Begin()
	if err != nil{
		return errors.InternalServerError("internalserver", err.Error())
	}

	_, err =  tx.QueryContext(
		ctx,
		"update users set firstname = $1, lastname = $2, email = $3, age = $4 where id = $5",
		req.FirstName, req.LastName, req.Email, req.Age, req.Uuid,
	)
	if err != nil{
		tx.Rollback()
		return errors.InternalServerError("internalserver", err.Error())
	}
	tx.Commit()
	return nil
}

