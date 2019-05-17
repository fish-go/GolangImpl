package function

import (
	"testing"
)

func Test_addFactory(t *testing.T) {
	add := addFactory()
	for i := 0; i < 10 ; i++ {
		add(i)
	}
}
