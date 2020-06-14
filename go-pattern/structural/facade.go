package structural

/*
外观模式（Facade Pattern）隐藏系统的复杂性，并向客户端提供了一个客户端可以访问系统的接口。
主要用于降低访问复杂系统的内部子系统时的复杂度，简化客户端与之的接口。

使用场景：
	1、为复杂的模块或子系统提供外界访问的模块。 2、子系统相对独立。 3、预防低水平人员带来的风险。
优点：
	1、减少系统相互依赖。 2、提高灵活性。 3、提高了安全性。
缺点：
	不符合开闭原则，如果要改东西很麻烦，继承重写都不合适。

refer to: https://www.runoob.com/design-pattern/facade-pattern.html
*/

//定义一个形状接口
type shape interface {
	doDraw() string
}

//圆形实现了形状
type circle struct {
	msg string
}

func (c *circle) doDraw() string {
	return c.msg
}

//三角形实现了形状
type rectangle struct {
	msg string
}

func (r *rectangle) doDraw() string {
	return r.msg
}

//方形实现了形状
type square struct {
	msg string
}

func (c *square) doDraw() string {
	return c.msg
}

/*
正常情况下，当用户想使用不同的形状doDraw时，将需要调用各自的square/rectangle/circle的doDraw
如果这些实现不想包外可见，隐藏实现。则使用一个facade对其进行封装，对外只提供执行的接口drawCircle/drawSquare/drawRectangle
*/
type shapeFacade struct {
	//facade中维护各个实现的引用
	circle
	square
	rectangle
}

//对外只提供一个facade封装
func NewShapeFacade() shapeFacade {
	return shapeFacade{
		circle:    circle{msg: "this is circle"},
		square:    square{msg: "this is square"},
		rectangle: rectangle{msg: "this is rectangle"},
	}
}

func (s *shapeFacade) drawCircle() string {
	return s.circle.doDraw()
}

func (s *shapeFacade) drawSquare() string {
	return s.square.doDraw()
}

func (s *shapeFacade) drawRectangle() string {
	return s.rectangle.doDraw()
}
