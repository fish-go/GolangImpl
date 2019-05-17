package types

/**
slice是数组的一个view
 */

func passByRef(slice []int) {
	for i, v := range slice{
		slice[i] = v * v
	}
}
