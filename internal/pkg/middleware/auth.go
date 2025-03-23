// @Author moruikang
// @Date 2025/3/23 19:13:00
// @Desc

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/vibrantbyte/go-antpath/antpath"
	"macro-mall/internal/admin/store/redis"
	"macro-mall/internal/pkg/constant"
	"strings"
)

func authorizator(data interface{}, g *gin.Context) {

}

func UrlMatcher(g *gin.Context) ([]string, error) {
	matcher := antpath.New()
	uri := g.Request.RequestURI
	urlMatchers := make([]string, 0)
	dataSource, err := redis.Factory.GetAllPrefixKey(constant.MallResourcePrefix)
	if err != nil {
		return nil, err
	}
	for _, url := range dataSource {
		trimUrl := strings.Split(url, constant.Delimiter)[1]
		if matcher.Match(trimUrl, uri) {
			urlMatchers = append(urlMatchers, trimUrl)
		}
	}
	return RemoveDeplicates(urlMatchers), nil
}

func RemoveDeplicates(newSlice []string) []string {
	m := make(map[string]bool)
	for _, iterm := range newSlice {
		m[iterm] = true
	}
	newSlice = make([]string, 0)
	for key := range m {
		newSlice = append(newSlice, key)
	}
	return newSlice
}
