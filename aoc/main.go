package aoc

import (
	"context"
	"flag"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"slices"
	"strings"
)

type Solver func(context.Context, []string) any

func getFunctionName(temp interface{}) string {
	path := runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()
	return path[strings.LastIndex(path, ".")+1:]
}

func Main(solversPart1, solversPart2 []Solver) {
	inputFlag := flag.String("input", "", "input file path")
	flag.Parse()
	lines := ReadInputLines(*inputFlag)
	var cmd string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	var solvers []Solver
	if cmd != "part2" || cmd == "part1" {
		solvers = append(solvers, solversPart1...)
	}
	if cmd != "part1" || cmd == "part2" {
		solvers = append(solvers, solversPart2...)
	}
	ctx := context.Background()
	if *inputFlag != "" {
		ctx = contextWithParams(ctx, paramsForInput(*inputFlag))
	}
	for _, solver := range solvers {
		fmt.Printf("%s: %v\n", getFunctionName(solver), solver(ctx, slices.Clone(lines)))
	}
}
