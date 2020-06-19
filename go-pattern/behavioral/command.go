package behavioral

import "fmt"

/*
命令模式（Command Pattern）是一种数据驱动的设计模式，它属于行为型模式。
命令最主要的作用是将命令的实际执行步骤进行了封装，比如我们有一个遥控器，遥控器可以控制灯光和空调。
对于遥控器来说，它只关心开和关这种操作，而灯or空调的开/关实现不需要它关心。（1.这就需要屏蔽实现）
另一方面，遥控器的开关应该是可扩展的，比如开关键也可以再控制电视，因此，就需要将开关这个行为抽象出来，作为command

关键代码：
	定义三个角色：1、received,也就是命令最终的实际执行/实现者
	2、真正的命令执行对象Command,这个Command的接口比较固定，exec就是执行，undo就是撤销等等
	3、invoker,使用命令对象的入口。客户最终要使用该对象去调用命令，即遥控器

优点： 1、降低了系统耦合度。 2、新的命令可以很容易添加到系统中去。
缺点：使用命令模式可能会导致某些系统有过多的具体命令类

refer to https://www.runoob.com/design-pattern/command-pattern.html
*/

//定义两个实体，一个灯，一个空调，也就是命令的最终实际执行者receiver
//灯接口
type light interface {
	//其中的方法是按receiver的实际而定，表示灯的行为
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
		fmt.Println(" my light off")
	}
}

//空调
type airCondition interface {
	//其中的方法是按receiver的实际而定，表示灯的行为
	on()
	off()
}
type myAirCondition struct {
	isOn bool
}

func (l *myAirCondition) on() {
	if !l.isOn {
		l.isOn = true
		fmt.Println(" my air condition on")
	}
}
func (l *myAirCondition) off() {
	if l.isOn {
		l.isOn = false
		fmt.Println(" my air condition off")
	}
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//定义命令接口，接口中描述了命令exec和undo
type command interface {
	execute()
	undo()
}

//各个具体命令的实现
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
	fmt.Println("cmd execute light undo")
}

//关灯命令，同样实现了command接口
type lightOffCmd struct {
	light light
}

func (lc *lightOffCmd) execute() {
	//执行时调用light的on，light就是receiver
	fmt.Println("cmd execute light off")
	lc.light.off()
}
func (lc *lightOffCmd) undo() {
	fmt.Println("cmd execute light undo")
}

//定义空调的打开命令
type airCondOnCmd struct {
	airCondition light
}

func (ac *airCondOnCmd) execute() {
	//执行时调用light的on，light就是receiver
	fmt.Println("cmd execute airCondition on")
	ac.airCondition.on()
}
func (ac *airCondOnCmd) undo() {
	fmt.Println("cmd execute airCondition undo")
}

//关灯命令，同样实现了command接口
type airCondOffCmd struct {
	airCondition light
}

func (ac *airCondOffCmd) execute() {
	//执行时调用light的on，light就是receiver
	fmt.Println("cmd execute airCondition off")
	ac.airCondition.off()
}
func (ac *airCondOffCmd) undo() {
	fmt.Println("cmd execute airCondition undo")
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//最后，将这些cmd通过invoker进行封装
type invoker struct {
	onCmd  []command //开按钮集合
	offCmd []command //关按钮集合
}

func (iv *invoker) addOnCMD(cmds ...command) {
	for _, v := range cmds {
		iv.onCmd = append(iv.onCmd, v)
	}
}
func (iv *invoker) addOffCMD(cmds ...command) {
	for _, v := range cmds {
		iv.offCmd = append(iv.offCmd, v)
	}
}

//执行所有on命令
func (iv *invoker) execOn() {
	for _, v := range iv.onCmd {
		v.execute()
	}
}

//执行所有off命令
func (iv *invoker) execOff() {
	for _, v := range iv.offCmd {
		v.execute()
	}
}

func NewInvoker() invoker {
	return invoker{}
}

func NewLightOnCMD(light light) command {
	return &lightOnCmd{light: light}
}
func NewLightOffCMD(light light) command {
	return &lightOffCmd{light: light}
}
func NewAirCondOnCMD(air airCondition) command {
	return &airCondOnCmd{airCondition: air}
}
func NewAirCondOffCMD(air airCondition) command {
	return &airCondOffCmd{airCondition: air}
}

func NewLight() light {
	return &myLight{isOn: false}
}
func NewAirCond() airCondition {
	return &myAirCondition{isOn: false}
}
