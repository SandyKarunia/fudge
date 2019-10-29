package utils

import (
	"math/rand"
)

// String ...
//go:generate mockery -name=String
type String interface {
	// GenerateRandomAlphanumeric generates a random string based on passed specifications
	GenerateRandomAlphanumeric(length int) string
}

type stringImpl struct{}

func (s *stringImpl) GenerateRandomAlphanumeric(length int) string {
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

// ProvideString ...
func ProvideString() String {
	return &stringImpl{}
}
