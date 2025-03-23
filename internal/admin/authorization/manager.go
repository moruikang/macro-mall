// @Author moruikang
// @Date 2025/3/23 09:45:00
// @Desc 参考: https://github.com/ory/ladon-community/blob/master/manager/redis/manager_redis.go

package authorization

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ory/ladon"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"macro-mall/internal/admin/store"
	r "macro-mall/internal/admin/store/redis"
	"macro-mall/internal/pkg/constant"
	"strings"
	"sync"
)

var (
	ErrPolicyExists  = errors.New("Policy exists")
	ErrBadConversion = errors.New("Could not convert policy from redis")
	AuthManager      *RedisManager
	once             sync.Once
)

const (
	prefixPolicy   = constant.LadonPolicyPrefix
	prefixResource = constant.LadonResourcePrefix
	prefixSubject  = constant.LadonSubjectPrefix
)

// Just returns strings.Join(vals, "_") for creating redis keys
func PrefixKey(vals ...string) string {
	return strings.Join(vals, "_")
}

// RedisManager is a redis implementation of Manager to store policies persistently.
type RedisManager struct {
	db        r.RedisFactory
	lock      *sync.RWMutex
	keyPrefix string
}

// NewRedisManager initializes a new RedisManager with no policies
func NewRedisManager(db r.RedisFactory, keyPrefix string) *RedisManager {

	once.Do(func() {
		if keyPrefix == "" {
			keyPrefix = "ladon"
		}

		m := &RedisManager{
			db:        db,
			lock:      new(sync.RWMutex),
			keyPrefix: keyPrefix,
		}
		SetManager(m)
	})
	return AuthManager
}

func SetManager(m *RedisManager) {
	if AuthManager == nil {
		AuthManager = m
	}
}

// Create a new policy in Redis. It will create a single key for the policy itself,
// and for each subject and resource the policy will also exist in a hashmap.
func (m *RedisManager) Create(ctx context.Context, policy ladon.Policy) error {
	// Make sure that the key doesn't already exist
	key := PrefixKey(m.keyPrefix, prefixPolicy, policy.GetID())
	if err := m.db.Get(ctx, key).Err(); err == nil {
		// TODO 待优化
		log.Errorf(ErrPolicyExists.Error())
		return ErrPolicyExists
	}

	p, err := json.Marshal(policy)
	if err != nil {
		return err
	}

	// Set the policy key
	cmd := m.db.Set(ctx, key, p, 0)

	if err := cmd.Err(); err != nil {
		return err
	}

	// Put this policy in the hashmap for each resource
	for _, v := range policy.GetResources() {

		policyID := policy.GetID()
		resourceID := strings.Split(policyID, constant.Delimiter)[1]
		hmkey := fmt.Sprintf("%s:%s", PrefixKey(m.keyPrefix, prefixResource, resourceID), v)
		field := policyID
		if err := m.db.HMSet(ctx, hmkey, map[string]interface{}{
			field: p,
		}).Err(); err != nil {
			return err
		}
	}

	// Put this policy in the hashmap for each subject
	for _, v := range policy.GetSubjects() {
		hmkey := PrefixKey(m.keyPrefix, prefixSubject, v)
		field := policy.GetID()
		if err := m.db.HMSet(ctx, hmkey, map[string]interface{}{
			field: p,
		}).Err(); err != nil {
			return err
		}
	}
	return nil
}

// GetAll retrieves all policies. (Equivelant of db.keys + db.Mget)
func (m *RedisManager) GetAll(ctx context.Context, limit int64, offset int64) (ladon.Policies, error) {
	key := PrefixKey(m.keyPrefix, prefixPolicy, "*")
	/*	keyscmd := m.db.Keys(key)
		if err := keyscmd.Err(); err != nil {
			return nil, err
		}

		keys, err := keyscmd.Result()
		if err != nil {
			return nil, err
		}*/
	keys, err := m.db.GetAllPrefixKey(key)
	if err != nil {
		return nil, err
	}

	mgetcmd := m.db.MGet(ctx, keys...)
	if err := mgetcmd.Err(); err != nil {
		return nil, err
	}

	values := mgetcmd.Val()

	policies := make(ladon.Policies, len(values))
	for i, v := range values {
		p := &ladon.DefaultPolicy{}
		b := []byte(v.(string))
		// if !ok {
		// 	return nil, errors.Wrapf(ErrBadConversion, "value %+v is not a byte array", v)
		// }
		if err := json.Unmarshal(b, p); err != nil {
			return nil, errors.Wrap(ErrBadConversion, err.Error())
		}
		policies[i] = p
	}

	if offset+limit > int64(len(policies)) {
		limit = int64(len(policies))
		offset = 0
	}

	return policies[offset:limit], nil
}

// Get retrieves a policy.
func (m *RedisManager) Get(ctx context.Context, id string) (ladon.Policy, error) {
	var (
		key    = PrefixKey(m.keyPrefix, prefixPolicy, id)
		cmd    = m.db.Get(ctx, key)
		policy = &ladon.DefaultPolicy{}
	)

	if err := cmd.Err(); err != nil {
		return nil, ladon.ErrNotFound
	}
	b, err := cmd.Bytes()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, policy); err != nil {
		return nil, errors.Wrap(ErrBadConversion, err.Error())
	}
	return policy, nil
}

