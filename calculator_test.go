package calculator

import "testing"

func TestCalculator(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{{
		input:  "1+2",
		output: 3,
	}, {
		input:  "200-100",
		output: 100,
	}}
	for _, test := range tests {
		output := Do(test.input)
		if output != test.output {
			t.Fatalf("expect to get %v; got %v", test.output, output)
		}
	}
}
