package main

import (
	"TSM-Server/cmd/tmd"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	//设置要接收的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
		os.Exit(1)
	}()

	start := make(chan bool)
	go tmd.Start(start)
	<-start

	res := tmd.Command("version")
	log.Println("res is", res)
	time.Sleep(5 * time.Second)
	tmd.Command("exit")

	time.Sleep(5 * time.Second)
	log.Println("restart")
	d := make(chan bool)
	go tmd.Start(d)
	<-d

	log.Println("现在服务器是否启动", tmd.CheckStart())

	<-done
	fmt.Println("进程被终止")

}
