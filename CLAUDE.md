# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is an Advent of Code solutions repository written in Go. It contains solutions for multiple years (2022, 2023, 2024) organized by year and day, along with a shared `aoc` package providing common utilities.

## Project Structure

- `YYYY/DD/` - Solutions for each day, where YYYY is the year and DD is the day (e.g., `2024/01/`)
  - `main.go` - Solution code with `SolvePart1` and `SolvePart2` functions
  - `main_test.go` - Test file that uses `aoc.RunTests()`
  - `tests.yaml` - Test configuration with input file paths and expected outputs
  - `input.txt` - Actual puzzle input (gitignored)
  - `input_ex*.txt` - Example inputs from problem descriptions
  - `part1.md`, `part2.md` - Problem descriptions (gitignored)
- `aoc/` - Shared utilities package
  - `aoc.go` - Core utilities (Ints, Blocks, GCD, LCM, etc.)
  - `main.go` - Main function runner for solutions
  - `test.go` - Test framework using YAML-based test definitions
  - `containers/` - Data structures (DSU, Set)
  - `fld/` - 2D field/grid utilities (Pos, Field)
  - `graphs/` - Graph algorithms (Dijkstra)
  - `queues/` - Priority queue implementations
  - `htmlparser/` - HTML parsing for problem descriptions
- `must/` - Must-style error handling utilities
- `template/` - Template for new day solutions
- `cmd/aoc-input/` - Tool to download inputs and problem descriptions from adventofcode.com

## Common Commands

### Running Solutions

From the repository root, navigate to a specific day's directory:
```bash
cd 2024/01
go run main.go < input.txt
```

Or run both parts from stdin:
```bash
cd 2024/01
cat input.txt | go run main.go
```

### Running Tests

Tests use a YAML-based configuration system. From a day's directory:
```bash
cd 2024/01
go test -v
```

Or from the root:
```bash
go test ./2024/01/...
```

To run a single test function:
```bash
cd 2024/01
go test -v -run Test/SolvePart1
```

### Linting

The project uses golangci-lint with extensive linter configuration:
```bash
golangci-lint run
```

To lint a specific directory:
```bash
golangci-lint run ./2024/01/...
```

The configuration at `.golangci.yaml` enables many linters with special exclusions for solution files matching `\d+/main.go`.

### Downloading Inputs

The `cmd/aoc-input` tool requires a session cookie configured in `~/.aoc-input.json`:
```json
{
  "sessionCookie": "your_session_cookie_here"
}
```

Run from repository root:
```bash
go run ./cmd/aoc-input/main.go
```

This automatically downloads missing `input.txt` files and problem descriptions for all year/day directories.

## Solution Structure

Each solution follows this pattern:

```go
package main

import (
    "context"
    "github.com/theyoprst/adventofcode/aoc"
)

func SolvePart1(ctx context.Context, lines []string) any {
    // Solution logic
    return result
}

func SolvePart2(ctx context.Context, lines []string) any {
    // Solution logic
    return result
}

var (
    solvers1 = []aoc.Solver{SolvePart1}
    solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
    aoc.Main(solvers1, solvers2)
}
```

## Testing Structure

Each day has a `tests.yaml` file defining test cases:

```yaml
inputs:
- path: input_ex1.txt
  wantPart1: "11"
  wantPart2: "31"
- path: input.txt
  wantPart1: "2904518"
  wantPart2: "18650129"
```

The test file uses `aoc.RunTests()`:

```go
func Test(t *testing.T) {
    aoc.RunTests(t, solvers1, solvers2)
}
```

This automatically runs all solvers against all inputs defined in `tests.yaml`.

## Key Utilities in `aoc` Package

- `aoc.Ints(s string)` - Extracts all integers from a string (handles any delimiters)
- `aoc.Blocks(lines []string)` - Splits lines into blocks separated by empty lines
- `aoc.ReadInputLines(path string)` - Reads input file or stdin
- `aoc.GCD(a ...int)`, `aoc.LCM(a ...int)` - Math utilities
- `aoc.Abs[T](a T)` - Generic absolute value
- `aoc.Reversed[S, E](a S)` - Returns reversed copy of slice
- `aoc/fld.Pos` - 2D position with direction support
- `aoc/fld.Field` - 2D grid/field utilities
- `aoc/containers.DSU` - Disjoint Set Union data structure
- `aoc/containers.Set[T]` - Generic set implementation
- `aoc/graphs.Dijkstra()` - Dijkstra's algorithm
- `aoc/queues.PriorityQueue[T]` - Generic priority queue

## Import Organization

The project uses `gci` for import formatting with three sections:
1. Standard library
2. External dependencies
3. Local imports with prefix `github.com/theyoprst/adventofcode`

## Linter Notes

- Solution files (`\d+/main.go`) have relaxed rules for `prealloc`, `bodyclose`, `contextcheck`, `rowserrcheck`, `sqlclosecheck`, and `errorlint`
- Security linter `gosec` excludes G404 (insecure random is acceptable for AoC)
- All code must pass `gofmt`, `gofumpt`, `goimports`, and `gci` formatting
- Comments should end with periods (enforced by `godot`)
