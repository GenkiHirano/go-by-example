package files

import (
	"crypto/sha1"
	"fmt"
)

func Task56() {
	fmt.Println()

	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
