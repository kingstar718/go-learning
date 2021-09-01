package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/**
有时我们希望我们的 Go 程序能够智能地处理Unix 信号。
例如，我们可能希望服务器在收到 . 时正常关闭SIGTERM，或者命令行工具在收到SIGINT.
这是在 Go 中使用通道处理信号的方法
*/

func main() {
	// Go 信号通知通过os.Signal 在通道上发送值来工作。
	//我们将创建一个通道来接收这些通知（我们还将创建一个通道以在程序可以退出时通知我们）。
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	//signal.Notify 注册给定的频道以接收指定信号的通知
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//这个 goroutine 对信号执行阻塞接收。当它得到一个时，它会打印出来，然后通知程序它可以完成。
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// 程序将在此处等待，直到它获得预期的信号（如上面的 goroutine 所指示的，在 上发送一个值done）然后退出。
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

	// 当我们运行这个程序时，它会阻塞等待信号。通过键入ctrl-C（终端显示为^C），我们可以发送一个SIGINT信号，使程序打印interrupt然后退出。
}
