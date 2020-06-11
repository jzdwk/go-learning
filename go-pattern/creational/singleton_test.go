package creational

import (
	"reflect"
	"testing"
)

func TestSingleton(t *testing.T) {

	actual1 := NewBySyncOnce()
	actual2 := NewBySyncMutex()
	if !reflect.DeepEqual(actual2, actual1) {
		t.Errorf("failed")
	}
	t.Log("success")
}
