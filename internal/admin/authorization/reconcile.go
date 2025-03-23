// @Author moruikang
// @Date 2025/3/23 09:45:00
// @Desc

package authorization

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ory/ladon"
	"github.com/pkg/errors"
	"macro-mall/internal/admin/store"
	"macro-mall/internal/admin/store/models"
	"strconv"
)

func (m *RedisManager) UpdateResourceRelaPolicy(ctx context.Context, resource *models.UmsResource) error {

	// 查询与resource关联的policy,逐一更新
	hmkey := fmt.Sprintf("%s:%s", PrefixKey(m.keyPrefix, prefixResource, strconv.FormatInt(resource.Id, 10)))
	allHmKey, err := m.db.HGetAllKey(hmkey)
	if err != nil {
		return err
	}

	for _, policyStr := range allHmKey {
		var policy *ladon.DefaultPolicy
		if err := json.Unmarshal([]byte(policyStr), &policy); err != nil {
			return errors.Wrap(ErrBadConversion, err.Error())
		}
		policy.Resources = []string{resource.Url}
		if err := m.Update(ctx, policy); err != nil {
			return err
		}
	}
	return nil
}

func (m *RedisManager) DeleteResourceRelaPolicy(ctx context.Context, resourceId int64) error {

	resource, err := store.Client().UmsResources().GetById(ctx, resourceId)
	if err != nil {
		return err
	}

	// 查询与resource关联的policy,逐一更新
	hmkey := fmt.Sprintf("%s:%s", PrefixKey(m.keyPrefix, prefixResource, strconv.FormatInt(resource.Id, 10)))
	allHmKey, err := m.db.HGetAllKey(hmkey)
	if err != nil {
		return err
	}

	for policyId, _ := range allHmKey {
		if err := m.Delete(ctx, policyId); err != nil {
			return err
		}
	}
	return nil
}

func (m *RedisManager) CreateRoleResourceRelaPolicy(ctx context.Context, rrr []*models.UmsRoleResourceRelation) error {

	for _, r := range rrr {
		resource, err := store.Client().UmsResources().GetById(ctx, r.ResourceId)
		if err != nil {
			return err
		}
		policy := NewDefaultPolicy(r.RoleId, r.ResourceId, resource.Url)
		err = m.Create(ctx, policy)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *RedisManager) DeleteRoleResourceRelaPolicy(ctx context.Context, rrr []*models.UmsRoleResourceRelation) error {

	for _, r := range rrr {
		resource, err := store.Client().UmsResources().GetById(ctx, r.ResourceId)
		if err != nil {
			return err
		}
		policyId := fmt.Sprintf("%d:%d:%s", r.RoleId, r.ResourceId, resource.Url)
		err = m.Delete(ctx, policyId)
		if err != nil {
			return err
		}
	}
	return nil
}
