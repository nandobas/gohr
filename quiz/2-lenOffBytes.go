package quiz

import (
	"fmt"
	"unicode/utf8"
)

func LenOffBytes() {
	sample := "a£c"
	fmt.Printf("Leng of bytes for given string is: %d\n", len(sample))

	fmt.Printf("True Leng of given string is: %d\n", utf8.RuneCountInString(sample))
}
