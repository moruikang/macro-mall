// @Author moruikang
// @Date 2025/3/12 21:38:00
// @Desc

package admin

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	engine := routers.Router()

	readTimeout := 60 * time.Second
	writeTimeout := 60 * time.Second
	endPoint := fmt.Sprintf("0.0.0.0:%d", 8080)
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
