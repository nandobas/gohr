package formulas

import "fmt"

// /* Implement a stack – in any language of preference – that satisfies the following conditions:
// - Provide implementations for:
//     - push(): inserts an element in the stack
//     - pop(): removes an element from the stack
//     - peek(): peeks the last element added in the stack
//     - sum(): sums the element in the stack
// - All operations must have a complexity of O(1).
// - A working program is not required, only logic will be taken into account
// */

type Stack struct {
	sumItems int
	top      int
	list     []int
}

func (s *Stack) Push(x int) error {
	s.list = append(s.list, x)
	s.top = x
	s.sumItems += x
	return nil
}

func (s *Stack) Pop() (int, error) {
	tempList := s.list[0 : len(s.list)-1] // todos os itens menos o ultimo
	decrement := s.list[len(s.list)-1]
	s.list = tempList
	if len(tempList) > 0 {
		s.top = tempList[len(tempList)] //recebe o q estava na ultima posição
	}

	s.sumItems -= decrement
	return s.top, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.list) > 0 {
		return s.top, nil
	}
	return 0, fmt.Errorf("peek: cannot peek, list empty")
}

func (s *Stack) Sum() int {
	return s.sumItems
}

func StackExecute() {
	stack1 := Stack{
		sumItems: 0,
		top:      0,
		list:     []int{},
	}

	err := stack1.Push(1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Push:", 1)

	givenVal, err := stack1.Peek()
	if err != nil {
		panic(err)
	}
	fmt.Println("Peek:", givenVal)

	givenVal, err = stack1.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Println("Pop:", givenVal)

	err = stack1.Push(2)
	fmt.Println("Push:", 2)
	err = stack1.Push(5)
	fmt.Println("Push:", 5)
	err = stack1.Push(4)
	fmt.Println("Push:", 4)

	if err != nil {
		panic(err)
	}
	fmt.Println("Sum:", stack1.Sum())

	givenVal, err = stack1.Peek()
	if err != nil {
		panic(err)
	}
	fmt.Println("Peek:", givenVal)

	fmt.Println("Sum:", stack1.Sum())

	givenVal, err = stack1.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Println("Pop:", givenVal)
	fmt.Println("Sum:", stack1.Sum())

}
