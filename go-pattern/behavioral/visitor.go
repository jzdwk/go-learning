package behavioral

import "fmt"

/*
在访问者模式（Visitor Pattern）中，我们使用了一个访问者类，它改变了元素类的执行算法。
通过这种方式，元素的执行算法可以随着访问者改变而改变。
这种类型的设计模式属于行为型模式。

使用场景：
	对于一个数据结构，结构中的各个字段/数据域是固定的（比如电脑，有型号和价格）;
	但是对于不同的使用者，它对于数据的访问策略是不一样的，比如购买者只关心价格，而维修者关心型号；
	传统的对于类/struct的结构为，除了定义类的数据，也实现了buyer/maintainer的业务，这样加入增加一个生产者，则需要修改类；
	因此，将数据的访问方式/策略，和数据的定义本身抽离；

关键代码:
	定义visitor接口以及访问方法，各个visitor实现自己的访问策略；
	数据对象定义一个accept方法，根据不同的visitor调用各自的访问策略
*/

type computer struct {
	types string
	price string
}

func (c *computer) accept(cv computerVisitor) {
	cv.visit(c)
}

//电脑的访问者
type computerVisitor interface {
	visit(c *computer)
}

//定义一个电脑的购买者，只关心电脑的价格
type computerBuyer struct {
}

func (c *computerBuyer) visit(com *computer) {
	fmt.Println("computer buyer do visit " + com.price)
}

//定义一个电脑的维修这，它根据电脑的型号做自身业务
type computerMaintainer struct {
}

func (c *computerMaintainer) visit(com *computer) {
	fmt.Println("computer maintainer do visit " + com.types)
}

func NewComputerBuyer() computerVisitor {
	return &computerBuyer{}
}
func NewComputerMaintainer() computerVisitor {
	return &computerMaintainer{}
}
func NewComputer() *computer {
	return &computer{types: "Mac", price: "9999"}
}
