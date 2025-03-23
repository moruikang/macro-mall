// @Author moruikang
// @Date 2025/3/23 19:13:00
// @Desc

package middleware

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"macro-mall/internal/admin/config"
	"macro-mall/internal/admin/dto"
	"macro-mall/internal/admin/store"
	"macro-mall/internal/admin/store/models"
	"macro-mall/internal/pkg/common/bcryptx"
	"macro-mall/pkg/response"
	"net/http"
	"strings"
	"time"
)

const username = "username"

type UmsAdminTokenInfo struct {
	Username string `json:"username"`
	UserId   int64  `json:"userId"`
}

type CustomClaims struct {
	jwt.MapClaims
	Username string  `json:"username"`
	UserId   int64   `json:"userId"`
	RoleIds  []int64 `json:"roleIds"`
}

type UserInfo struct {
	user  *models.UmsAdmin
	roles []*models.UmsRole
}

func SetJwtInfo() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.Unauthorized(c)
			c.Abort()
			return
		}
		info := strings.Split(token, ".")[1]
		payload, _ := base64.RawURLEncoding.DecodeString(info)

		var user UmsAdminTokenInfo
		_ = json.Unmarshal(payload, &user)
		c.Request.Header.Set("Username", user.Username)
		c.Request.Header.Set("UserId", fmt.Sprintf("%d", user.UserId))
		c.Next()
	}
}

func payloadFunc() func(data interface{}) jwt.MapClaims {

	return func(data interface{}) jwt.MapClaims {
		claims := jwt.MapClaims{}
		if v, ok := data.(*UserInfo); ok {
			claims[username] = v.user.Username
			claims["userId"] = v.user.Id
			roleIds := make([]int64, 0)
			for _, role := range v.roles {
				roleIds = append(roleIds, role.Id)
			}
			claims["roleIds"] = roleIds
		}
		return claims
	}
}

func authenticator() func(c *gin.Context) (interface{}, error) {

	return func(c *gin.Context) (interface{}, error) {

		var loginDTO *dto.AdminLoginDTO
		if err := c.ShouldBindJSON(loginDTO); err != nil {
			return nil, jwt.ErrMissingLoginValues
		}
		// 校验用户信息
		umsAdmin, err := store.Client().UmsAdmins().GetByUserName(c, loginDTO.Username)
		if err != nil {
			return nil, jwt.ErrFailedAuthentication
		}
		if !bcryptx.VerifyPassword(loginDTO.Password, umsAdmin.Password) {
			return nil, jwt.ErrFailedAuthentication
		}
		now := time.Now()
		if err = store.Client().UmsAdmins().UpdateLoginTime(c, umsAdmin.Id, &now); err != nil {
			return nil, err
		}
		userRoles, err := store.Client().UmsAdmins().GetUserRoles(c, umsAdmin.Id)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Fail to get user's roles, Error=%s", err.Error()))
		}
		return &UserInfo{
			user:  umsAdmin,
			roles: userRoles,
		}, nil
	}
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	uid := int64(claims["userId"].(float64))
	roleIDs := make([]int64, 0)
	roleIds := claims["roleIds"].([]interface{})
	if len(roleIds) > 0 {
		for _, roleId := range roleIds {
			roleIDs = append(roleIDs, int64(roleId.(float64)))
		}
	}
	customClaims := CustomClaims{
		Username:  claims[username].(string),
		UserId:    uid,
		RoleIds:   roleIDs,
		MapClaims: claims,
	}

	return customClaims
}

func NewGinJwtMiddleware() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "macro-mall",
		Key:             []byte(config.GlobalConfig.Jwt.Key),
		Timeout:         time.Duration(config.GlobalConfig.Jwt.Timeout) * time.Hour,
		MaxRefresh:      time.Duration(config.GlobalConfig.Jwt.MaxRefresh) * time.Hour,
		IdentityKey:     username,
		PayloadFunc:     payloadFunc(),
		IdentityHandler: identityHandler,
		Authenticator:   authenticator(),
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		Unauthorized: func(c *gin.Context, code int, message string) {
			errMsg := c.Request.Header.Get("error")
			if errMsg != "" {
				response.InternalServerError(c, errMsg)
				return
			} else {
				response.Unauthorized(c)
			}
		},
		LoginResponse: func(c *gin.Context, code int, message string, time time.Time) {
			if code == http.StatusOK {
				data := map[string]interface{}{
					"tokenHead": config.GlobalConfig.Jwt.TokenHead,
					"token":     message,
				}
				response.Success(c, data)
			} else {
				response.Fail(c, fmt.Sprintf("login fail for reason: %s", message))
			}
		},
		LogoutResponse: func(c *gin.Context, code int) {
			response.Success(c, nil)
		},
		SendCookie: true,
	})

	if err != nil {
		return nil, err
	}

	return authMiddleware, nil
}
