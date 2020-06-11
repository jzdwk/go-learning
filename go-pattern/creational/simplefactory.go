package creational

//简单工厂模式

//顶层接口
type store interface {
	doStore(msg string) (out string, err error)
}

//实现接口的结构1
type storeImpl1 struct {
	msg string
}

func (imp1 *storeImpl1) doStore(msg string) (string, error) {
	return imp1.msg + msg, nil
}

//实现接口的结构2
type storeImpl2 struct {
	msg string
}

func (imp2 *storeImpl2) doStore(msg string) (string, error) {
	return imp2.msg + msg, nil
}

func NewSimpleFactory(storeType string) store {
	switch storeType {
	case "Impl1":
		return &storeImpl1{"impl1:"}
	case "Impl2":
		return &storeImpl2{"impl2:"}
	default:
		return nil
	}
}
