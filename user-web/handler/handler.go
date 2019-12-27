package handler

import (
	"fmt"
	"net/http"
	client2 "outback/micro-go/user-web/client"

	"github.com/micro/go-micro/util/log"
)

type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

// Login 登录入口
func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Errorf("login parse form error ", err.Error())
	}
	name := r.Form.Get("userName")
	prePwd := r.Form.Get("pwd")
	user, err := client2.QueryUserByName(name)
	log.Error("name ", name, prePwd)
	if err != nil {
		log.Errorf("client2.QueryUserByName error ", err.Error())
		return
	}

	if user.Pwd == prePwd {
		log.Error("密码成功配对")
		token, err := client2.GetToken(user)
		if err != nil {
			log.Errorf("client2.GetToken error ", err.Error())
			return
		}
		fmt.Printf("token is %s \n", token)
	}

	fmt.Printf("服务运行成功 user is %+v", user)

}
