package concurrency

import (
	"fmt"
	"time"
)

func TryForSelect() {
	go func() {
		for {
			select {
			default:
				fmt.Println("Doing some work")
			}
		}
	}()

	time.Sleep(5 * time.Second)
}
