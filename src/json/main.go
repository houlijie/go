package main

import (
	"encoding/json"
	"fmt"
)

type stu struct {
	Name string `json:"name"`
}

func main() {
	a1 := stu{Name: "zhangsan"}

	j1, _ := json.Marshal(a1)
	fmt.Println(string(j1)) // {"name":"zhangsan"}

	var j3 stu
	err := json.Unmarshal(j1, &j3)
	if err != nil {
		return
	}
	fmt.Println(j3.Name)
}
