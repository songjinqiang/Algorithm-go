package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Producer sending:", i)
		ch <- i                 // 发送数据到通道
		time.Sleep(time.Second) // 模拟生产过程
	}
	close(ch) // 关闭通道
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Println("Consumer received:", num)
		time.Sleep(3 * time.Second) // 模拟消费过程
	}
}

func main() {
	ch := make(chan int) // 创建一个通道

	go producer(ch) // 启动生产者 Goroutine
	go consumer(ch) // 启动消费者 Goroutine

	time.Sleep(16 * time.Second) // 等待一段时间，模拟程序运行

	fmt.Println("Done")
}
