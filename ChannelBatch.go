package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // 模拟任务处理
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job
	}
}

func main() {
	numJobs := 10
	numWorkers := 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动多个协程
	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, results)
	}

	// 发送任务到通道
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs) // 关闭任务通道

	// 获取任务结果
	for i := 1; i <= numJobs; i++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
}
