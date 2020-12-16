package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "projecttest/api/app/user/v1"
	"projecttest/internal/app/biz"
)

type UserService struct {
	v1.UnimplementedUserServiceServer
	usc *biz.UserUsecase
}

func NewUserService(usc *biz.UserUsecase) *UserService  {
	return &UserService{usc: usc}
}

func (u *UserService)GetUserInfo(ctx context.Context, req *v1.UserRequest) (*v1.User, error) {
	user, err := u.usc.GetUser(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound,"user not found with err: %s",err.Error())
	}
	return &v1.User{Id: user.Id,Name: user.Name,Age: user.Age,Sex: user.Sex},nil
}