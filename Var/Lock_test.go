package Var

import (
	"sync"
	"testing"
)

/**
非安全的
*/
func TestWithoutLock(t *testing.T) {
	counter := 1
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}
	t.Log(counter)
}

/**
并发编程 使用mutex,很粗暴
*/
func TestLock(t *testing.T) {
	var mutex sync.Mutex
	counter := 1
	for i := 0; i < 1000; i++ {
		go func() {
			defer func() {
				//解锁
				mutex.Unlock()
			}()
			//很暴力,lock之后,其他协程到这里就会阻塞
			mutex.Lock()
			counter++
		}()
	}
	t.Log(counter)
}

/**
waitgroup
*/
func TestWaitgroup(t *testing.T) {
	var wg sync.WaitGroup
	var lock sync.Mutex
	counter := 1

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				lock.Unlock()
			}()
			lock.Lock()
			counter++
		}()
	}
	wg.Wait()
	t.Log(counter)
}
