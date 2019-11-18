package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"time"
)

const (
	// signatureMaxRequestDuration is the max difference of timestamp between
	// server time and the one in the request.
	// It's used to prevent replay attack.
	// If value is too high, then there's a high risk of replay attack;
	// if it's too low, then it could fail many genuine requests
	// (if latency between client and server is high).
	signatureMaxRequestDuration = int64(2 * time.Minute)
)

// Signature ...
//go:generate mockery -name=Signature
type Signature interface {
	// IsValidHMACSignature checks whether the given signature
	// can be reconstructed with the other payloads.
	IsValidHMACSignature(
		signature string,
		now time.Time,
		timestamp time.Time,
		nonce string,
		secret string,
		payload *string,
	) bool
}

type signatureImpl struct {
}

func (s *signatureImpl) IsValidHMACSignature(
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
	if durInt > signatureMaxRequestDuration {
		return false
	}

	return signature == s.hmacSign(timestamp, nonce, secret, payload)
}

func (*signatureImpl) hmacSign(timestamp time.Time, nonce string, secret string, payload *string) string {
	if payload == nil {
		s := ""
		payload = &s
	}

	h := hmac.New(sha256.New, []byte(secret))

	timestampMsec := timestamp.UnixNano() / 1000
	timestampStr := strconv.FormatInt(timestampMsec, 10)

	msg := timestampStr + *payload + nonce
	h.Write([]byte(msg))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// ProvideSignature ...
func ProvideSignature() Signature {
	return &signatureImpl{}
}
