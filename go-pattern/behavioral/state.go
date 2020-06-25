package behavioral

import "fmt"

/*
在状态模式（State Pattern）中，类的行为是基于它的状态改变的。这种类型的设计模式属于行为型模式。
在状态模式中，我们创建表示各种状态的对象和一个行为随着状态对象改变而改变的 context 对象。

主要解决：
	对象的行为依赖于它的状态（属性），并且可以根据它的状态改变而改变它的相关行为。

使用场景：
	代码中包含大量与对象状态有关的条件语句。

关键代码：
	1. 将各种具体的状态类抽象出来。
	2. 定义状态接口，各个抽象出来的状态类实现该接口。
	3. 状态接口定义了业务逻辑中的执行方法，只是在不同的状态实现中，对应的业务方法表现不同。
	4. 状态实现通过引用context给context赋值，context通过state的引用调用state对应的业务方法。

和命令模式的差别：

	1. 通常命令模式的接口中只有一个方法。而状态模式的接口中有一个或者多个方法。
	2. 状态模式的实现类的方法，一般返回值，或者是改变实例变量的值。也就是说，状态模式一般和对象的状态有关。实现类的方法有不同的功能，覆盖接口中的方法。状态模式和命令模式一样，也可以用于消除 if...else 等条件选择语句。
*/

//定义状态接口
type state interface {
	//注意，状态接口中需要定义业务逻辑执行的各个方法
	doRun()
	doStop()
}

//1. 开始状态
type startState struct {
	//state实现中持有context的引用，用于向context中注入迁移状态
	ct context
}

func (s *startState) doRun() {
	fmt.Println("start state cannot run again")
}
func (s *startState) doStop() {
	fmt.Println("stop running, state changes to stop")
	s.ct.setState(&stopState{s.ct})
}

//2 结束状态
type stopState struct {
	//state实现中持有context的引用，用于向context中注入迁移状态
	ct context
}

func (s *stopState) doRun() {
	fmt.Println("start running, state changes to start")
	s.ct.setState(&startState{s.ct})
}
func (s *stopState) doStop() {
	fmt.Println("stop state cannot stop again")
}

//上下文接口最主要的功能就是setState
type context interface {
	//业务逻辑
	//1.启动
	doRun()
	//2.停止
	doStop()
	//状态赋值
	setState(stt state)
}

type myContext struct {
	//业务上下文中保存了状态变量
	stt state
}

func (c *myContext) setState(stt state) {
	c.stt = stt
}
func (c *myContext) doRun() {
	//执行业务代码将调用状态变量的相应代码
	c.stt.doRun()
}
func (c *myContext) doStop() {
	c.stt.doStop()
}

func NewMyContext() context {
	//context 初始为stop状态
	ct := &myContext{}
	stop := &stopState{ct: ct}
	ct.stt = stop
	return ct
}
