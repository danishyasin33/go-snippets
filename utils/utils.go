package utils

import (
	"crypto/rand"
	"fmt"
	"math"
	"time"
)

// timeSince returns a duration in milliseconds between a start time and now.
func TimeSince(start time.Time) int {
	return int(math.Floor(float64(time.Since(start))/float64(time.Millisecond) + 0.5))
}

// V4 returns a random UUID v4 string.
func V4() string {
	b := make([]byte, 16)
	rand.Read(b)
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
