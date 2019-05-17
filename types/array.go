package types

/**
声明一个数组，数组的长度是参与到类型检查的
所以Go的数组类型检查比其他语言更加精细
 */
var arr [3]int = [...]int{1, 2, 3}

// 索引递增来变量数组
func sumForLoop(arr [5]int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

/**
通过range来循环，拿到index和value
 */
func sumRangeLoop(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// 修改数组中的元素的值，但是它是应用传递的，修改对外部不可见
func changeArrayValue(arr [3]int){
	arr[0] = arr[0] * arr[0]
}
