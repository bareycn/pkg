package rand

import (
	crypto "crypto/rand"
	"encoding/base64"
	"math/rand"
)

const (
	number       = "0123456789"
	lower        = "abcdefghijklmnopqrstuvwxyz"
	upper        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letter       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberLetter = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// String 生成指定长度的随机字符串
func String(length int) string {
	return Generate(length, numberLetter)
}

// Num 生成指定长度的随机数字字符串
func Num(length int) string {
	return Generate(length, number)
}

// Lower 生成指定长度的随机小写字母字符串
func Lower(length int) string {
	return Generate(length, lower)
}

// Upper 生成指定长度的随机大写字母字符串
func Upper(length int) string {
	return Generate(length, upper)
}

// Letter 生成指定长度的随机字母字符串
func Letter(length int) string {
	return Generate(length, letter)
}

func Nonce(length int) string {
	b := make([]byte, 16)
	_, err := crypto.Read(b)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// Generate 生成指定长度的随机数字字符串
func Generate(length int, str string) string {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = str[rand.Intn(len(str))]
	}
	return string(bytes)
}
