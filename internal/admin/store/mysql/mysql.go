// @Author moruikang
// @Date 2025/3/16 10:45:00
// @Desc

package mysql

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"macro-mall/internal/admin/config"
	"macro-mall/internal/admin/store"
	"macro-mall/internal/pkg/options"
	"sync"
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) UmsAdmins() store.UmsAdminStore {
	return newUmsAdmin(ds)
}

func (ds *datastore) UmsRoles() store.UmsRoleStore {
	return newUmsRole(ds)
}

func (ds *datastore) UmsMenus() store.UmsMenuStore {
	return newUmsMenu(ds)
}

func (ds *datastore) UmsResources() store.UmsResourceStore {
	return newUmsResource(ds)
}

func (ds *datastore) UmsResourceCategorys() store.UmsResourceCategoryStore {
	return newUmsResourceCategory(ds)
}

// 确保实现了Factory工厂接口
var _ store.Factory = (*datastore)(nil)

var (
	mysqlFactory store.Factory
	once         sync.Once
)

func GetMySQLFactoryOr(opts *options.MySQLOptions) (store.Factory, error) {
	if opts == nil && mysqlFactory == nil {
		opts = &config.GlobalConfig.Mysql
	}

	var err error
	var dbIns *gorm.DB
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			opts.Username,
			opts.Password,
			opts.Host,
			opts.Database)

		dbIns, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("MySQL Connected Fail, err: %s", err.Error())
		}
		sqlDB, err := dbIns.DB()
		if err != nil {
			log.Fatalf("MySQL Connected DB Fail, err: %s", err.Error())
		}

		sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)
		sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)
		sqlDB.SetConnMaxIdleTime(opts.MaxConnectionLifeTime)

		mysqlFactory = &datastore{dbIns}
		store.SetClient(mysqlFactory)
	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}

	return mysqlFactory, nil
}
