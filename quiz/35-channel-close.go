package quiz

import (
	"fmt"
)

func ChannelOutputClose() {
	ch := make(chan int)
	close(ch)
	val := <-ch
	fmt.Println(val)
}
