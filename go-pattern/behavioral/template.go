package behavioral

import "fmt"

/*
在模板模式（Template Pattern）中，一个抽象类公开定义了执行它的方法的方式/模板。
它的子类可以按需要重写方法实现，但调用将以抽象类中定义的方式进行。
这种类型的设计模式属于行为型模式。

关键代码：
	1. 定义一个操作中的算法的骨架，而将一些步骤延迟到子类中。
	2. 模板方法使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。
使用场景：
	1、有多个子类共有的方法，且逻辑相同。
	2、重要的、复杂的方法，可以考虑作为模板方法。
注意：
	对于java等面向对象语言，可以通过抽象类/方法的定义，将需要实现的个性化算法步骤延延迟到子类中去个性化实现，
	但是对于go，没有严格意义上的继承，父struct无法定义“空方法”，并通过子struct去具体实现。
	因此，父struct中持有一个template接口，用于调用子struct的实现。
*/
//定义一个基本的算法接口
type baseTp interface {
	doFirst()
	doSecond()
	doThird()
	doTp()
}

//一个顶层的实现
type baseTemplate struct {
	//一个base接口，通过赋值具体的子struct实现来得到延迟实现的场景
	sub baseTp
	//一个用于控制算法流程的hook，继承该struct的struct可以通过这个hook改变流程
	hook bool
}

//各个步骤的内部实现
func (b *baseTemplate) doFirst() {
	fmt.Println("do first step")
}

func (b *baseTemplate) doThird() {
	fmt.Println("do third step")
}

//对外的方法
func (b *baseTemplate) doTp() {
	b.doFirst()
	//调用子struct的实现
	if b.sub != nil {
		b.sub.doSecond()
	}
	//根据hook，决定3是否执行
	if b.hook {
		b.doThird()
	}
}

//set
func (b *baseTemplate) setTp(sub baseTp) {
	b.sub = sub
}

//两个子实现继承了baseTemplate，从而完成了方法doFirst和doThird
type subOneTp struct {
	//subOne继承了baseTemplate的方法
	*baseTemplate
}

func (s *subOneTp) doSecond() {
	//second使用自己的实现
	fmt.Println("do sub one second step")
}

type subTwoTp struct {
	//subTwo继承了baseTemplate的方法
	*baseTemplate
}

func (s *subTwoTp) doSecond() {
	//second使用自己的实现
	fmt.Println("do sub two second step")
	s.baseTemplate.hook = false
	fmt.Println("set hook to false")
}

func NewSubOneTp(base *baseTemplate) *subOneTp {
	return &subOneTp{base}
}
func NewSubTwoTp(base *baseTemplate) *subTwoTp {
	return &subTwoTp{base}
}
func NewBaseTp(sub baseTp) *baseTemplate {
	return &baseTemplate{hook: true, sub: sub}
}
