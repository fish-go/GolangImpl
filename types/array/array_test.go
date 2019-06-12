package array

import (
	"testing"

	"gotest.tools/assert"
)

func Test_sumForLoop(t *testing.T) {
	sum := sumForLoop([5]int{1, 2, 3, 4, 5})
	assert.Equal(t, 15, sum)
}

func Test_sumRangeLoop(t *testing.T) {
	sum := sumRangeLoop([5]int{1, 2, 3, 4, 5})
	assert.Equal(t, 15, sum)
}

/**
golang中，数组都是值传递的
 */
func Test_passByValue(t *testing.T)  {
	arr := [3]int{1, 2, 3}
	changeArrayValue(arr)

	// 函数体中对array的修改不会影响到原来的数组
	// 这个位置的值依旧是1
	assert.Equal(t, 1, arr[0])
}
