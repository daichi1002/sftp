package util

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func GetUlid() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
