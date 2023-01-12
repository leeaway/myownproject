package channeldemo

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/1/9 17:49
 */

func Test_(t *testing.T) {
	Convey("", t, func() {
		Convey(" test1", func() {
			ChannelStartDemo()
		})
	})
}

/*
	1. 向一个nil channel发送消息，会死锁。"all goroutines are asleep - deadlock!"；
	2. 向已经close的channel发送消息，会panic；
	3. channel关闭后不可以再发送消息，但可以继续接受消息；
	4. 当channel关闭，且缓冲区为空时，继续从channel接受会获取到相应的0值；
*/
func Test_SendToNilChannel(t *testing.T) {
	Convey("SendToNilChannel", t, func() {
		//Convey("SendToNilChannel test1", func() {
		//	var ch chan int
		//	ch <- 2
		//})

		//Convey("channel已经关闭", func() {
		//	ch := make(chan int)
		//	close(ch)
		//	ch <- 2
		//})

		//Convey("channel已经关闭，继续接受", func() {
		//	ch := make(chan int, 2)
		//	ch <- 1
		//	ch <- 2
		//	num1, ok := <-ch
		//	if ok {
		//		fmt.Println(num1)
		//	}
		//	close(ch)
		//	num2, ok := <-ch
		//	if ok {
		//		fmt.Println(num2)
		//	}
		//})

		Convey("channel已经关闭，且缓冲区为空，继续接受", func() {
			ch := make(chan int, 2)
			ch <- 1
			ch <- 2
			num1, ok := <-ch
			if ok {
				fmt.Println(num1)
			}
			num2, ok := <-ch
			if ok {
				fmt.Println(num2)
			}
			close(ch)
			empty, _ := <-ch
			fmt.Println(empty)

		})
	})
}

func Test_ForRangeChannel(t *testing.T) {
	Convey("ForRangeChannel", t, func() {
		Convey("ForRangeChannel test1", func() {
			ch := make(chan int, 4)
			for i := 0; i < cap(ch); i++ {
				ch <- i
			}
			close(ch)
			for i := range ch {
				fmt.Println(i)
			}
		})
	})
}

func Test_UnbufferedChannel(t *testing.T) {
	Convey("UnbufferedChannel", t, func() {
		Convey("UnbufferedChannel test1", func() {
			ch1 := make(chan int)
			defer func() {
				close(ch1)
			}()
			go func() {
				fmt.Println(<-ch1)
			}()
			ch1 <- 2
		})
	})
}

func Test_ForSelectChannel(t *testing.T) {
	Convey("ForSelectChannel", t, func() {
		Convey("ForSelectChannel test1", func() {
			ch := make(chan int, 2)
			for i := 0; i < cap(ch); i++ {
				ch <- i + 1
			}
			close(ch)

			ch2 := make(chan int, 2)
			for i := 0; i < cap(ch2); i++ {
				ch2 <- i + 1
			}
			close(ch2)

			chClosed, ch2Closed := false, false
			for {
				if chClosed && ch2Closed {
					return
				}
				select {
				case i, ok := <-ch:
					if ok {
						fmt.Println(i)
					} else {
						chClosed = true
						fmt.Println("ch closed")
					}

				case i, ok := <-ch2:
					if ok {
						fmt.Println(i)
					} else {
						ch2Closed = true
						fmt.Println("ch2 closed")
					}
				default:
					fmt.Println("waiting")
				}
			}
		})
	})
}

func Test_CyclePrintCatDogCow(t *testing.T) {
	Convey("CyclePrintCatDogCow", t, func() {
		Convey("CyclePrintCatDogCow test1", func() {
			CyclePrintCatDogCow2()
		})
	})
}
