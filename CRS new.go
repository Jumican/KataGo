package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Text struct {
	Content string
	Spaces string
}

func (t *Text) textModifier() {
	for {
		t.Spaces = strings.ReplaceAll(t.Content, "  ", " ")

		if t.Content == t.Spaces {
			break
		}

		t.Content = t.Spaces
	}
	
	runeArray := []rune(t.Content)

	for i := 1; i < len(runeArray)-1; i++ {
		if runeArray[i] == 45 {

				temp := runeArray[i-1]
				runeArray[i-1] = runeArray[i+1]
				runeArray[i+1] = temp
				runeArray[i] = 0

				i = 1

				t.Content = string(runeArray)
		}
	}

	t.Content = strings.ReplaceAll(t.Content, "-", "")

	t.Content = strings.ReplaceAll(t.Content, "+", "!")

	t.Content = strings.ReplaceAll(t.Content, "0", "")

	runeArray = []rune(t.Content)
	c := 0
	nums := map[rune] int {
		49: 1,
		50: 2,
		51: 3,
		52: 4,
		53: 5,
		54: 6,
		55: 7,
		56: 8,
		57: 9,
	}

	for i := 0; i < len(runeArray); i++ {

		if nums[runeArray[i]] != 0 {

			c += nums[runeArray[i]]
			runeArray[i] = 0
			i = 0

			t.Content = string(runeArray)
		}
	}

	if c != 0 {
		t.Content += " " + strconv.Itoa(c)
	}

	fmt.Println(t.Content)
}

func main() {
	text := &Text{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите строку:")

	for scanner.Scan() {
		text.Content = scanner.Text()
		text.textModifier()
	}
}
