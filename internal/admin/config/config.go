// @Author moruikang
// @Date 2025/3/16 16:28:00
// @Desc

package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() {

	viper.SetConfigType("yml")
	viper.SetConfigName("admin-config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("../config/")

	viper.AutomaticEnv()

	var err error

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error read config: %s", err.Error())
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error unmarshal config: %s", err.Error())
	}

	SetGlobalConfig(&config)

}
