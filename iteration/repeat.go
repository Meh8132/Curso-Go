package iteration

import "fmt"

func Repeat(character string, number int) string {
	var repeated string
	for i := 0; i < number; i++ {
		repeated += character
	}
	return repeated
}

func ExampleRepeat() {
	repeated := Repeat("a", 6)
	fmt.Println(repeated)
	// Output: aaaaaa
}
