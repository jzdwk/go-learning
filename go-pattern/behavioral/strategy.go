package behavioral

import "fmt"

/*
在策略模式（Strategy Pattern）中，一个类的行为或其算法可以在运行时更改。这种类型的设计模式属于行为型模式。
在策略模式中，我们创建表示各种策略的对象和一个行为随着策略对象改变而改变的 context 对象。策略对象改变 context 对象的执行算法。

使用场景：
	1、一个系统需要动态地在几种算法/策略中选择一种。
	2、如果一个对象有很多的行为，针对每一种行为，不同的具体实现上行为表现又不同。

关键代码：
	1、将对象的行为分离出来，比如跑步行为、走路行为，定义为接口
	2、针对不同对象，实现具体的行为，比如人类的走路行为实现和动物的走路行为实现
	3、context（主体）中持有各种策略接口的引用，将适当的策略注入，并在主体的行为中调用策略的具体实现

和状态模式的不同：
	最大的不同点在于：状态需要在内部维持状态迁移，并且各个状态都需要对主体定义的业务行为进行实现。

*/

//方法1，将strategy抽象为接口后，定义不同的struct实现
//定义策略接口
type strategy interface {
	//执行策略
	doStrategy()
	//get strateMsg
	getStrategyMsg() string
}

//不同的策略实现
type strategyA struct {
	msg string
}

func (s *strategyA) doStrategy() {
	fmt.Println("do strategy A. ")
}
func (s *strategyA) getStrategyMsg() string {
	return s.msg
}

type strategyB struct {
	msg string
}

func (s *strategyB) doStrategy() {
	fmt.Println("do strategy B. ")
}
func (s *strategyB) getStrategyMsg() string {
	return s.msg
}

//传统stg
type contextStg struct {
	//持有引用
	stg strategy
}

func (c *contextStg) setStg(stg strategy) {
	c.stg = stg
}
func (c *contextStg) doAction() {
	c.stg.doStrategy()
}

//将所有的策略封装进map后，不需要再执行if..else
type contextStgWithMap struct {
	//持有引用，并通过map来定位不同的执行
	stg map[string]strategy
}

func (c *contextStgWithMap) setStg(stg strategy) {
	if _, ok := c.stg[stg.getStrategyMsg()]; !ok {
		c.stg[stg.getStrategyMsg()] = stg
	}
}
func (c *contextStgWithMap) doAction(msg string) {
	if k, ok := c.stg[msg]; ok {
		k.doStrategy()
	}
}

//方法2，直接使用函数变量作为strategy的具体实现
type contextStgWithFunc struct {
	stg map[string]func()
}

func (c *contextStgWithFunc) setStgFunc(msg string, f func()) {
	if _, ok := c.stg[msg]; !ok {
		c.stg[msg] = f
	}
}

//调用具体执行逻辑
func (c *contextStgWithFunc) doAction(msg string) {
	if k, ok := c.stg[msg]; ok {
		k()
	}
}

//定义两个逻辑
var strategyAFunc = func() {
	fmt.Println("do strategy A")
}
var strategyBFunc = func() {
	fmt.Println("do strategy B")
}

//返回一个已经持有A,B策略的context
func NewContextStgWithFunc() *contextStgWithFunc {
	return &contextStgWithFunc{stg: map[string]func(){"A": strategyAFunc, "B": strategyBFunc}}
}
func NewContextStgWithMap() *contextStgWithMap {
	return &contextStgWithMap{stg: make(map[string]strategy)}
}
func NewContextSt() *contextStg {
	return &contextStg{}
}

func NewStrategyA() *strategyA {
	return &strategyA{msg: "A"}
}
func NewStrategyB() *strategyB {
	return &strategyB{msg: "B"}
}
