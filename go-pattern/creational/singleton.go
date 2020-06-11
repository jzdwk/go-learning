package creational

import (
	"sync"
	"sync/atomic"
)

type singleton struct {
	properties interface{} //对象属性
}

var (
	once     sync.Once
	instance *singleton //private global var
)

//单例模式常见于初始化db连接池等操作，被创建的对象全局唯一，单例模式的主要问题在于线程安全，常见的实现方式有：
//1. sync.once，属于懒汉模式
func NewBySyncOnce() *singleton {
	// use sync.once
	once.Do(func() {
		instance = &singleton{"some properties"} //init
	})
	return instance
}

//2.sync.Mutex，通过互斥锁的懒汉模式
var mu sync.Mutex

func NewBySyncMutex() *singleton {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &singleton{"some properties"}
	}
	return instance
}

//3.double-check+atomic
var initialized uint32

func NewByDoubleAtomic() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}
