package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	xerrors "github.com/pkg/errors"
	"projecttest/internal/app/biz"
)

var (
	_ biz.UserRepo = new(userRepo)
	NotFoundErr = errors.New("record not found")
)

var Dataset = wire.NewSet(NewDb, NewUserRepo)

func NewUserRepo(db *sql.DB) biz.UserRepo  {
	return &userRepo{db: db}
}

func NewDb() (db *sql.DB,err error) {
	db, err = sql.Open("mysql", "root:w135790@tcp(localhost:3306)/test?charset=utf8")
	return
}

type userRepo struct {
	db *sql.DB
}

func (u *userRepo) GetUserInfo(ctx context.Context,id int32) (*biz.User, error) {
	var user biz.User
	stmt, err := u.db.Prepare("select * from testuser where id=?")
	if err != nil {
		return nil, xerrors.Wrap(err,"sql prepare failed")
	}
	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.Age, &user.Sex)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, xerrors.Wrap(NotFoundErr,fmt.Sprintf("sql query failed with err:%s",err.Error()))
		}
		return nil, xerrors.Wrap(err,"sql query failed")
	}
	return &biz.User{Id: 3,Name: "zhangsan",Age: 22,Sex: "man"},nil
}