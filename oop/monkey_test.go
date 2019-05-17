package oop

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func TestMonkey_getAge(t *testing.T) {
	monkey := Monkey{
		23,
		"Sun WuKong",
		999,
	}
	fmt.Println(monkey.getAge())
	// t.Error("sdf")
}

func TestMonkey_setAge(t *testing.T) {
	monkey := Monkey{
		23,
		"Sun WuKong",
		999,
	}
	// 依旧是999
	monkey.setAgeByValue(4)
	assert.Equal(t, monkey.age, int64(999))

	// 通过传递指针才可以修改值
	monkey.setAge(666)
	assert.Equal(t, monkey.age, int64(666))
}
