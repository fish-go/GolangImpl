package schema

/**
定义一个接口，表示这个对象必须有Run() 方法
 */
type Runnable interface {
	Run()
}

/**
作为契约来实现
 */
func Runner(runnable Runnable)  {
	runnable.Run()
}
