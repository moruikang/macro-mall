// @Author moruikang
// @Date 2025/3/16 20:50:00
// @Desc

package bcryptx

import (
	"golang.org/x/crypto/bcrypt"
)

// EncodingPassword 加密密码，使用bcrypt算法
func EncodingPassword(password string) (string, error) {
	cost := 16
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hashedPassword), err
}

// VerifyPassword 验证密码
func VerifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
