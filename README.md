一个个包都是安装在gopath的src目录下的


Java是一个个类组成的

go是一个个函数和type组成的

别人import一个文件，应该会执行其中的代码

go文件的格式

- type

- var

- func
  - 静态方法
  - 成员方法

go file中不能有运算


包
- 包名和目录名可以不一致
- 为结构体定义方法只能在一个包里面，
  肯定的，不然就乱套了
- public private是针对包来说的

```bash
$ go test ./oop/monkey_test.go ./oop/monkey.go
ok      command-line-arguments  0.005s

$ go test ./oop/monkey_test.go                
# command-line-arguments [command-line-arguments.test]
oop/monkey_test.go:10:12: undefined: Monkey
oop/monkey_test.go:20:12: undefined: Monkey
FAIL    command-line-arguments [build failed]

# 目前只能这样，测试整个包，服了
$ go test ./oop
ok      _/Users/apple/Desktop/program/go/GolangImpl/oop 0.084s


```
这个地方我是服了，直接在IDE中点击运行测试的话会触发
下面的命令，导致找不到函数！！！服了

只能右键run go test PACKAGE

这样也还行把，就是麻烦，可能挡误时间，结果不明确

一次只能测试一整个包。。需要color来显眼

没有pom.xml 以及 package.json
包依赖管理动态获取到GOPATH，
编译流程由go 命令行接管