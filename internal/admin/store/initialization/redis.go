// @Author moruikang
// @Date 2025/3/25 21:01:00
// @Desc

package initialization

import (
	log "github.com/sirupsen/logrus"
	"macro-mall/internal/admin/store/redis"
)

func Redis() {

	_, err := redis.GetRedisFactoryOr(nil)
	if err != nil {
		log.Fatal(err)
	}
}
