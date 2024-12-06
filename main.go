package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var op, x1, x2 string
	var sum, n1, n2 int
	var isRome bool

	_, err := fmt.Scanln(&x1, &op, &x2)
	
	if err != nil {
		fmt.Println("Panic!")
		os.Exit(0)
	}

	if (x1 == "1" || x1 == "2" || x1 == "3" || x1 == "4" || x1 == "5" ||
		  x1 == "6" || x1 == "7" || x1 == "8" || x1 == "9" || x1 == "10") &&
		 (x2 == "1" || x2 == "2" || x2 == "3" || x2 == "4" || x2 == "5" ||
		  x2 == "6" || x2 == "7" || x2 == "8" || x2 == "9" || x2 == "10") {
		n1, _ = strconv.Atoi(x1)
		n2, _ = strconv.Atoi(x2)
		isRome = false
	} else {
			roman := map[string]int {
				"I":    1,
				"II":   2,
				"III":  3,
				"IV":   4,
				"V":    5,
				"VI":   6,
				"VII":  7,
				"VIII": 8,
				"IX":   9,
				"X":    10,
			}
			n1 = roman[x1]
			n2 = roman[x2]
			isRome = true
			if n1 == 0 || n2 == 0 {
				fmt.Println("Panic!")
				os.Exit(0)
			}
		}

	if (n1 > 10 || n1 < 0) || (n2 > 10 || n2 < 0) {
		fmt.Println("Panic!")
		os.Exit(0)
	}

	if op == "+" {
			sum = n1 + n2
	} else if op == "-" {
			sum = n1 - n2
	} else if op == "*" {
			sum = n1 * n2
	} else if op == "/" {
			sum = n1 / n2
	} else {
			fmt.Println("Panic!")
			os.Exit(0)
	}

	if isRome == false {
			fmt.Println(sum)
	}	else if sum <= 0 {
			fmt.Println("Panic!")
			os.Exit(0)
	} else {
			arabic := map[int]string{
				1:"I",
				2:"II",
				3:"III",
				4:"IV",
				5:"V",
				6:"VI",
				7:"VII",
				8:"VIII",
				9:"IX",
			}

			d := arabic[sum%10]

			if sum < 10 {
					fmt.Println(d)
			} else if sum < 20 {
					fmt.Println("X"+d)
			} else if sum < 30 {
					fmt.Println("XX"+d)
			} else if sum < 40 {
					fmt.Println("XXX"+d)
			} else if sum < 50 {
					fmt.Println("XL"+d)
			} else if sum < 60 {
					fmt.Println("L"+d)
			} else if sum < 70 {
					fmt.Println("LX"+d)
			} else if sum < 80 {
					fmt.Println("LXX"+d)
			} else if sum < 90 {
					fmt.Println("LXXX"+d)
			} else if sum < 100 {
					fmt.Println("XC"+d)
			} else if sum == 100 {
					fmt.Println("C")
			}
	}
}