// Delete removes a policy.
func (m *RedisManager) Delete(ctx context.Context, id string) error {
	key := PrefixKey(m.keyPrefix, prefixPolicy, id)
	getCmd := m.db.Get(ctx, key)
	if err := getCmd.Err(); err != nil {
		return ladon.ErrNotFound
	}
	policy := &ladon.DefaultPolicy{}
	res, err := getCmd.Result()
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(res), policy); err != nil {
		return errors.Wrap(ErrBadConversion, err.Error())
	}

	if err := m.db.Del(ctx, key).Err(); err != nil {
		return err
	}

	// Put this policy in the hashmap for each resource
	for _, v := range policy.GetResources() {
		/*		hmkey := PrefixKey(m.keyPrefix, prefixResource, v)
				field := policy.GetID()*/
		policyID := policy.GetID()
		resourceID := strings.Split(policyID, constant.Delimiter)[1]
		hmkey := fmt.Sprintf("%s:%s", PrefixKey(m.keyPrefix, prefixResource, resourceID), v)
		field := policyID
		if err := m.db.HDel(ctx, hmkey, field).Err(); err != nil {
			return err
		}
	}

	// Put this policy in the hashmap for each subject
	for _, v := range policy.GetSubjects() {
		hmkey := PrefixKey(m.keyPrefix, prefixSubject, v)
		field := policy.GetID()
		if err := m.db.HDel(ctx, hmkey, field).Err(); err != nil {
			return err
		}
	}

	return nil
}

// FindPoliciesForResource returns policies that could match the resource. It either returns
// a set of policies that apply to the resource, or a superset of it.
// If an error occurs, it returns nil and the error.
func (m *RedisManager) FindPoliciesForResource(ctx context.Context, resource string) (ladon.Policies, error) {
	policies := ladon.Policies{}

	var (
		rKey    = PrefixKey(m.keyPrefix, prefixResource, resource)
		rGetCmd = m.db.HGetAll(ctx, rKey)
	)
	if err := rGetCmd.Err(); err != nil {
		return nil, err
	}

	rPolicies, err := rGetCmd.Result()
	if err != nil {
		return nil, err
	}

	for _, v := range rPolicies {
		p := &ladon.DefaultPolicy{}
		b := []byte(v)
		if err := json.Unmarshal(b, p); err != nil {
			return nil, errors.Wrap(ErrBadConversion, err.Error())
		}
		policies = append(policies, p)
	}

	return policies, nil
}

// FindPoliciesForSubject returns policies that could match the subject. It either returns
// a set of policies that applies to the subject, or a superset of it.
// If an error occurs, it returns nil and the error.
func (m *RedisManager) FindPoliciesForSubject(ctx context.Context, subject string) (ladon.Policies, error) {
	policies := ladon.Policies{}

	var (
		sKey    = PrefixKey(m.keyPrefix, prefixSubject, subject)
		sGetCmd = m.db.HGetAll(ctx, sKey)
	)
	if err := sGetCmd.Err(); err != nil {
		return nil, err
	}

	sPolicies, err := sGetCmd.Result()
	if err != nil {
		return nil, err
	}

	for _, v := range sPolicies {
		p := &ladon.DefaultPolicy{}
		b := []byte(v)
		if err := json.Unmarshal(b, p); err != nil {
			return nil, errors.Wrap(ErrBadConversion, err.Error())
		}
		policies = append(policies, p)
	}

	return policies, nil
}

// FindRequestCandidates returns candidates that could match the request object. It either returns
// a set that exactly matches the request, or a superset of it. If an error occurs, it returns nil and
// the error.
func (m *RedisManager) FindRequestCandidates(ctx context.Context, r *ladon.Request) (ladon.Policies, error) {
	policies := ladon.Policies{}
	var (
		rKey    = PrefixKey(m.keyPrefix, prefixResource, r.Resource)
		sKey    = PrefixKey(m.keyPrefix, prefixSubject, r.Subject)
		rGetCmd = m.db.HGetAll(ctx, rKey)
		sGetCmd = m.db.HGetAll(ctx, sKey)
	)
	if err := rGetCmd.Err(); err != nil {
		return nil, err
	}
	if err := sGetCmd.Err(); err != nil {
		return nil, err
	}

	rPolicies, err := rGetCmd.Result()
	if err != nil {
		return nil, err
	}
	sPolicies, err := sGetCmd.Result()
	if err != nil {
		return nil, err
	}

	for _, v := range rPolicies {
		p := &ladon.DefaultPolicy{}
		b := []byte(v)
		// if !ok {
		// 	return nil, errors.Wrapf(ErrBadConversion, "value %+v is not a byte array", v)
		// }
		if err := json.Unmarshal(b, p); err != nil {
			return nil, errors.Wrap(ErrBadConversion, err.Error())
		}
		policies = append(policies, p)
	}

	for _, v := range sPolicies {
		p := &ladon.DefaultPolicy{}
		b := []byte(v)
		// if !ok {
		// 	return nil, errors.Wrapf(ErrBadConversion, "value %+v is not a byte array", v)
		// }
		if err := json.Unmarshal(b, p); err != nil {
			return nil, errors.Wrap(ErrBadConversion, err.Error())
		}
		policies = append(policies, p)
	}

	return policies, nil
}

