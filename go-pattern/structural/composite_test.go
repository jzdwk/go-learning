package structural

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {
	top := NewMyOrg("top")
	sub1 := NewMyOrg("sub1")
	sub2 := NewMyOrg("sub2")
	top.addSub(sub1)
	top.addSub(sub2)
	fmt.Println(top.print())
	top.delSub("sub1")
	fmt.Println(top.print())
}
