package quiz

import "fmt"

type area func(int, int) int

func Area() {
	var fArea area = func(n1 int, n2 int) int {
		return n1 * n2
	}
	var gArea = getArea()
	print(2, 2, fArea)
	print(2, 2, gArea)
}

func getArea() area {
	return func(x int, y int) int {
		return x * y
	}
}

func print(a int, b int, f area) {
	fmt.Printf("Area is: %d\n", f(a, b))
}
