package schema

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
	"testing"
)

type Sheep struct {
}

/**
Sheep有了Run方法等于自动实现了runnable接口
*/
func (sheep Sheep) Run() {
	fmt.Println(Green("a sheep is running!!"))
}

func TestRunner(t *testing.T) {
	sheep := Sheep{}
	Runner(sheep)
}
