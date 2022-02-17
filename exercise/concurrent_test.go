package exercise

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func Test(t *testing.T) {
	n := 7
	fmt.Println(n >> 3)
	fmt.Println(n & 1)
}

func TestMyMutex_TryLock(t *testing.T) {
	Convey("TestMyMutex_TryLock", t, func() {
		So(try(0), ShouldEqual, true)
		So(try(3), ShouldEqual, false)
	})
}

func try(sleepSeconds int) bool {
	var mu MyMutex

	go func() {
		mu.Lock()
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
		mu.Unlock()
	}()

	time.Sleep(time.Second)

	if mu.TryLock() {
		fmt.Println("got the lock")
		mu.Unlock()
		return true
	}

	fmt.Println("can not get the lock")
	return false
}

func TestCount(t *testing.T) {
	var mu MyMutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			mu.Unlock()
		}()
	}
	fmt.Printf("Total: %v, Locked: %v, Woken: %v, Starving: %v", mu.Count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}
