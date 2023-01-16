package formulas

import (
	"fmt"
	"sync"
	"time"
)

func sleepFun(sec time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(sec * time.Second)
	fmt.Println("goroutine exit %s:", sec.String())
}

func SyncOutput() {
	var wg sync.WaitGroup

	wg.Add(1)
	go sleepFun(3, &wg)
	go sleepFun(1, &wg)
	wg.Wait()
	fmt.Println("Main goroutine exit")

}
