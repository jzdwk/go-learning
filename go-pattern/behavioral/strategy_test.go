package behavioral

import "testing"

func TestStrategy(t *testing.T) {
	//方法1，仅使用context，此时，调用者同样需要根据需求调用strategyA、B,同样出现if..else，治标不治本
	ctx := NewContextSt()
	stgA := NewStrategyA()
	ctx.setStg(stgA)
	ctx.doAction()
	stgB := NewStrategyB()
	ctx.setStg(stgB)
	ctx.doAction()
	//方法2，使用context with map
	ctx2 := NewContextStgWithMap()
	ctx2.setStg(stgA)
	ctx2.setStg(stgB)
	//根据key来map到具体的strategy
	ctx2.doAction("A")
	//方法3， 直接func变量的context，此种方法需要保证各个strategy的方法签名相同，且strategy逻辑单一
	ctx3 := NewContextStgWithFunc()
	ctx3.doAction("B")
}
