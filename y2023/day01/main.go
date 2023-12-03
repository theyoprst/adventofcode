package main

import (
	"fmt"
	"strings"
)

func s2i(s string) int {
	var first, last int
	for _, r := range s {
		if '0' <= r && r <= '9' {
			if first == 0 {
				first = int(r - '0')
			}
			last = int(r - '0')
		}
	}
	return first*10 + last
}

func main() {
	var sum1, sum2 int
	for {
		var s string
		fmt.Scan(&s)
		if len(s) == 0 {
			break
		}
		sum1 += s2i(s)
		s = strings.ReplaceAll(s, "one", "o1e")
		s = strings.ReplaceAll(s, "two", "t2o")
		s = strings.ReplaceAll(s, "three", "t3e")
		s = strings.ReplaceAll(s, "four", "f4r")
		s = strings.ReplaceAll(s, "five", "f5e")
		s = strings.ReplaceAll(s, "six", "s6x")
		s = strings.ReplaceAll(s, "seven", "s7n")
		s = strings.ReplaceAll(s, "eight", "e8t")
		s = strings.ReplaceAll(s, "nine", "n9e")
		sum2 += s2i(s)
	}
	fmt.Println("Ans1:", sum1)
	fmt.Println("Ans2:", sum2)
}
