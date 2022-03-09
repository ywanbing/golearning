package main

import (
	"context"
	"runtime"
	"time"
)

func main() {
	timer := time.NewTimer(60 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	cpuNum := runtime.NumCPU()

	for i := 0; i < cpuNum; i++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					break
				default:
					for n := 0; n < 30_000_000; n++ {
					}
					time.Sleep(10 * time.Millisecond)
				}
			}
		}()
	}

	select {
	case <-timer.C:
		cancel()
	}

	timer.Stop()
}
