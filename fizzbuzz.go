package fizzbuzz

import "fmt"

func GetOutput(i int) string {
	x := i % 3
	y := i % 5

	switch {
		case x == 0 && y == 0:
			return "FizzBuzz"
		case x == 0:
			return "FizzBuzz"
		case y == 0:
			return "FizzBuzz"
		default:
			return fmt.Sprintf("%d", i)		
	}
}