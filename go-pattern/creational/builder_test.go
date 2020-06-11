package creational

import "testing"

func TestNewBuilder(t *testing.T) {
	expect := "this is builder with plugin1 with plugin2"
	b := NewBuilder()
	b.addPlugin1(" with plugin1").addPlugin2(" with plugin2")
	if b.print() != expect {
		t.Error("failed")
	}
	t.Log("success")
}
