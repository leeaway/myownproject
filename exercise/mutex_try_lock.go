package exercise

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

/*
	为sync.Mutex实现TryLock
*/

const (
	mutexLocked      = 1 << iota // 加锁标识位置
	mutexWoken                   // 唤醒标识位置
	mutexStarving                // 锁饥饿标识位置
	mutexWaiterShift             //	标识Waiter的起始bit位置
)

type MyMutex struct {
	sync.Mutex
}

func (m *MyMutex) TryLock() bool {
	//如果能成功抢到🔐,使用cas更新为locked
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}

	//如果处于唤醒、加锁或饥饿状态，这次请求就不参与竞争了，直接返回false
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if (old & (mutexLocked | mutexWoken | mutexStarving)) != 0 {
		return false
	}

	//尝试在竞争状态去请求锁
	new := old | mutexLocked //将锁标识位更改
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new)
}
