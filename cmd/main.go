package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/wangkekekexili/calculator"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		input = strings.Trim(input, " \r\n")
		if strings.TrimSpace(input) == "" {
			continue
		}
		result, err := calculator.Do(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if result == math.Floor(result) {
			fmt.Printf("%.0f\n", result)
		} else {
			fmt.Printf("%.2f\n", result)
		}
	}
}
