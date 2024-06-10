package main_test

import (
	"testing"

	"github.com/nandobas/gohr/formulas"
	"github.com/stretchr/testify/suite"
)

type mainTestSuite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(mainTestSuite))
}

func (t *mainTestSuite) TestFibonacci() {
	var givenSequence []int
	expectedSequence := []int{
		0,
		1,
		1,
		2,
		3,
		5,
		8,
		13,
		21,
		34}

	f := formulas.Fibonacci()
	for i := 0; i < 10; i++ {
		givenInt := f()
		givenSequence = append(givenSequence, givenInt)
	}

	t.ElementsMatch(givenSequence, expectedSequence)
}
