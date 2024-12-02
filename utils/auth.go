package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const secret = "lirous.com"

func EncryptPassword(oPassword string) string {
	// 创建一个 SHA-256 哈希对象
	h := sha256.New()

	// 将秘密值写入哈希对象
	h.Write([]byte(secret))

	// 将原始密码写入哈希对象，计算哈希值，返回十六进制表示的哈希字符串
	fmt.Println(hex.EncodeToString(h.Sum([]byte(oPassword))))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
