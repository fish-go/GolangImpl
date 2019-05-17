package types

import (
	"fmt"
	"testing"
)

func Test_createTreeNode(t *testing.T) {
	node := createTreeNode(3)

	fmt.Println(*node)
	// {<nil> <nil> 3}
}
