package concurrency

import "fmt"

func sliceToChan(nums []int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range nums {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func getSquare(inC <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		for v := range inC {
			ch <- v * v
		}
		close(ch)
	}()
	return ch
}

func TryPipeline() {
	nums := []int{1, 2, 3, 4, 5}
	dataCh := sliceToChan(nums)
	finalCh := getSquare(dataCh)
	for v := range finalCh {
		fmt.Println(v)
	}
}
