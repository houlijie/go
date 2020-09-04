package testcase02

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m *Monster) Store() bool {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err=", err)
		return false
	}

	filePath := "monster.txt"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write err=", err)
		return false
	}

	return true
}

func (m *Monster) Restore() bool {
	filePath := "monster.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read err=", err)
		return false
	}

	json.Unmarshal(data, m)

	fmt.Println("m=", m)
	return true
}
