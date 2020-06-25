package structural

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	//根据规则需要，创建不同的sub实现
	realSub := &realSubject{msg: "real sub"}
	//直接创建代理
	proxySub := &proxySubject{msg: "proxy:"}
	//proxy:real sub
	//和装饰者的差别在于,proxy中不暴露被代理的sub参数，即对外无感知
	doTest(proxySub)
	//real sub
	doTest(realSub)
	//使用函数变量
	proxy := ProxyDoSth(realSub.doSth)
	fmt.Println(proxy())
}

//使用代理后，所有调用的接口的入参定义sub接口即可
func doTest(sub subject) {
	fmt.Println(sub.doSth())
}
