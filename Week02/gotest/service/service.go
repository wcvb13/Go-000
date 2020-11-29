package service

import (
	"gotest/internal/dao"
	"gotest/internal/model"
)

func GetUser(id int32) (*model.User,error)  {
	return dao.GetUser(id)
}
