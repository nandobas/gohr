package formulas

import (
	"fmt"
	"math"
)

func Baskara(a, b, c float64) (error, float64, float64) {

	delta := b*b - 4*a*c

	if b == 0 || c == 0 {
		return fmt.Errorf("incomplete equation"), 0, 0
	}

	if delta < 0 {
		return fmt.Errorf("equation dont have real square"), 0, 0
	}

	if delta == 0 {
		fmt.Println("equation have two equals real squares")
	} else {
		fmt.Println("equation have two dont equals real squares")
	}

	fmt.Println("(-b + ou - raiz(delta)) / (2 * a)")
	fmt.Println("(-", b, " + ou - raiz(", delta, ")) / (2 * ", a, ")")
	fmt.Println("(", -1*b, " + ou - ", math.Sqrt(delta), ") / ", 2*a)

	fmt.Println("\nx1 = (", -1*b+math.Sqrt(delta), ") / ", 2*a)
	x1 := (-1*b + math.Sqrt(delta)) / (2 * a)

	fmt.Printf("\nx2 = (", -1*b-math.Sqrt(delta), ") / ", 2*a)
	x2 := (-1*b - math.Sqrt(delta)) / (2 * a)
	return nil, x1, x2
}
