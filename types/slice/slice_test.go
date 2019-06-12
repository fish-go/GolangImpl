package slice

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func Test_passByRef(t *testing.T) {
	var arr [3]int = [...]int{1, 2, 3}

	var slice []int = arr[:]

	// 是在原来的数组上进行的修改
	passByRef(slice)

	// 所以数组这步的值发生了变化
	assert.Equal(t, 9, slice[2])
}


func Test_getAddressOfItem(t *testing.T) {
	var arr [3]int = [...]int{1, 2, 3}

	var slice []int = arr[:]

	a := getAddressOfItem(slice)

	*a = 999

	fmt.Println(arr)
}
