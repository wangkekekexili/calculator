package calculator

// Do performs arithmetic calculation based on the input string.
func Do(input string) (float64, error) {
	return newInterpreter(input).expr()
}
