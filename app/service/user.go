package service

import (
	"GoFrame-weibo/app/dao"
	"GoFrame-weibo/app/model"
	"context"
)

// 用户管理服务
var User = &userService{

}

type userService struct {

}

func (s *userService) Create(ctx context.Context,name,email,password string) error  {
	_user := model.Users{
		Name: name,
		Email: email,
		Password: password,
	}
	 dao.Users.Ctx(ctx).Insert(&_user)
	return nil
}