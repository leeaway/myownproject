package exercise

import (
	"fmt"
	"sync"
	"time"
)

/*
	@author: Away Lee (2416144794@qq.com)
	这里count++非原子操作，存在并发问题，每次执行结果应该都不一样
	使用 go run -race main.go 可以检测出数据竞争问题
*/
func ConcurrentCountWithoutMutex() {
	var count = 0
	//使用WaitGroup等待10个goroutine完成，类似于java的countdownlatch
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func ConcurrentCountWithMutex() {
	var count = 0
	var mu sync.Mutex
	//使用WaitGroup等待10个goroutine完成，类似于java的countdownlatch
	var wg sync.WaitGroup
	wg.Add(10)
	beginTime := time.Now().Unix()
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			//此处应放在for循环外，减少加减锁的次数
			mu.Lock()
			for j := 0; j < 10000000; j++ {
				count++
			}
			mu.Unlock()
		}()
	}
	wg.Wait()
	endTime := time.Now().Unix()
	fmt.Printf("consume: %v\n", endTime-beginTime)
	fmt.Println(count)
}
