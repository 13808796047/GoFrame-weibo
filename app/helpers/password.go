package helpers

import (
	"github.com/gogf/gf/os/glog"
	"golang.org/x/crypto/bcrypt"
)

// Hash 使用bcrypt对密码进行加密
func Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		glog.Error(err)
	}
	return string(bytes)
}

// CheckHash 对比明文密码和数据库的哈希值
func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return false
	}
	return err == nil
}

// IsHashed 判断字符串是否哈希过
func IsHashed(str string) bool {
	// bcrypt 加密后的长度等于60
	return len(str) == 60
}
