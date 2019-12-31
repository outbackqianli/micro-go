package basic

import (
	"outback/micro-go/basic/config"
	"outback/micro-go/basic/db"
	"outback/micro-go/basic/redis"

	"github.com/micro/go-micro/util/log"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
	log.SetLevel(2)
}
