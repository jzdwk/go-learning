package structural

/*
装饰器模式（Decorator Pattern）允许向一个现有的对象添加新的功能，同时又不改变其结构。
这种类型的设计模式属于结构型模式，它是作为现有的类的一个包装。

关键代码：
	1、Component 类充当抽象角色，不应该具体实现。
	2、修饰类引用和继承 Component 类，具体扩展类重写父类方法。

使用场景：
	1、扩展一个类的功能。
	2、动态增加功能，动态撤销。
注意：
	该模式从代码结构上和代理模式相似，不同点在于：
	1、代理模式不会向外暴露被代理的对象，装饰者需要指定被装饰的对象
	2、两者都是对类的方法进行增强，但装饰器模式强调的是增强自身，在被装饰之后你能够够在被增强的类上使用增强后的方法。增
	强过后还是你，只不过能力变强了。
		而代理模式则强调要别人帮你去做一些本身与你业务没有太多关系的职责。
	代理模式是为了实现对象的控制，因为被代理的对象往往难以直接获得或者是其内部不想暴露出来。

refer to:
	https://zhuanlan.zhihu.com/p/62783012
	https://www.runoob.com/design-pattern/decorator-pattern.html
*/

//定义要被装饰的对象
type shapeDct interface {
	doDraw() string
}

//实现的实例
type circleDct struct {
	msg string
}

func (c *circleDct) doDraw() string {
	return c.msg
}

//实现1，使用传统的接口和继承
//定义shapeDecoratorA，对circleDct的doDraw做扩展
type shapeDecoratorA struct {
	msg  string
	base shapeDct
}

func (c *shapeDecoratorA) doDraw() string {
	//add func
	if c.base != nil {
		c.addFuncA()
		return c.msg + c.base.doDraw()
	}
	return ""
}
func (c *shapeDecoratorA) addFuncA() {
	c.msg = "A decorated"
}
func NewShapeDecoratorA(base shapeDct) shapeDct {
	return &shapeDecoratorA{base: base}
}

//当需要另一个对circle的扩展，则定义B
type shapeDecoratorB struct {
	msg  string
	base shapeDct
}

func (c *shapeDecoratorB) doDraw() string {
	//add func
	if c.base != nil {
		c.addFuncB()
		return c.msg + c.base.doDraw()
	}
	return ""
}
func (c *shapeDecoratorB) addFuncB() {
	c.msg = "B decorated"
}
func NewShapeDecoratorB(base shapeDct) shapeDct {
	return &shapeDecoratorB{base: base}
}

//实现2,直接定义func变量doDraw
type doDraw func() string

func DecorateDraw(do doDraw) doDraw {
	//add decorate func
	return func() string {
		rst := do()
		//add func
		rst += " decorated "
		return rst
	}
}
