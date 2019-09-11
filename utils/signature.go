package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"time"
)

const (
	// maxRequestDuration is the max difference of timestamp between
	// server time and the one in the request.
	// It's used to prevent replay attack.
	// If value is too high, then there's a high risk of replay attack;
	// if it's too low, then it could fail many genuine requests
	// (if latency between client and server is high).
	maxRequestDuration = int64(2 * time.Minute)
)

func isValidHMACSignature(
	signature string,
	now time.Time,
	timestamp time.Time,
	nonce string,
	secret string,
	payload *string,
) bool {
	dur := now.Sub(timestamp)
	durInt := int64(dur)
	if durInt < 0 {
		durInt *= -1
	}
	// To prevent replay attack that happens sometime after the original request.
	if durInt > maxRequestDuration {
		return false
	}

	if payload == nil {
		s := ""
		payload = &s
	}
	return signature == hmacSign(timestamp, nonce, secret, payload)
}

func hmacSign(timestamp time.Time, nonce string, secret string, payload *string) string {
	if payload == nil {
		return ""
	}

	h := hmac.New(sha256.New, []byte(secret))

	timestampMsec := timestamp.UnixNano() / 1000
	timestampStr := strconv.FormatInt(timestampMsec, 10)

	msg := timestampStr + *payload + nonce
	h.Write([]byte(msg))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
