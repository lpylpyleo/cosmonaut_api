package util

import (
	"math/rand"
	"strings"
)

const all = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const allLen = len(all)

func RandomString(l int) string {
	var b strings.Builder
	for i := 0; i < l; i++ {
		b.WriteString(string(all[rand.Intn(allLen)]))
	}
	return b.String()
}
