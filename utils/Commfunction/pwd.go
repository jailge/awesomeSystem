package Commfunction

import "golang.org/x/crypto/bcrypt"

// ***************
// 使用Bcrypt实现加密或验证密码
// ***************

// PasswordHash 密码加密
func PasswordHash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

// PasswordVerify 密码验证
func PasswordVerify(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
