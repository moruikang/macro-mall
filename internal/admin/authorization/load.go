// @Author moruikang
// @Date 2025/3/23 09:45:00
// @Desc

package authorization

import (
	log "github.com/sirupsen/logrus"
	"macro-mall/internal/admin/store/redis"
	"macro-mall/internal/pkg/constant"
)

func InitialAuthorization() {

	var err error
	manager := NewRedisManager(redis.Factory, constant.MallPrefix)
	err = manager.ClearPolicyPool()
	if err != nil {
		log.Fatalf("Error initial ClearPolicyPool error: %v", err)
	}
	err = manager.LoadPolicyPool()
	if err != nil {
		log.Fatalf("Error initial LoadPolicyPool error: %v", err)
	}
}
