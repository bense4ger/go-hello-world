package helper

import (
	"math/rand"
	"time"
)

// RandomDuration produces a random duation of up to five seconds
func RandomDuration() time.Duration {
	r := rand.Intn(5000)

	return time.Duration(r) * time.Millisecond
}
