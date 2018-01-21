package main

import b64 "encoding/base64"
import "fmt"

func main() {
	puts := fmt.Println
	data := "asd123asd1!?$*&()'-=@~"
	puts("standard bas64:")
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	puts("encoded: ", sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	puts("decoded: ", string(sDec))
	puts("-----------------------")

	puts("urlsafe base64:")
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	puts("encoded: ", uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	puts("decoded: ", string(uDec))
}
