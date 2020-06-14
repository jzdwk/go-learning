package structural

import (
	"fmt"
	"sync"
)

/*
享元模式（Flyweight Pattern）主要用于减少创建对象的数量，以减少内存占用和提高性能。
这种类型的设计模式属于结构型模式，它提供了减少对象数量从而改善应用所需的对象结构的方式。
享元模式尝试重用现有的同类对象(color)，如果未找到匹配的对象，则创建新对象。

场景： 1、系统有大量相似对象。 2、需要缓冲池的场景。
优点：
	大大减少对象的创建，降低系统的内存，使效率提高。
缺点：
	提高了系统的复杂度，需要分离出外部状态和内部状态，而且外部状态具有固有化的性质，不应该随着内部状态的变化而变化，否则会造成系统的混乱
实例： db的连接池，string类型设计

*/

//shape接口
type shapeFw interface {
	doDraw(x int, y int) string
}

//circle实现了shape接口，问题在于，当存在相同color的circle时，并不想重新新建，而是直接返回
type circleFw struct {
	color string
}

func (c *circleFw) doDraw(x int, y int) string {
	//除了颜色，根据剩下的x,y的不同，执行各自的draw逻辑
	return fmt.Sprintf("draw circle x= %d, y= %d, color = %v", x, y, c.color)
}

//var circleMap = sync.Map{} //global map to store circle
var circleMap2 = make(map[string]shapeFw)
var muxLock = sync.Mutex{}

//使用一个factory
type CircleFactory struct {
}

//需要考虑线程安全问题
func (c *CircleFactory) getCircle(color string) shapeFw {
	muxLock.Lock()
	defer muxLock.Unlock()
	//如果color已经在cache中，则返回，color就是创建各个对象时，需要执行/共享的相同的逻辑
	if v, ok := circleMap2[color]; ok {
		return v
	}
	fmt.Println("create new circle")
	v := circleFw{
		color: color,
	}
	circleMap2[color] = &v
	return &v
}
