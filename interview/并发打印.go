package interview

import (
	"fmt"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/1/16 19:33
 */

func ConcurrentPrint123() {
	ch1, ch2, ch3 := make(chan struct{}, 1), make(chan struct{}, 1), make(chan struct{}, 1)
	ch1 <- struct{}{}
	for i := 0; i < 1000; i++ {
		go Print1(ch1, ch2)
		go Print2(ch2, ch3)
		go Print3(ch3, ch1)
	}
}

func Print1(ch1, ch2 chan struct{}) {
	<-ch1
	fmt.Println(1)
	ch2 <- struct{}{}
}

func Print2(ch2, ch3 chan struct{}) {
	<-ch2
	fmt.Println(2)
	ch3 <- struct{}{}
}

func Print3(ch3, ch1 chan struct{}) {
	<-ch3
	fmt.Println(3)
	ch1 <- struct{}{}
}
