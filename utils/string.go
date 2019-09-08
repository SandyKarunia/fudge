package utils

import "math/rand"

// GenerateRandomAlphanumericString generates a random string based on passed specifications
func GenerateRandomAlphanumericString(length int) string {
	if length <= 0 {
		return ""
	}

	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	res := make([]byte, length)
	for i := 0; i < length; i++ {
		res[i] = chars[rand.Intn(len(chars))]
	}
	return string(res)
}
