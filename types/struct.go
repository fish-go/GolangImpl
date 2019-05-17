package types

/**
声明一个结构体
 */
type TreeNode struct {
	Left, Right *TreeNode
	value int
}

/**
在函数体中返回一个局部变量的地址在C++中是大忌
但是在Go中，是函数式编程，有闭包，可以的

编译器根据具体的用法来决定在栈或者堆上分配内存
 */
func createTreeNode(value int) *TreeNode  {
	return &TreeNode{value:value}
}


var root TreeNode = TreeNode{
	value: 3,
}
