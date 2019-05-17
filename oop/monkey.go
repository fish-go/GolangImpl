package oop

import "fmt"

/**
以文件（或者包）为单位声明对象的格式
 */

type Monkey struct {
	id int32
	name string
	age int64
}

/**
专门服务于这个结构体的函数，
可以理解为对象的方法
 */
func (monkey Monkey) getAge() int64  {
	return monkey.age
}

/**
上下两种写法只是语法上的不同
*/
func getAgeFromMonkey(monkey Monkey) int64  {
	return monkey.age
}

/**
依旧是值传递
 */
func (monkey Monkey) setAgeByValue(age int64) {
	monkey.age = age
}

/**
这里通过指针达到引用传递
*/
func (monkey *Monkey) setAge(age int64) {
	monkey.age = age
}

/**
这个也算是静态方法了
 */
func SayHello()  {
	fmt.Println("Hell world")
}