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
			fmt.Println(err)
			continue
		}
		result, err := calculator.Do(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}
}
