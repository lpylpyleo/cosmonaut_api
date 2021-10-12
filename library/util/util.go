package util

import (
	"math/rand"
	"strings"
)

const (
	all   = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	upper = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//lower = "0123456789abcdefghijklmnopqrstuvwxyz"
)

func RandomString(l int) string {
	var b strings.Builder
	for i := 0; i < l; i++ {
		b.WriteString(string(all[rand.Intn(len(all))]))
	}
	return b.String()
}

func GenUid() string {
	var b strings.Builder
	l := len(upper)
	for i := 0; i < 8; i++ {
		b.WriteString(string(upper[rand.Intn(l)]))
	}
	return b.String()
}
