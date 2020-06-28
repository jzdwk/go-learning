package behavioral

import "testing"

func TestVisitor(t *testing.T) {
	c := NewComputer()
	//使用buyer去访问computer
	c.accept(NewComputerBuyer())
	//使用maintainer去访问computer
	c.accept(NewComputerMaintainer())
}
