package structural

/*
桥接（Bridge）是用于把抽象化与实现化解耦，使得二者可以独立变化。
这种类型的设计模式属于结构型模式，它通过提供抽象化和实现化之间的桥接结构，来实现二者的解耦。

通俗来说，就是将实体（类）中变化的部分独立出来，应用“组合”来完成实体的描述。
举个例子，当描述一个图形，可以从形状和颜色两个维度区分，形状有圆、方，颜色有红、黄。
图形的draw方法用于打印图像。如果使用单一继承，则需要4个实现；如果再增加一种红色，又多需要两种。

使用场景：
	 一个类存在两个独立变化的维度，且这两个维度都需要进行扩展。
关键代码：
	将变化的部分抽象（首相并不一定针对具体实物，可能是抽象的概念，比如颜色）
*/

//shapeBg 形状接口
type shapeBg interface {
	draw() string
}

//使用组合，将颜色与形状组合
type squareBg struct {
	msg   string
	color color
}

func (s *squareBg) draw() string {
	return s.msg + s.color.draw()
}

type circleBg struct {
	msg   string
	color color
}

func (s *circleBg) draw() string {
	return s.msg + s.color.draw()
}

//颜色接口，将颜色独立抽象
type color interface {
	draw() string
}
type redBg struct {
	msg string
}

func (s *redBg) draw() string {
	return s.msg
}

type yellowBg struct {
	msg string
}

func (s *yellowBg) draw() string {
	return s.msg
}
