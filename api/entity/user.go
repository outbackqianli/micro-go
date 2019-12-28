package entity

type User struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Pwd         string `json:"pwd"`
	Token       string `json:"token"`
	CreateTime  int64  `json:"createTime"`
	UpdatedTime int64  `json:"updatedTime"`
}
