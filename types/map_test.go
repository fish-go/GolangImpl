package types

import (
	"fmt"
	"testing"
)

func Test_CRUD(t *testing.T) {
	// 尝试获得不存在的值会返回空字符串
	value, ok := tutorials["Math"]
	if ok {
		fmt.Println("Math 属性不存在")
	} else {
		fmt.Println(value)
	}

	// 添加数据
	tutorials["PythonImpl"] = "Python实战课程"

	fmt.Println(tutorials)

	// 删除字段
	delete(tutorials, "PythonImpl")
}

func Test_loop(t *testing.T) {
	fmt.Println(tutorials)
	loop()
}
