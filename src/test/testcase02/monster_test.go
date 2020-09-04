package testcase02

import (
	"fmt"
	"testing"
)

func TestStore(t *testing.T) {
	monster := Monster{
		Name:  "hahah",
		Age:   10,
		Skill: "ahaha",
	}
	res := monster.Store()
	if res != true {
		fmt.Println("测试失败")
	}
}

func TestRestore(t *testing.T) {
	monster := Monster{}
	ok := monster.Restore()
	if ok != true {
		fmt.Println("测试失败")
	}
}
