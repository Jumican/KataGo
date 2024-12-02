package main

import (
	"fmt"
	"os"
)

func main() {

	var op string
	var n1 int
	var n2 int
	var sum int

	_, err := fmt.Scanln(&n1, &op, &n2)
	
	if err != nil {
		fmt.Println("Panic!")
		os.Exit(0)
	}

	if (n1 > 10 || n1 < 0) || (n2 > 10 || n2 < 0) {
		fmt.Println("Panic!")
		os.Exit(0)
	}

	if op == "+" {
		sum = n1 + n2
	} else {
			if op == "-" {
				sum = n1 - n2
			} else {
					if op == "*" {
						sum = n1 * n2
					} else {
							if op == "/" {
								sum = n1 / n2
							} else {
									fmt.Println("Panic!")
									os.Exit(0)
								}
						}
				}
		}

	fmt.Println(sum)

}