// @Author moruikang
// @Date 2025/3/25 21:01:00
// @Desc

package initialization

import (
	log "github.com/sirupsen/logrus"
	"macro-mall/internal/admin/store/mysql"
)

func Store() {

	_, err := mysql.GetMySQLFactoryOr(nil)
	if err != nil {
		log.Fatal("failed to get mysql store factory, error: " + err.Error())
	}
}
