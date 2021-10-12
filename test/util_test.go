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

func TestGetDigest(t *testing.T) {
	result := util.GetSha256Digest("qwe123456")
	t.Log(len(result), result)
}

func BenchmarkGetDigest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		util.GetSha256Digest("qwe123")
	}
}

func TestGenUid(t *testing.T) {
	for i := 0; i < 10; i++ {
		result := util.GenUid()
		t.Log(result)
	}
}
