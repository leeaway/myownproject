package concurent

import "fmt"

/**
 * @author 2416144794@qq.com
 * @date 2023/1/9 15:46
 */

func ConcurrentPrint0To100(goRoutineNum int) {
	if goRoutineNum < 2 {
		goRoutineNum = 2
	}
	var chanSlice []chan int
	quitChan := make(chan int)
	for i := 0; i < goRoutineNum; i++ {
		chanSlice = append(chanSlice, make(chan int, 1))
	}

	res := 1
	j := 0
	for i := 0; i < goRoutineNum; i++ {
		go func(i int) {
			for {
				<-chanSlice[i]
				if res > 100 {
					quitChan <- 1
					break
				}

				fmt.Println(fmt.Sprintf("goRoutine: %d, print %d", i, res))
				res++

				if j == goRoutineNum-1 {
					j = 0
				} else {
					j++
				}

				chanSlice[j] <- 1
			}
		}(i)
	}
	chanSlice[0] <- 1

	select {
	case <-quitChan:
		fmt.Println("end")
	}
}
