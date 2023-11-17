package xsync

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	var swg XWaitGroup
	swg.Add(1)
	go func() {
		time.Sleep(5 * time.Second) // simulate a long task
		swg.Done()
	}()
	// swg.Done() // extra Done call, will not cause panic
	if swg.Wait(3 * time.Second) {
		fmt.Println("Timed out")
	} else {
		fmt.Println("XWaitGroup finished")
	}
}
