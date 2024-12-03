package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/spf13/viper"
)

var secret string

func init() {
	secret = viper.GetString("auth.encrypt_secret")
}

func EncryptPassword(oPassword string) string {
	// 创建一个 SHA-256 哈希对象
	h := sha256.New()

	// 将秘密值写入哈希对象
	h.Write([]byte(secret))

	// 将原始密码写入哈希对象，计算哈希值，返回十六进制表示的哈希字符串
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
