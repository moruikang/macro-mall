// @Author moruikang
// @Date 2025/3/16 16:28:00
// @Desc

package config

import "macro-mall/internal/pkg/options"

var (
	GlobalConfig *Config
)

type Config struct {
	Server Server
	Jwt    Jwt
	Mysql  options.MySQLOptions
	Redis  options.RedisOptions
}

type Server struct {
	Bind string `json:"bind" mapstructure:"bind"`
	Port int    `json:"port" mapstructure:"port"`
}

type Jwt struct {
	TokenHeader string `json:"tokenHeader" mapstructure:"tokenHeader"`
	TokenHead   string `json:"tokenHead" mapstructure:"tokenHead"`
	RealmName   string `json:"realmName" mapstructure:"realmName"`
	Key         string `json:"key" mapstructure:"key"`
	Timeout     int64  `json:"timeout" mapstructure:"timeout"`
	MaxRefresh  int64  `json:"maxRefresh" mapstructure:"maxRefresh"`
}

func SetGlobalConfig(c *Config) {
	GlobalConfig = c
}
