package main

import (
	"flag"
	"fmt"
)
var xx string
var age int
func init()  {
	flag.IntVar(&age, "age", 0, "commender age")
	flag.StringVar(&xx, "name", "", "commender name")
}
func main()  {
	flag.Parse()
	fmt.Println(xx, " age is ", age)
}
