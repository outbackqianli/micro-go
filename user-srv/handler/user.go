package handler

import (
	"context"
	"outback/micro-go/api/entity"
	"outback/micro-go/basic/db"
	"time"

	"github.com/go-log/log"
)

type UserHandler struct {
}

func (u *UserHandler) QueryUserByName(ctx context.Context, request string, response *entity.User) error {

	queryString := `SELECT user_id, user_name, pwd FROM user WHERE user_name = ?`
	//获取数据库
	o := db.GetDB()
	// 查询
	err := o.QueryRow(queryString, request).Scan(&response.Id, &response.Name, &response.Pwd)
	if err != nil {
		log.Logf("[QueryUserByName] 查询数据失败，err：%s", err)
		return err
	}
	time.Sleep(time.Second * 2)
	//return errors.New("执行出错")
	return nil
}

func (u *UserHandler) QueryUserByName2(ctx context.Context, request string, response *entity.User) error {

	queryString := `SELECT user_id, user_name, pwd FROM user WHERE user_name = ?`
	//获取数据库
	o := db.GetDB()
	// 查询
	err := o.QueryRow(queryString, request).Scan(&response.Id, &response.Name, &response.Pwd)
	if err != nil {
		log.Logf("[QueryUserByName] 查询数据失败，err：%s", err)
		return err
	}
	time.Sleep(time.Second * 2)
	//return errors.New("执行出错")
	return nil
}
