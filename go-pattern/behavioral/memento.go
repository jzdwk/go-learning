package behavioral

/*
备忘录模式（Memento Pattern）保存一个对象的某个状态，以便在适当的时候恢复对象。备忘录模式属于行为型模式。

使用场景：
	1、需要保存/恢复数据的相关状态场景。 2、提供一个可回滚的操作。
应用实例：
	1、Windows 里的 ctri + z。 2、IE 中的后退。 3、数据库的事务管理。
关键代码：
	1、 Memento:备忘录。主要用来存储原发器对象的内部状态，但是具体需要存储哪些数据是由原发器对象来决定的。
		另外备忘录应该只能由原发器对象来访问它内部的数据，原发器外部的对象不应该能访问到备忘录对象的内部数据。
    2、 Originator:原发器。使用备忘录来保存某个时刻原发器自身的状态，也可以使用备忘录来恢复内部状态。
	3、 Caretaker:备忘录管理者，或者称为备忘录负责人。主要负责保存备忘录对象，但是不能对备忘录对象的内容进行操作或检查

refer to: https://www.runoob.com/design-pattern/memento-pattern.html
		  https://juejin.im/post/5c075105e51d45598b76f4b0

*/

// 备忘录memento，用于记录需要被保存的对象的状态信息
type memento struct {
	state string // 保存的状态,这个字段与originator的state对饮
}

func (m *memento) SetState(s string) {
	m.state = s
}
func (m *memento) GetState() string {
	return m.state
}

// 发起人originator，需要被保存的对象，除了日常操作，需要定义一个将当前状态保存得到memento的方法
type originator struct {
	state string // 要保存的状态，例子中为字符串，实际为对象的状态信息
}

func (o *originator) SetState(s string) {
	o.state = s
}
func (o *originator) GetState() string {
	return o.state
}

// 这里就是规定了要保存的状态，返回memento结构，注意originator对外并不暴露memento，而是将创建过程封装
func (o *originator) CreateMemento() *memento {
	return &memento{state: o.state}
}

// 从备忘录还原状态
func (o *originator) GetFromMemento(mem *memento) {
	o.state = mem.GetState()
}

// 负责人，只提供memento的维护，而没有对内部状态的操作
type caretaker struct {
	memento []*memento
}

func (c *caretaker) GetMemento(index int) *memento {
	return c.memento[index]
}

func (c *caretaker) SetMemento(m *memento) {
	c.memento = append(c.memento, m)
}

func NewOriginator() *originator {
	return &originator{}
}

func NewCaretaker() *caretaker {
	return &caretaker{}
}
