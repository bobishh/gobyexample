package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	seconds := now.Unix()
	nanos := now.UnixNano()
	puts := fmt.Println

	puts(now)

	millis := nanos / 1000000

	puts(nanos)
	puts(millis)
	puts(seconds)

	puts(time.Unix(seconds, 0))
	puts(time.Unix(0, nanos))
}
