package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wangkekekexili/calculator"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		fmt.Println(calculator.Do(input))
	}
}
