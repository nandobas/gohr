package lib

import "fmt"

type Message struct {
	ID      string
	Message string
}

type List struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value Message
	Next  *Node
}

func (l *List) Append(m Message) {
	node := &Node{Value: m}

	if l.Head == nil {
		l.Head = node
	}
	if l.Tail != nil {
		l.Tail.Next = node
	}
	l.Tail = node
}
func (l *List) Display() {
	node := l.Head
	for node != nil {
		fmt.Println(node.Value.Message)
		node = node.Next
	}
}

func ExecLinkedListTest() {
	var l List
	l.Append(Message{
		ID:      "1",
		Message: "Hello",
	})
	l.Append(Message{
		ID:      "2",
		Message: "World",
	})
	l.Append(Message{
		ID:      "3",
		Message: "By by!",
	})

	l.Display()
}
