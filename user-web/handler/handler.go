package handler

import (
	"net/http"
	authClient "outback/micro-go/auth/client"
	"outback/micro-go/auth/model"
	userClient "outback/micro-go/user-web/client"

	"github.com/micro/go-micro/util/log"
)

//var (
//	serviceClient client.Client
//	authClient    client.Client
//)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

//func Init() {
//	serviceClient = client.NewClient()
//	authClient = client.NewClient()
//}

// Login 登录入口
func Login(w http.ResponseWriter, r *http.Request) {
	// 只接受POST请求
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()

	// 调用后台服务
	userName := r.Form.Get("userName")
	user, err := userClient.QueryUserByName(userName)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// 增加token验证

	if user.Pwd == r.Form.Get("pwd") {
		log.Logf("[Login] 密码校验完成，生成token...")

		// 生成token
		rsp2, err := authClient.MakeAccessToken(&model.Request{
			UserId:   user.Id,
			UserName: user.Name,
		})
		log.Info("login token is ", rsp2)

		if err != nil {
			log.Logf("[Login] 创建token失败，err：%s", err)
			http.Error(w, err.Error(), 500)
			return
		}
		//
		//	log.Logf("[Login] token %s", rsp2.Token)
		//	response["token"] = rsp2.Token
		//
		//	// 同时将token写到cookies中
		//	w.Header().Add("set-cookie", "application/json; charset=utf-8")
		//	// 过期30分钟
		//	expire := time.Now().Add(30 * time.Minute)
		//	cookie := http.Cookie{Name: "remember-me-token", Value: rsp2.Token, Path: "/", Expires: expire, MaxAge: 90000}
		//	http.SetCookie(w, &cookie)
		//}
		//
		//// 返回结果
		//response := map[string]interface{}{
		//	"ref":  time.Now().UnixNano(),
		//	"user": user,
		//}
		//
		//w.Header().Add("Content-Type", "application/json; charset=utf-8")
		//
		//// 返回JSON结构
		//if err := json.NewEncoder(w).Encode(response); err != nil {
		//	http.Error(w, err.Error(), 500)
		//	return
	}
}
