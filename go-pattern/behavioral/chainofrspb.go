package behavioral

import (
	"fmt"
	"strings"
)

/*
顾名思义，责任链模式（Chain of Responsibility Pattern）为请求创建了一个接收者对象的链。属于行为型模式。
职责链上的处理者负责处理请求，客户只需要将请求发送到职责链上即可，无须关心请求的处理细节和请求的传递.
所以职责链将请求的发送者和请求的处理者解耦了。

关键代码：
	Handler里面聚合它自己（即例子中的next方法），所有的实现都继承自Handler接口。
根据链中的每个元素执行自己的业务逻辑，然后递归调用下一个handler。

refer to 	https://github.com/pibigstar/go-demo/blob/master/design/chain/chain.go
			https://www.runoob.com/design-pattern/chain-of-responsibility-pattern.html
*/

//统一的接口定义，handle方法用于处理业务逻辑，next方法引用自身
type Handler interface {
	Handle(content string)
	//关键代码
	next(handler Handler, content string)
}

// 广告过滤
type AdHandler struct {
	nextHandler Handler
}

func (ad *AdHandler) Handle(content string) {
	fmt.Println("执行广告过滤。。。")
	newContent := strings.Replace(content, "广告", "**", 1)
	fmt.Println(newContent)
	//关键代码
	ad.next(ad.nextHandler, newContent)
}

func (ad *AdHandler) next(handler Handler, content string) {
	if handler != nil {
		handler.Handle(content)
	}
}

// 涉黄过滤
type YellowHandler struct {
	nextHandler Handler
}

func (yellow *YellowHandler) Handle(content string) {
	fmt.Println("执行涉黄过滤。。。")
	newContent := strings.Replace(content, "涉黄", "**", 1)
	fmt.Println(newContent)
	yellow.next(yellow.nextHandler, newContent)
}

func (yellow *YellowHandler) next(handler Handler, content string) {
	if handler != nil {
		handler.Handle(content)
	}
}

// 敏感词过滤
type SensitiveHandler struct {
	nextHandler Handler
}

func (sensitive *SensitiveHandler) Handle(content string) {
	fmt.Println("执行敏感词过滤。。。")
	newContent := strings.Replace(content, "敏感词", "***", 1)
	fmt.Println(newContent)
	sensitive.next(sensitive.nextHandler, newContent)
}

func (sensitive *SensitiveHandler) next(handler Handler, content string) {
	if handler != nil {
		handler.Handle(content)
	}
}
