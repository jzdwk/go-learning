package behavioral

import "testing"

func TestState(t *testing.T) {
	ct := NewMyContext()
	ct.doRun()
	ct.doRun()
	ct.doStop()
}
