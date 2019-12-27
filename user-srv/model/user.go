package model

import (
	"context"
	"fmt"
	"outback/micro-go/basic/db"

	"github.com/sirupsen/logrus"
)

type User struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Pwd         string `json:"pwd"`
	Token       string `json:"token"`
	CreateTime  int64  `json:"createTime"`
	UpdatedTime int64  `json:"updatedTime"`
}

func (this *User) QueryUserByName(ctx context.Context, req *User, response *User) error {
	fmt.Println("(this *User) QueryUserByName is there ")

	queryString := `SELECT user_id, user_name, pwd FROM user WHERE user_name = ?`
	//获取数据库
	o := db.GetDB()
	// 查询
	err := o.QueryRow(queryString, req.Name).Scan(&response.Id, &response.Name, &response.Pwd)
	if err != nil {
		logrus.Errorf("[QueryUserByName] 查询数据失败，err：%s", err)
		return err
	}
	return nil
}
