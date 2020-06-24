package behavioral

import "fmt"

/*
当对象间存在一对多关系时，且一侧的对象需要根据另一侧对象的行为去执行业务逻辑(监听)，则使用观察者模式（Observer Pattern）。
比如，当一个对象被修改时，则会自动通知依赖它的对象。观察者模式属于行为型模式。

关键代码：
	通过一个集合数据域来存放观察者们。
*/

//定义观察者接口，subject将会维护一个observer的集合
type observer interface {
	//当监听到sub的变化，需要doSth
	doSth()
}

// 具体的观察者定义
type observerA struct {
	msg string
}

func (o *observerA) doSth() {
	fmt.Println("This is observer " + o.msg)
}

type observerB struct {
	msg string
}

func (o *observerB) doSth() {
	fmt.Println("This is observer " + o.msg)
}

type subject struct {
	//使用一个map充当set，即注册的观察者列表
	observers map[observer]struct{}
}

//当发生一些状态改变/事件时，触发notify
func (s *subject) change() {
	fmt.Println("sth changed")
	s.notifyAll()
}

//将新的观察者加入列表，此处返回自身是为了测试
func (s *subject) register(ob observer) *subject {
	if _, ok := s.observers[ob]; !ok {
		s.observers[ob] = struct{}{}
	}
	return s
}

//将新的观察者加入列表
func (s *subject) deRegister(ob observer) {
	if _, ok := s.observers[ob]; ok {
		delete(s.observers, ob)
	}
}
func (s *subject) notifyAll() {
	for k := range s.observers {
		k.doSth()
	}
}

//new
func NewObserverA() observer {
	return &observerA{msg: "A"}
}
func NewObserverB() observer {
	return &observerB{msg: "B"}
}
func NewSubject() *subject {
	return &subject{observers: make(map[observer]struct{})}
}

//官方实现
/*type (
	// Event defines an indication of a point-in-time occurrence.
	Event struct {
		// Data in this case is a simple int, but the actual
		// implementation would depend on the application.
		Data int64
	}

	// Observer defines a standard interface for instances that wish to list for
	// the occurrence of a specific event.
	Observer interface {
		// OnNotify allows an event to be "published" to interface implementations.
		// In the "real world", error handling would likely be implemented.
		OnNotify(Event)
	}

	// Notifier is the instance being observed. Publisher is perhaps another decent
	// name, but naming things is hard.
	Notifier interface {
		// Register allows an instance to register itself to listen/observe
		// events.
		Register(Observer)
		// Deregister allows an instance to remove itself from the collection
		// of observers/listeners.
		Deregister(Observer)
		// Notify publishes new events to listeners. The method is not
		// absolutely necessary, as each implementation could define this itself
		// without losing functionality.
		Notify(Event)
	}
)

type (
	eventObserver struct{
		id int
	}

	eventNotifier struct{
		// Using a map with an empty struct allows us to keep the observers
		// unique while still keeping memory usage relatively low.
		observers map[Observer]struct{}
	}
)

func (o *eventObserver) OnNotify(e Event) {
	fmt.Printf("*** Observer %d received: %d\n", o.id, e.Data)
}

func (o *eventNotifier) Register(l Observer) {
	o.observers[l] = struct{}{}
}

func (o *eventNotifier) Deregister(l Observer) {
	delete(o.observers, l)
}

func (p *eventNotifier) Notify(e Event) {
	for o := range p.observers {
		o.OnNotify(e)
	}
}
*/
