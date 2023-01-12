package channeldemo

import "fmt"

/**
 * @author 2416144794@qq.com
 * @date 2023/1/9 17:53
 */

func ChannelStartDemo() {
	//初始化一个类型为int，缓存区大小为1的channel
	ch := make(chan int, 1)
	defer func() {
		close(ch)
	}()
	//向channel发送一个2
	ch <- 2
	//ans 接受channel中的值
	ans, ok := <-ch
	if ok {
		fmt.Println(ans)
	}
}
