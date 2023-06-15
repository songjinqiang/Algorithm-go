package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
