package dal

import (
	"app/user/biz/dal/mysql"
	"app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
