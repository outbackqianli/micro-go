package handler

import (
	"context"
	"outback/micro-go/auth/model"
	"outback/micro-go/auth/model/access"
	"strconv"

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

type AuthService struct{}

// MakeAccessToken 生成token
func (s *AuthService) MakeAccessToken(ctx context.Context, req *model.Request, rsp *model.Response) error {
	log.Info("[MakeAccessToken] 收到创建token请求")

	token, err := accessService.MakeAccessToken(&access.Subject{
		ID:   strconv.FormatUint(uint64(req.UserId), 10),
		Name: req.UserName,
	})
	if err != nil {
		rsp.Error = &model.Error{
			Detail: err.Error(),
		}

		log.Logf("[MakeAccessToken] token生成失败，err：%s", err)
		return err
	}

	rsp.Token = token
	return nil
}

// DelUserAccessToken 清除用户token
func (s *AuthService) DelUserAccessToken(ctx context.Context, req *model.Request, rsp *model.Response) error {
	log.Log("[DelUserAccessToken] 清除用户token")
	err := accessService.DelUserAccessToken(req.Token)
	if err != nil {
		rsp.Error = &model.Error{
			Detail: err.Error(),
		}

		log.Logf("[DelUserAccessToken] 清除用户token失败，err：%s", err)
		return err
	}

	return nil
}
