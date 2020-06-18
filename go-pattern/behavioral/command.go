package behavioral

import "fmt"

/*
命令模式（Command Pattern）是一种数据驱动的设计模式，它属于行为型模式。
命令模式将每一个请求都抽象出来，以命令的形式包裹封装对象中，并传给调用对象。
调用对象寻找可以处理该命令的合适的对象，并把该命令传给相应的对象，该对象执行命令。
*/

//定义两个实体，一个灯，一个空调
type light interface {
	on()
	off()
}
type myLight struct {
	isOn bool
}

func (l *myLight) on() {
	if !l.isOn {
		l.isOn = true
		fmt.Println(" my light on")
	}

}
func (l *myLight) off() {
	if l.isOn {
		l.isOn = false
		fmt.Println(" my light on")
	}
}

type airCondition interface {
	on()
	off()
}
type myAirCondition struct {
	isOn bool
}

func (l *myAirCondition) on() {
	if !l.isOn {
		l.isOn = true
		fmt.Println(" my light on")
	}

}
func (l *myAirCondition) off() {
	if l.isOn {
		l.isOn = false
		fmt.Println(" my light on")
	}
}

//定义命令接口，接口中描述了命令执行和撤下
type command interface {
	execute()
	undo()
}

//定义一个灯的打开命令
type lightOnCmd struct {
	light light
}

func (lc *lightOnCmd) execute() {
	//执行时调用light的on，light就是receiver
	fmt.Println("cmd execute light on")
	lc.light.on()
}
func (lc *lightOnCmd) undo() {
	fmt.Println("cmd execute light on undo")
}

type lightOffCmd struct {
	light light
}

func (lc *lightOffCmd) execute() {
	//执行时调用light的on，light就是receiver
	fmt.Println("cmd execute light on")
	lc.light.off()
}
func (lc *lightOffCmd) undo() {
	fmt.Println("cmd execute light on undo")
}
