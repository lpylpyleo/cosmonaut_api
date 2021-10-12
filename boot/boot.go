package boot

import (
	_ "cosmonaut_api/packed"
	_ "github.com/lib/pq"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
