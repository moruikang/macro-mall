// @Author moruikang
// @Date 2025/3/25 21:01:00
// @Desc

package initialization

import (
	log "github.com/sirupsen/logrus"
	"macro-mall/internal/admin/store/mysql"
	"macro-mall/internal/admin/store/redis"
)

func InitFactory() {

	_, err := mysql.GetMySQLFactoryOr(nil)
	if err != nil {
		log.Fatal("failed to get mysql store factory, error: " + err.Error())
	}

	_, err = redis.GetRedisFactoryOr(nil)
	if err != nil {
		log.Fatal("failed to get redis store factory, error: " + err.Error())
	}
}
