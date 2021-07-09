package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"strings"
)

func main()  {
	e := 0.0
	er := 1.0
	fmt.Println(e/er)
	var a bytes.Buffer
	a.WriteString("aaaa")
	fmt.Println(a.Bytes())

	var b strings.Builder
	b.WriteString("key=123&")
	b.WriteString("aaa你好")

	h := sha1.New()

	fmt.Println(h.Sum([]byte(b.String())))

	_, _ = io.WriteString(h, "key=123&aaa你好")
	fmt.Sprintf("%x", h.Sum(nil))

	str := "ABCDEFGHI"

	for i := 0; i < len(str) ; i++ {
		fmt.Println(str[i])
	}

}
