package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// GenerateSalt 生成指定长度的随机盐值
func GenerateSalt(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("salt length must be positive")
	}

	salt := make([]byte, length)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(salt), nil
}

// HashWithSalt 对数据进行MD5加盐哈希
func HashWithSalt(data string, salt string) string {
	// 将数据和盐值组合
	saltedData := data + salt

	// 计算MD5哈希
	hash := md5.Sum([]byte(saltedData))

	// 返回十六进制字符串
	return hex.EncodeToString(hash[:])
}

// VerifyWithSalt 验证数据与加盐哈希是否匹配
func VerifyWithSalt(data string, salt string, expectedHash string) bool {
	actualHash := HashWithSalt(data, salt)
	return actualHash == expectedHash
}
