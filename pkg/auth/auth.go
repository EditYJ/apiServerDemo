// 账户加密验证相关
package auth

import "golang.org/x/crypto/bcrypt"

// 加密
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// 验证密码的正确性
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}