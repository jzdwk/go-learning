package structural

/*
在代理模式（Proxy Pattern）中，一个类代表另一个类的功能。这种类型的设计模式属于结构型模式。
当直接访问对象时，如果存在问题，比如说：
1. 要访问的对象在远程的机器上。
2. 在面向对象系统中，有些对象由于某些原因（比如对象创建开销很大，或者某些操作需要安全控制，或者需要进程外的访问），直接访问会给使用者或者系统结构带来很多麻烦。
我们可以在访问此对象时加上一个对此对象的访问层。
3. 进行一些和访问实际对象无关的权限控制、过滤、LB等方案

refer to：https://www.runoob.com/design-pattern/proxy-pattern.html
		  https://zhuanlan.zhihu.com/p/26141688
*/

//访问对象接口，代理实例和被代理实例均要实现
type subject interface {
	doSth() string
}

//代理subject
type proxySubject struct {
	msg     string
	subject //代理中维护一个真实sub，执行完代理逻辑后调用真实sub的doSth
}

func (s *proxySubject) doSth() string {
	//代理逻辑为s.msg，真实sub的业务逻辑为s.subject.doSth
	return s.msg + s.subject.doSth()
}

//真实sub
type realSubject struct {
	msg string
}

func (s *realSubject) doSth() string {
	return s.msg
}
