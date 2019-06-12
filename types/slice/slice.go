package slice

/**
slice是数组的一个view
 */

func passByRef(slice []int) {
	for i, v := range slice{
		slice[i] = v * v
	}
}

func concatenateTwoSlices(a []int, b []int) []int {
	return append(a, b...)
}

/**
获得切面中，某一个元素的地址
 */
func getAddressOfItem(array []int) *int {
	return &array[0]
}