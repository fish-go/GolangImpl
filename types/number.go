package types

import (
	"fmt"
	"unsafe"
)

func sizeOfInteger()  {
	// 平台相关
	var a int
	var b int32
	var c uint64

	fmt.Println(
		unsafe.Sizeof(a), // 8
		unsafe.Sizeof(b), // 4
		unsafe.Sizeof(c)) // 8
}