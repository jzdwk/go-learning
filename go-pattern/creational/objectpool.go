package creational

import (
	"fmt"
	"sync"
)

/*
Pool用于存储那些被分配了但是没有被使用，而未来可能会使用的值，以减小垃圾回收的压力，在需要用到的时候直接从pool中取。
*/

//1. 自定义实现
//维护一个channel作为对象的pool，这种实现非线程安全
type Pool chan PoolInstance

//pool中对象的接口
type PoolInstance interface {
	doSth()
}

func NewObjectPool(total int) *Pool {
	p := make(Pool, total)
	for i := 0; i < total; i++ {
		p <- newInstance() //对象入pool，返回pool指针
	}
	return &p
}

type poolImpl struct {
}

func (p *poolImpl) doSth() {
	fmt.Println("do sth.")
}

func newInstance() PoolInstance {
	return &poolImpl{}
}

//2. 使用sync.Pool
//使用sync.Pool创建线程安全的pool
var bytePool = sync.Pool{
	New: func() interface{} {
		instance := poolImpl{} //new instance
		return &instance
	},
}

func NewObjectPoolBySyncPool() *sync.Pool {
	return &bytePool
}
