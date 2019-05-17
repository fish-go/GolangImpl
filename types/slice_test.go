package types

import (
	"gotest.tools/assert"
	"testing"
)

func Test_passByRef(t *testing.T) {
	var arr [3]int = [...]int{1, 2, 3}

	var slice []int = arr[:]

	// 是在原来的数组上进行的修改
	passByRef(slice)

	assert.Equal(t, 9, slice[2])
}
