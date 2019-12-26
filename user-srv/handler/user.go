package handler

import (
	"context"
	"outback/micro-go/basic/db"
	"outback/micro-go/user-srv/model"

	"github.com/go-log/log"
)

func (this *model.User) QueryUserByName(ctx context.Context, req *model.User, response *model.User) error {
	queryString := `SELECT user_id, user_name, pwd FROM user WHERE user_name = ?`
	//获取数据库
	o := db.GetDB()
	// 查询
	err := o.QueryRow(queryString, req.Name).Scan(&response.Id, &response.Name, &response.Pwd)
	if err != nil {
		log.Logf("[QueryUserByName] 查询数据失败，err：%s", err)
		return err
	}
	return nil
}