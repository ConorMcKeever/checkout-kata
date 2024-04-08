package main

import (
	"bufio"
	"checkout-kata/pkg"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	checkout := pkg.NewCheckout()

	fmt.Println("Enter SKUs to scan:")
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		err := checkout.Scan(line)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}

	if scanner.Err() != nil {
		fmt.Printf("Error reading input: %v\n", scanner.Err())
	}
}
