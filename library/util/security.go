package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GetSha256Digest(src string) string {
	key := []byte("Cosmonaut_sign_v1")
	bytes := []byte(src)
	h := hmac.New(sha256.New, key)
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}
