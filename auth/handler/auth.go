package handler

import (
	"context"
	"outback/micro-go/user-srv/model"
	"strconv"

	"outback/micro-go/auth/model/access"

	"github.com/micro/go-micro/util/log"
)

var (
	accessService access.Service
)

// Init 初始化handler
func Init() {
	var err error
	accessService, err = access.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误，%s", err)
		return
	}
}

type Service struct{}

// MakeAccessToken 生成token
func (s *Service) MakeAccessToken(ctx context.Context, req *model.User, rsp *model.User) error {
	log.Log("[MakeAccessToken] 收到创建token请求,accessService", accessService)

	token, err := accessService.MakeAccessToken(&access.Subject{
		ID:   strconv.FormatInt(req.Id, 10),
		Name: req.Name,
	})
	if err != nil {

		log.Logf("[MakeAccessToken] token生成失败，err：%s", err)
		return err
	}

	rsp.Token = token
	return nil
}

// DelUserAccessToken 清除用户token
//func (s *Service) DelUserAccessToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
//	log.Log("[DelUserAccessToken] 清除用户token")
//	err := accessService.DelUserAccessToken(req.Token)
//	if err != nil {
//		rsp.Error = &auth.Error{
//			Detail: err.Error(),
//		}
//
//		log.Logf("[DelUserAccessToken] 清除用户token失败，err：%s", err)
//		return err
//	}
//
//	return nil
//}
