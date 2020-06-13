package creational

import (
	"testing"
)

func TestNewSimpleFactory(t *testing.T) {
	expect1, expect2 := "lenovo:this is lenovo mouse", "hp:this is hp mouse"
	type1, type2 := "lenovo", "hp"
	instance1 := NewSimpleFactory(type1)
	rst1, _ := instance1.doCreate("this is lenovo mouse")
	instance2 := NewSimpleFactory(type2)
	rst2, _ := instance2.doCreate("this is hp mouse")
	if rst1 != expect1 || rst2 != expect2 {
		t.Errorf("failed")
	}
	t.Log("success")
}
