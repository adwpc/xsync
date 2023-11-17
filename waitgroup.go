package xsync

import (
	"sync"
	"time"
)

type XWaitGroup struct {
	sync.Mutex
	wg      sync.WaitGroup
	counter int
}

func (swg *XWaitGroup) Add(delta int) {
	swg.Lock()
	defer swg.Unlock()
	swg.counter += delta
	swg.wg.Add(delta)
}

func (swg *XWaitGroup) Done() {
	swg.Lock()
	defer swg.Unlock()
	if swg.counter > 0 {
		swg.counter--
		swg.wg.Done()
	}
}

func (swg *XWaitGroup) Wait(timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		swg.wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
