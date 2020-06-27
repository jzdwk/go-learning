package main

import (
	"context"
	"log"
	_ "net/http/pprof"
	"time"
)

/*
官方说明：
Package context defines the Context type,
which carries deadlines, cancelation signals,
and other request-scoped values across API boundaries and between processes.

通过context，我们可以方便地对同一个请求所产生地goroutine进行约束管理，
可以设定超时、deadline，甚至是取消这个请求相关的所有goroutine。

注意：
1. 不要把context存储在结构体中，而是要显式地进行传递
2. 把context作为第一个参数，并且一般都把变量命名为ctx
3. 就算是程序允许，也不要传入一个nil的context，如果不知道是否要用context的话，用context.TODO()
4. context.WithValue()只用来传递请求范围的值，不要用它来传递可选参数
5. 就算是被多个不同的goroutine使用，context也是安全的

refer to:
	http://www.nljb.net/default/Golang%E4%B9%8BContext%E7%9A%84%E4%BD%BF%E7%94%A8/
	https://juejin.im/post/5a6873fef265da3e317e55b6

*/
func main() {

	//WithCancel
	doWithCancel()

	//WithTimeout定义了一个超时的ctx，到时自动退出
	//ctx2,cancel := context.WithTimeout(context.Background(),3*time.Second)

}

func doWithCancel() {
	//定义一个WithCancel函数，该函数返回了一个context的副本以及一个cancel函数，这个副本有一个Done的channel。
	// 当函数被调用时，Done将close
	ctx, cancel := context.WithCancel(context.Background())

	/*	go func() {
		time.Sleep(3 * time.Second)
		//cancel的一般用法为在defer里定义，当主协程退出后，依次退出各个子协程
		cancel()
	}()*/
	defer cancel()
	//将ctx作为参数给函数A，注意ctx为参数表的第一个参数
	log.Println(withCancelA(ctx))
	time.Sleep(3 * time.Second)
	//cancel的一般用法为在defer里定义，当主协程退出后，依次退出各个子协程

}
func withCancelB(ctx context.Context) string {
	select {
	case <-ctx.Done():
		return "B Done"
	default:
		return "B stop"
	}
}

func withCancelA(ctx context.Context) string {
	go log.Println(withCancelB(ctx))
	select {
	case <-ctx.Done():
		return "A Done"
	default:
		return "A stop"
	}
}
