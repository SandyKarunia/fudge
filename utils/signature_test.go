package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testPayload = "{\"a\":123.4,\"bb\":\"xyz\",\"ccc\":true,\"dddd\":[\"\",9],\"eeeee\":{\"ffffff\":false}}"

func toStrPtr(str string) *string {
	return &str
}

func TestHMACSign(t *testing.T) {
	tests := []struct {
		desc           string
		timestamp      time.Time
		nonce          string
		secret         string
		payload        *string
		expectedOutput string
	}{
		{
			"happy path",
			time.Unix(1400011122, 33),
			"4092540c-aae6-49c5-8a8c-69de7317fde9",
			"d6qOq9Bu9vQ76ClWrV8J",
			toStrPtr(testPayload),
			"pHu3MHpNq8vdTJCI862zFCLz8TRu5ROmh5U4r4R2SvI=",
		},
		{
			"empty payload",
			time.Unix(1032177401, 10183),
			"16844d31-037c-4ef4-b80b-7bd80ba515eb",
			"oS2vtm1A",
			nil,
			"",
		},
	}

	obj := &signatureImpl{}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			res := obj.hmacSign(test.timestamp, test.nonce, test.secret, test.payload)
			assert.Equal(t, test.expectedOutput, res, test.desc)
		})
	}
}

func TestIsValidHMACSignature(t *testing.T) {
	now := time.Unix(1400011122, 33)
	tests := []struct {
		desc           string
		signature      string
		timestamp      time.Time
		nonce          string
		secret         string
		payload        *string
		expectedOutput bool
	}{
		{
			"happy path",
			"pHu3MHpNq8vdTJCI862zFCLz8TRu5ROmh5U4r4R2SvI=",
			now,
			"4092540c-aae6-49c5-8a8c-69de7317fde9",
			"d6qOq9Bu9vQ76ClWrV8J",
			toStrPtr(testPayload),
			true,
		},
		{
			"timestamp too early",
			"pHu3MHpNq8vdTJCI862zFCLz8TRu5ROmh5U4r4R2SvI=",
			now.Add(-time.Duration(signatureMaxRequestDuration) - time.Second),
			"4092540c-aae6-49c5-8a8c-69de7317fde9",
			"d6qOq9Bu9vQ76ClWrV8J",
			toStrPtr(testPayload),
			false,
		},
		{
			"timestamp too late",
			"pHu3MHpNq8vdTJCI862zFCLz8TRu5ROmh5U4r4R2SvI=",
			now.Add(time.Duration(signatureMaxRequestDuration) + time.Second),
			"4092540c-aae6-49c5-8a8c-69de7317fde9",
			"d6qOq9Bu9vQ76ClWrV8J",
			toStrPtr(testPayload),
			false,
		},
		{
			"wrong signature",
			"abc123",
			now,
			"4092540c-aae6-49c5-8a8c-69de7317fde9",
			"d6qOq9Bu9vQ76ClWrV8J",
			toStrPtr(testPayload),
			false,
		},
	}

	obj := &signatureImpl{}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			res := obj.IsValidHMACSignature(test.signature, now, test.timestamp, test.nonce, test.secret, test.payload)
			assert.Equal(t, test.expectedOutput, res, test.desc)
		})
	}
}
