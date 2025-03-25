// @Author moruikang
// @Date 2025/3/12 21:38:00
// @Desc

package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"macro-mall/internal/admin/config"
	"macro-mall/internal/admin/router"
	"macro-mall/internal/admin/store/initialization"
	"net/http"
	"time"
)

func init() {

	config.InitConfig()
	initialization.Store()
	initialization.Redis()

}

func main() {

	engine := router.Router()

	readTimeout := 60 * time.Second
	writeTimeout := 60 * time.Second
	endPoint := fmt.Sprintf("0.0.0.0:%d", config.GlobalConfig.Server.Port)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        engine,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

}
