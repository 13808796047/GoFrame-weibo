package service

import (
	"GoFrame-weibo/app/dao"
	"GoFrame-weibo/app/helpers"
	"GoFrame-weibo/app/model"
	"context"
	"database/sql"
	"github.com/gogf/gf/frame/g"
)

// 用户管理服务
var User = &userService{}

type userService struct {
}

func (s *userService) Create(ctx context.Context, name, email, password string) error {
	_user := model.Users{
		Name:     name,
		Email:    email,
		Password: helpers.Hash(password),
	}
	dao.Users.Ctx(ctx).Insert(&_user)
	return nil
}
func (s *userService) Show(ctx context.Context, id string) *model.Users {
	_user := model.Users{}
	err := dao.Users.Ctx(ctx).FindScan(&_user, "id", id)
	if err == nil {
		return &_user
	}
	return nil
}

func (s *userService) Update(ctx context.Context, id, name, email, introduction string) (sql.Result, error) {
	result, err := dao.Users.Ctx(ctx).Data(g.Map{"name": name, "email": email, "introduction": introduction}).Where("id", id).Update()
	if err != nil {
		return result, err
	}
	return result, nil
}
