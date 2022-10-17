package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

func GetRandomIndex(max int) int {
	return rand.Intn(max)
}
