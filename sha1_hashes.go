package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 dis please"
	h := sha1.New()

	h.Write([]byte(s))

	// This gets the finalized hash result as a byte slice. The argument to Sum can be used to append to an existing byte slice: it usually isn’t needed.
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
