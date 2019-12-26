package basic

import (
	"outback/micro-go/basic/config"
	"outback/micro-go/basic/db"
	"outback/micro-go/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
