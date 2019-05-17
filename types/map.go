package types

import (
	"fmt"
)

/**
这个内置map实战是太完美了，JS的对象字面量写法完全享受到了
 */

var tutorials = map[string]string{
	"GolangImpl": "Golang实战课程",
}

func loop()  {
	tutorials["JavaImpl"] = "Java课程实战"
	for k, v := range tutorials {
		fmt.Println(k, v)
	}
}
