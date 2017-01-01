package calculator

import "testing"

func TestCalculator(t *testing.T) {
	tests := []struct {
		input  string
		output float64
	}{{
		input:  "1+2",
		output: 3,
	}, {
		input:  "200-100",
		output: 100,
	}, {
		input:  " 10 +  11 ",
		output: 21,
	}, {
		input:  "1+2+3",
		output: 6,
	}, {
		input:  "2 + 7 * 4",
		output: 30,
	}, {
		input:  "7 - 8 / 4",
		output: 5,
	}, {
		input:  "14 + 2 * 3 - 6 / 2",
		output: 17,
	}, {
		input:  "7 + 3 * (10 / (12 / (3 + 1) - 1))",
		output: 22,
	}}
	for _, test := range tests {
		output, err := Do(test.input)
		if err != nil {
			t.Fatal(err)
		}
		if output != test.output {
			t.Fatalf("expect to get %v; got %v", test.output, output)
		}
	}
}
