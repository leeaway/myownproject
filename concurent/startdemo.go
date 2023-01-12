package concurent

import (
	"fmt"
	"time"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/1/9 15:59
 */

func StartDemo() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(fmt.Sprintf("goroutine %d print", i))
		}(i)
	}
	fmt.Println("main goroutine print")
	time.Sleep(time.Second)
}
