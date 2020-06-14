package structural

import (
	"fmt"
	"testing"
)

func TestShapeFacade(t *testing.T) {
	facade := NewShapeFacade()
	//调用封装的方法屏蔽实现
	fmt.Println(facade.drawCircle())
	fmt.Println(facade.drawSquare())
	fmt.Println(facade.drawRectangle())
}
