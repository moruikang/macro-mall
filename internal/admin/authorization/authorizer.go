// @Author moruikang
// @Date 2025/3/23 09:45:00
// @Desc

package authorization

import (
	"context"
	"fmt"
	"github.com/ory/ladon"
	log "github.com/sirupsen/logrus"
	"macro-mall/internal/admin/store/redis"
	"macro-mall/internal/pkg/constant"
	"strconv"
)

type Authorizer struct {
	warden ladon.Warden
}

func NewAuthorizer() *Authorizer {
	return &Authorizer{
		warden: &ladon.Ladon{
			Manager: NewRedisManager(redis.Factory, constant.MallPrefix),
		},
	}
}

func (a *Authorizer) Authorize(ctx context.Context, requests *ladon.Request) bool {
	log.Debug("authorize requests", requests)
	if err := a.warden.IsAllowed(ctx, requests); err == nil {
		return true
	}
	return false
}

func NewDefaultPolicy(roleId, resourceId int64, resourceUrl string) *ladon.DefaultPolicy {

	return &ladon.DefaultPolicy{
		ID:       fmt.Sprintf("%d:%d:%s", roleId, resourceId, resourceUrl),
		Subjects: []string{strconv.FormatInt(roleId, 10)},
		Actions:  []string{"*"},
		//Actions:     []string{"GET","POST", "PUT", "DELETE"},
		Resources: []string{resourceUrl},
		Effect:    ladon.AllowAccess,
	}
}

func NewRequests(roleId int64, resourceUrl, method string) *ladon.Request {
	return &ladon.Request{
		Resource: resourceUrl,
		Action:   method,
		Subject:  strconv.FormatInt(roleId, 10),
	}
}
