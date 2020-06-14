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
例子：
	MVC框架，controller/service/dao层均向上层提供了封装

refer to: https://www.runoob.com/design-pattern/facade-pattern.html
*/

//定义一个形状接口
type shapeFcd interface {
	doDraw() string
}

//圆形实现了形状
type circleFcd struct {
	msg string
}

func (c *circleFcd) doDraw() string {
	return c.msg
}

//三角形实现了形状
type rectangleFcd struct {
	msg string
}

func (r *rectangleFcd) doDraw() string {
	return r.msg
}

//方形实现了形状
type squareFcd struct {
	msg string
}

func (c *squareFcd) doDraw() string {
	return c.msg
}

/*
正常情况下，当用户想使用不同的形状doDraw时，将需要调用各自的square/rectangleFcd/circle的doDraw
如果这些实现不想包外可见，隐藏实现。则使用一个facade对其进行封装，对外只提供执行的接口drawCircle/drawSquare/drawRectangle
*/
type shapeFacade struct {
	//facade中维护各个实现的引用
	circleFcd
	squareFcd
	rectangleFcd
}

//对外只提供一个facade封装
func NewShapeFacade() shapeFacade {
	return shapeFacade{
		circleFcd:    circleFcd{msg: "this is circleFcd"},
		squareFcd:    squareFcd{msg: "this is squareFcd"},
		rectangleFcd: rectangleFcd{msg: "this is rectangleFcd"},
	}
}

func (s *shapeFacade) drawCircle() string {
	return s.circleFcd.doDraw()
}

func (s *shapeFacade) drawSquare() string {
	return s.squareFcd.doDraw()
}

func (s *shapeFacade) drawRectangle() string {
	return s.rectangleFcd.doDraw()
}
