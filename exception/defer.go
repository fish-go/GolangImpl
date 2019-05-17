package exception

import "fmt"

/**
有一个defer栈，在程序即将退出(无论是异常还是return)时执行
 */
func tryDefer()  {
	defer fmt.Println(1)
	defer fmt.Println(2)

	fmt.Println(3)

	return

	fmt.Println(4)
}
// 3
// 2
// 1