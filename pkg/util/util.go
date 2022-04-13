package util

import (
	"fmt"
	"math/rand"
	"time"
)

// return 6 digtal random num
func CreateCaptcha() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
