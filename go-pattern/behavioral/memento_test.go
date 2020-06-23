package behavioral

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	org := NewOriginator()
	taker := NewCaretaker()
	org.SetState("state A")
	//将A状态保存到备忘录
	taker.SetMemento(org.CreateMemento())
	org.SetState("state B")
	//B状态保存
	taker.SetMemento(org.CreateMemento())
	org.SetState("state C")

	fmt.Println("current state: ", org.GetState())
	org.GetFromMemento(taker.GetMemento(1))
	fmt.Println("current state: ", org.GetState())
	org.GetFromMemento(taker.GetMemento(0))
	fmt.Println("current state: ", org.GetState())
}
