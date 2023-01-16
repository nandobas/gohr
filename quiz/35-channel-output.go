package quiz

import (
	"fmt"
)

func ChannelOutputDeadLock() {
	ch := make(chan int)
	ch <- 7
	val := <-ch
	fmt.Println(val)
}

func ChannelOutput() {
	ch := make(chan int)
	go func() {
		ch <- 7
	}()

	select {
	case val := <-ch:
		fmt.Println(val)
	}
}
