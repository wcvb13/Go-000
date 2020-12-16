package biz

import "context"

type User struct {
	Id int32
	Name string
	Age int32
	Sex string
}

type UserRepo interface {
	GetUserInfo(ctx context.Context,id int32) (*User,error)
}

func NewUserUsercase(repo UserRepo) *UserUsecase  {
	return &UserUsecase{repo: repo}
}

type UserUsecase struct {
	repo UserRepo
}


func (u * UserUsecase) GetUser(ctx context.Context, id int32) (*User,error)  {
	return u.repo.GetUserInfo(ctx,id)
}