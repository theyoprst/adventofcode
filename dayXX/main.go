package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var ans1, ans2 int
	for scanner.Scan() {
		line := scanner.Text()
		_ = line
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}
