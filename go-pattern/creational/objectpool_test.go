package creational

import "testing"

func TestNewObjectPool(t *testing.T) {
	p := NewObjectPool(3)
	select {
	case obj := <-*p:
		obj.doSth()
		//*p <- obj
	default:
		// No more objects left — retry later or fail
		return
	}

	p2 := NewObjectPoolBySyncPool()
	if poolInstance, ok := p2.Get().(PoolInstance); ok {
		poolInstance.doSth()
	}
}
