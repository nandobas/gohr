package formulas

func Fibonacci() func() int {
	a := 0
	out := 0
	last := 0
	return func() int {
		last = a + out
		out = a
		a = last
		if a == 0 {
			a++
		}
		return out
	}
}
