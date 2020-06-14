package structural

import (
	"fmt"
	"testing"
)

func TestFlyweight(t *testing.T) {
	cf := CircleFactory{}
	circle1 := cf.getCircle("green")
	fmt.Println(circle1.doDraw(1, 2))
	circle2 := cf.getCircle("green")
	fmt.Println(circle2.doDraw(2, 3))
	circle3 := cf.getCircle("red")
	fmt.Println(circle3.doDraw(3, 4))
}
