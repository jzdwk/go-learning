package creational

/*
构造器模式，将创建一个对象的步骤分离，每一步都是对原对象的补充
*/

//构造器接口，结构中每一个方法用于组成最终对象的一部分
type builder interface {
	addPlugin1(plugin string) builder
	addPlugin2(plugin string) builder
	addPlugin3(plugin string) builder
	print() string
}

//构造器的一个实现，重要的地方为每次构造返回接口本身，这个才可以继续调用后续功能
type builderImpl struct {
	msg string
}

func (b *builderImpl) addPlugin1(plugin string) builder {
	b.msg += plugin
	return b
}

func (b *builderImpl) addPlugin2(plugin string) builder {
	b.msg += plugin
	return b
}

func (b *builderImpl) addPlugin3(plugin string) builder {
	b.msg += plugin
	return b
}
func (b *builderImpl) print() string {
	return b.msg
}

func NewBuilder() builder {
	return &builderImpl{"this is builder"}
}
