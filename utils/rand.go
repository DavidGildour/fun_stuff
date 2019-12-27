package utils

import (
	"math/rand"
	"time"
)

func GetRand() *rand.Rand {
	s := rand.NewSource(time.Now().UnixNano())
	return rand.New(s)
}
