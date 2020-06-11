package creational

import (
	"testing"
)

func TestNewSimpleFactory(t *testing.T) {
	expect1, expect2 := "impl1:this is type1", "impl2:this is type2"
	type1, type2 := "Impl1", "Impl2"
	instance1 := NewSimpleFactory(type1)
	rst1, _ := instance1.doStore("this is type1")
	instance2 := NewSimpleFactory(type2)
	rst2, _ := instance2.doStore("this is type2")
	if rst1 != expect1 || rst2 != expect2 {
		t.Errorf("failed")
	}
	t.Log("success")
}
