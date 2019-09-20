package utils

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGenerateRandomAlphanumericString(t *testing.T) {
	allowedChars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tests := []struct {
		description    string
		length         int
		expectedLength int
	}{
		{
			description:    "length = 0",
			length:         0,
			expectedLength: 0,
		},
		{
			description:    "length = -1",
			length:         -1,
			expectedLength: 0,
		},
		{
			description:    "length = 1000",
			length:         1000,
			expectedLength: 1000,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			obj := &stringImpl{}
			res := obj.GenerateRandomAlphanumeric(test.length)
			assert.Equal(t, test.expectedLength, len(res), test.description)
			// check for allowed chars
			for _, c := range res {
				if !strings.Contains(allowedChars, string(c)) {
					t.Fatalf("Result contains illegal character: %c", c)
				}
			}
		})
	}
}
