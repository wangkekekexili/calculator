package calculator

import "fmt"

// Do performs arithmetic calculation based on the input string.
func Do(input string) (float64, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic recovered: ", r)
		}
	}()
	return newInterpreter(input).calculate()
}
