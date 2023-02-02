package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	done := make(chan int)

	numCPU := runtime.NumCPU()

	fmt.Println(numCPU)

	for i := 0; i < numCPU; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}

	time.Sleep(time.Second * 86400)
	close(done)
}
