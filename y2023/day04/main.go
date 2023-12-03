package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/theyoprst/adventofcode/must"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var ans1, ans2 int
	for scanner.Scan() {
		line := scanner.Text()
		_ = line
	}
	must.NoError(scanner.Err())

	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}
