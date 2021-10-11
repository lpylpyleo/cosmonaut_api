package test

import (
	"cosmonaut_api/library/util"
	"testing"
)

func TestRandomString(t *testing.T) {
	result := util.RandomString(10)
	t.Log(result)
}

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		util.RandomString(20)
	}
}
