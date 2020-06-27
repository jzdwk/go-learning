package behavioral

import "testing"

func TestTemplate(t *testing.T) {
	//这里是一个互相引用，因此需要3步创建
	tp := NewBaseTp(nil)
	subOne := NewSubOneTp(tp)
	subTwo := NewSubTwoTp(tp)

	tp.setTp(subOne)
	tp.doTp()
	tp.setTp(subTwo)
	tp.doTp()
}
