package entities

import "fmt"

type myTypesSum interface {
	int64 | float64 | string
}

func sum[K comparable, V myTypesSum](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func PrintGenerics() {

	listInt64 := map[string]int64{
		"a": 40,
		"b": 30,
		"c": 50,
	}
	listFloat64 := map[int]float64{
		1:  10.5,
		2:  15.30,
		30: 20.222,
	}

	fmt.Println("Int64:", sum(listInt64))
	fmt.Println("Float64:", sum(listFloat64))

}
