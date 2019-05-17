package function

import (
	"fmt"
	"github.com/logrusorgru/aurora"
)

/**
sum是闭包变量
 */
func addFactory() func(value int) int {
	sum := 0
	return func(value int) int {
		sum = sum + value
		fmt.Println(aurora.Green(sum))
		return sum
	}
}
