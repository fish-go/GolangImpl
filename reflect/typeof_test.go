package reflect

import (
	"fmt"
	"testing"
)

func Test_typeof(t *testing.T) {
	fmt.Println(typeof(123)) // int
	fmt.Println(typeof(fmt.Println)) // func(...interface {}) (int, error)

	arr :=  [...]int{1,2}
	fmt.Println(typeof(arr)) // [2]int

	sli := arr[:]
	fmt.Println(typeof(sli)) // []int
}
