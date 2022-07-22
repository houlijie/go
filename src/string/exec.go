package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
)

func main() {
	e := 0.0
	er := 1.0
	fmt.Println(e / er)
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

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

}

const letter = "abcdefghijklmnopqrstuvxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStr(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}

func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}

	return s
}

func fmtSprintf(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}

	return s
}

func strBuilder(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}

	return builder.String()
}

func strBuffer(n int, s string) string {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(s)
	}

	return buf.String()
}

func strBytes(n int, s string) string {
	buf := make([]byte, 0, n*len(s))
	for i := 0; i < n; i++ {
		buf = append(buf, s...)
	}

	return string(buf)
}
