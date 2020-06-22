package behavioral

import "fmt"

/*
中介者模式（Mediator Pattern）是用来降低多个对象和类之间的通信复杂性。
这种模式提供了一个中介类，该类通常处理不同类之间的通信，并支持松耦合，使代码易于维护。
中介者模式属于行为型模式。

使用场景：
	1. 多个组件间都需要互相的通信，形成了网状结构。因此需要一个“中介”，协调/处理各个组件间的通信，从而转化为星状结构。
	2. 多用于调度场景；
	3. MVC架构中的Controller就是M和V的“中介者”，View只需要和Controller通信.
	根据api url，Controller层路由到具体的Controller实现完成业务逻辑
关键代码：
	组件之间的通信封装到一个类中单独处理。Mediator为中介者，Colleague为交互的组件，Colleague持有Mediator的引用
*/
//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//中介者接口
type mediator interface {
	//根据名称，将msg发送给实际的colleague
	sendMsg(name string, msg string)
	//colleague的注册接口
	register(name string, colleague colleague)
}

//中介者的具体实现，其中维护和各个colleague的通信逻辑
//本例子中通过一个route进行路由,进行一对一的转发，也可以为组播/广播等等，根据规则来定（类比于api定义）
type myMediator struct {
	colleagues map[string]colleague
}

func (m *myMediator) sendMsg(name string, msg string) {
	if v, ok := m.colleagues[name]; ok {
		v.receive(msg)
	}
}
func (m *myMediator) register(name string, colleague colleague) {
	if _, ok := m.colleagues[name]; !ok {
		m.colleagues[name] = colleague
	}
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//通信组件接口，各个组件的通用接口定义
type colleague interface {
	sendTo(name string, msg string)
	receive(msg string)
}

//组件A
type colleagueA struct {
	//维护中介者的引用，使用中介者通信
	mediator mediator
}

func (c *colleagueA) sendTo(name string, msg string) {
	fmt.Println("A send msg: " + msg)
	c.mediator.sendMsg(name, msg)
}
func (c *colleagueA) receive(msg string) {
	fmt.Println("A get msg: " + msg)
}

//组件B
type colleagueB struct {
	//维护中介者的引用，使用中介者通信
	mediator mediator
}

func (c *colleagueB) sendTo(name string, msg string) {
	fmt.Println("B send msg: " + msg)
	c.mediator.sendMsg(name, msg)
}
func (c *colleagueB) receive(msg string) {
	fmt.Println("B get msg: " + msg)
}

func NewMyMediator() mediator {
	return &myMediator{colleagues: make(map[string]colleague)}
}

func NewColleagueA(m mediator) colleague {
	return &colleagueA{mediator: m}
}

func NewColleagueB(m mediator) colleague {
	return &colleagueB{mediator: m}
}
