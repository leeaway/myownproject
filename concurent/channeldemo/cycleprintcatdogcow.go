package channeldemo

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/1/9 20:32
 */

/*
	使用三个协程按顺序打印cat、dog、cow，打印100次
*/
func CyclePrintCatDogCow() {
	wg := sync.WaitGroup{}
	ch := make(chan struct{})
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- struct{}{}
			if i&1 == 0 {
				fmt.Println("cat")
			}
		}
		close(ch)
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			<-ch
			if i&1 > 0 {
				fmt.Println("dog")
			}
		}
		wg.Done()
	}()

	wg.Wait()
}

func CyclePrintCatDogCow2() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	cnt := int32(0)

	go PrintCat(&wg, &cnt)
	go PrintDog(&wg, &cnt)
	go PrintCow(&wg, &cnt)

	wg.Wait()
}

func PrintCow(wg *sync.WaitGroup, cnt *int32) {
	for {
		if *cnt >= 30 {
			wg.Done()
			return
		}
		if *cnt%3 == 2 {
			fmt.Println("cow")
			atomic.AddInt32(cnt, 1)
		}
	}
}

func PrintDog(wg *sync.WaitGroup, cnt *int32) {
	for {
		if *cnt >= 30 {
			wg.Done()
			return
		}
		if *cnt%3 == 1 {
			fmt.Println("dog")
			atomic.AddInt32(cnt, 1)
		}
	}
}

func PrintCat(wg *sync.WaitGroup, cnt *int32) {
	for {
		if *cnt >= 30 {
			wg.Done()
			return
		}
		if *cnt%3 == 0 {
			fmt.Println("cat")
			atomic.AddInt32(cnt, 1)
		}
	}
}
