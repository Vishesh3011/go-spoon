package concurrency

import "fmt"

func TrySelect() {
	myChannel := make(chan string)
	myOtherChannel := make(chan string)

	go func() {
		myChannel <- "Hey from my channel!"
	}()

	go func() {
		myOtherChannel <- "Hey from my other channel!"
	}()
	//msg := <-myOtherChannel
	//fmt.Println(msg)
	//msg = <-myChannel
	//fmt.Println(msg)

	select {
	case msg := <-myChannel:
		fmt.Println(msg)
	case msg := <-myOtherChannel:
		fmt.Println(msg)
	}
}
