package dao

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	xerrors "github.com/pkg/errors"
	"gotest/internal/model"
)

var (
	NotFoundErr = errors.New("record not found")
	db *sql.DB
)

func init()  {
	var err error
	db, err = sql.Open("mysql", "root:w135790@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
}

func GetUser(id int32) (*model.User,error) {
	var user model.User
	stmt, err := db.Prepare("select name,age from testuser where id=?")
	if err != nil {
		return &user, xerrors.Wrap(err,"sql prepare failed")
	}
	err = stmt.QueryRow(id).Scan(&user.Name, &user.Age)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, xerrors.Wrap(NotFoundErr,"sql query failed")
		}
		return nil, xerrors.Wrap(err,"sql query failed")
	}
	return &user,nil
}