package entities

import "fmt"

type Message struct {
	msg string
	n   int
}
type Result struct {
	rc  int
	msg string
}

func worker(m Message) Result {
	var r Result
	r.msg = m.msg
	r.rc = 1
	return r
}

func WorkOutput() {
	var m Message
	ch := make(chan Result)

	fmt.Println("Start a goroutine that will call worker.")

	/* go func() {
		r := worker(m)
		ch <- r
	} */ // function must be invoked in go statement

	/* go func() {
		r := worker(m)
		r -> ch // expected operand, found '>' (and 2 more errors)
	} () */

	go func() {
		r := worker(m)
		ch <- r
	}()

	// go ch <- worker(m) // function must be invoked in go statement (and 1 more errors)

	fmt.Println("Done.")
}