func (m *RedisManager) Update(ctx context.Context, policy ladon.Policy) error {
	// Make sure that the key doesn't already exist
	key := PrefixKey(m.keyPrefix, prefixPolicy, policy.GetID())
	if err := m.db.Get(ctx, key).Err(); err != nil {
		return ladon.ErrNotFound
	}

	p, err := json.Marshal(policy)
	if err != nil {
		return err
	}

	m.lock.Lock()
	defer m.lock.Unlock()
	// Set the policy key
	cmd := m.db.Set(ctx, key, p, 0)

	if err := cmd.Err(); err != nil {
		return err
	}

	// Put this policy in the hashmap for each resource
	for _, v := range policy.GetResources() {
		/*hmkey := PrefixKey(m.keyPrefix, prefixResource, v)
		field := policy.GetID()*/
		policyID := policy.GetID()
		resourceID := strings.Split(policyID, constant.Delimiter)[1]
		hmkey := fmt.Sprintf("%s:%s", PrefixKey(m.keyPrefix, prefixResource, resourceID), v)
		field := policyID
		if err := m.db.HMSet(ctx, hmkey, map[string]interface{}{
			field: p,
		}).Err(); err != nil {
			return err
		}
	}

	// Put this policy in the hashmap for each subject
	for _, v := range policy.GetSubjects() {
		hmkey := PrefixKey(m.keyPrefix, prefixSubject, v)
		field := policy.GetID()
		if err := m.db.HMSet(ctx, hmkey, map[string]interface{}{
			field: p,
		}).Err(); err != nil {
			return err
		}
	}

	return nil
}

// 以下是业务代码

func (m *RedisManager) ListAllPolicies(ctx context.Context) (ladon.Policies, error) {

	key := PrefixKey(m.keyPrefix, prefixPolicy, "*")
	keys, err := m.db.GetAllPrefixKey(key)
	if err != nil {
		return nil, err
	}
	if len(keys) == 0 {
		return nil, nil
	}

	mGetCmd := m.db.MGet(ctx, keys...)
	if err := mGetCmd.Err(); err != nil {
		return nil, err
	}
	values := mGetCmd.Val()
	policies := make(ladon.Policies, len(values))
	for i, v := range values {
		p := &ladon.DefaultPolicy{}
		b := []byte(v.(string))
		if err := json.Unmarshal(b, p); err != nil {
			return nil, errors.Wrap(ErrBadConversion, err.Error())
		}
		policies[i] = p
	}
	return policies, nil
}

/*
ClearPolicyPool
@Author: moruikang
@Description: 清除全部策略(TODO后续在看需不需要这个操作)
@receiver m
@return error
*/
func (m *RedisManager) ClearPolicyPool() error {

	ctx := context.Background()
	policies, err := m.ListAllPolicies(ctx)
	if err != nil {
		return err
	}
	for _, v := range policies {
		err := m.Delete(ctx, v.GetID())
		if err != nil {
			return err
		}
	}
	return nil
}

/*
	方案1:
	列出所有user
	获取每个user -role - resource 关系
	以用户维度来组装DefaultPolicy： 以umsAdmin id 为ladon subject、 resources url为ladon resource，组装ladon.DefaultPolicy(即一个角色可以访问那些资源)
	调用ladon.Manager 创建ladon Policy pool
	缺点: 需要维护、监听 user - role - resource 三者关联关系变化，维护复杂度大
	优点: 策略池是直接user - policy直接映射，简单明了

	方案2:
	列出所有role
	获取每个role - resource 关系
	以角色维度来组装DefaultPolicy： 以role id 为ladon subject、 resources url为ladon resource，组装ladon.DefaultPolicy(即一个用户可以访问那些资源)
	调用ladon.Manager 创建ladon Policy pool：以role_id:resource_id:resource_url 为redis key，创建一条ladon.DefaultPolicy
	缺点: 需要先通过user id获取user role列表，再通过role 列表鉴权
	优点: 仅需要维护一个role - resource 关联关系，维护简单

	最终采取方案2，RBAC
*/

func (m *RedisManager) LoadPolicyPool() error {

	ctx := context.Background()
	var err error
	allRoles, err := store.Client().UmsRoles().ListAll(ctx)
	if err != nil {
		return err
	}
	for _, role := range allRoles {
		resources, err := store.Client().UmsRoles().ListResources(ctx, role.Id)
		if err != nil {
			return err
		}
		for _, resource := range resources {
			policy := NewDefaultPolicy(role.Id, resource.Id, resource.Url)
			err := m.Create(ctx, policy)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
