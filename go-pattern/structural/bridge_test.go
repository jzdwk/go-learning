package structural

import (
	"fmt"
	"testing"
)

func TestBridge(t *testing.T) {
	red := &redBg{msg: " red color "}
	shape := circleBg{color: red, msg: "circle"}
	fmt.Println(shape.draw())
}
