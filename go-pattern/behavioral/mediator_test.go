package behavioral

import "testing"

func TestMediator(t *testing.T) {
	//创建一个明确的中介者
	mediator := NewMyMediator()
	//创建两个通信组件
	colleagueA := NewColleagueA(mediator)
	colleagueB := NewColleagueB(mediator)
	//将通信组件注册至mediator，类比beego/springmvc的http handler处理
	mediator.register("Colleague A", colleagueA)
	mediator.register("Colleague B", colleagueB)

	//A发送信息
	colleagueA.sendTo("Colleague B", "Hello world!")

}
