package service

import (
	"context"
)

// 用户管理服务
var User = &userService{

}

type userService struct {

}

func (s userService) Create(ctx context.Context)  {



	//err := g.Validator().Messages(&messages).Rules(&rules).CheckMap(&ctx)
	//if err != nil {
	//	g.Dump(err.Maps())
	//}
}