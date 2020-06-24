package behavioral

import "testing"

func TestNewObserver(t *testing.T) {
	obA := NewObserverA()
	obB := NewObserverB()
	sub := NewSubject()
	sub.register(obA).register(obB).register(obA)
	sub.change()

}
