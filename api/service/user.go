package service

import (
	"context"
	"outback/micro-go/api/constent"
	"outback/micro-go/api/entity"
	"sync"

	"github.com/micro/go-micro/client"
)

type UserService interface {
	QueryUserByName(ctx context.Context, in string, opts ...client.CallOption) (*entity.User, error)
}

type userService struct {
	Clint client.Client
	Name  string
}

var (
	usersrv userService
	once    sync.Once
)

func NewUserService(c client.Client) *userService {
	once.Do(func() {
		if c == nil {
			c = client.NewClient()
		}
		usersrv = userService{
			Clint: c,
			Name:  constent.ServiceName,
		}
	})
	return &usersrv
}
