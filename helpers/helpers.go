package helpers

import (
	"bufio"
	"os"
	"strings"

	"github.com/theyoprst/adventofcode/must"
)

func ReadInputLines() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	must.NoError(scanner.Err())
	return lines
}

func AddBorder2D(a []string, r rune) []string {
	b := string(r)
	cols := len(a[0]) + 2
	res := make([]string, 0, len(a)+2)
	res = append(res, strings.Repeat(b, cols))
	for _, s := range a {
		res = append(res, b+s+b)
	}
	res = append(res, strings.Repeat(b, cols))
	return res
}

func IsDigit[T byte | rune](ch T) bool {
	return '0' <= ch && ch <= '9'
}
