package creational

import (
	"reflect"
	"testing"
)

func TestSingleton(t *testing.T) {

	/*actual1 := NewBySyncOnce()
	actual2 := NewBySyncMutex()
	if !reflect.DeepEqual(actual2, actual1) {
		t.Errorf("failed")
	}
	t.Log("success")*/
	ins := NewBySyncOnce()
	ins2 := NewBySyncOnce()
	ins3 := &singletonSub{"some sub properties"}
	ins4 := &singleton{"some properties", ins3}
	if !reflect.DeepEqual(ins, ins2) {
		t.Errorf("1:ins failed")
	}
	if !reflect.DeepEqual(ins, ins4) {
		t.Errorf("2:ins failed")
	}

	if !reflect.DeepEqual(ins.singletonSub, ins2.singletonSub) {
		t.Errorf("3:ins sub failed")
	}
	if !reflect.DeepEqual(ins.singletonSub, ins4.singletonSub) {
		t.Errorf("4:ins sub failed")
	}

	if ins == ins2 {
		t.Errorf("5:ins failed")
	}
	if ins != ins4 {
		t.Errorf("6:ins failed")
	}
	if ins.singletonSub != ins3 {
		t.Errorf("7:ins failed")
	}
	t.Log("success")
}
