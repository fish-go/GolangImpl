package types

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"strconv"
	"testing"
)

func Test_multipleString(t *testing.T) {
	str := "Jon"
	str = multipleString(str, 10)
	fmt.Println(aurora.BrightGreen(str))
}

func Test_PrintMultilineString(t *testing.T) {
	fmt.Println(aurora.Cyan(multipleLineString))
}

/**
int 类型转化为字符型
 */
func Test_IntToString(t *testing.T) {
	var s string = strconv.Itoa(123)
	fmt.Println(aurora.Yellow(s))
}

func Test_StringToInt(t *testing.T) {
	s, _ := strconv.Atoi("123456")
	fmt.Println(aurora.Red(s))
}
