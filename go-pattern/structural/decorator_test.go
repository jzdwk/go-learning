package structural

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	base := circleDct{msg: " circle "}
	//实现1：和代理proxy不同在于，创建decorator时引用base/other decorator作为入参
	decoratorA := NewShapeDecoratorA(&base)
	fmt.Println(decoratorA.doDraw())
	decoratorB := NewShapeDecoratorB(decoratorA)
	fmt.Println(decoratorB.doDraw())
	//实现2：直接使用func变量
	draw := DecorateDraw(base.doDraw)
	fmt.Println(draw())
}
