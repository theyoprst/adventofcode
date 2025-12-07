# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is an Advent of Code solutions repository with solutions in multiple languages:
- **Go** (2022-2024): Solutions using a shared `aoc` package with common utilities
- **Swift** (2025+): Solutions organized as a Swift Package with language-agnostic YAML-based testing

## Multi-Language Structure

### 2022-2024 (Go)
Solutions use the established Go pattern with the shared `aoc` package.

### 2025+ (Swift)
Swift solutions are organized differently while maintaining compatibility:
- Solutions live directly in `2025/DD/` directories (no language subdirectory)
- Each day has `Solution.swift` and `SolutionTests.swift`
- Tests parse the same `tests.yaml` format as Go (language-agnostic)
- Package.swift at year level defines executable targets per day
- Input files (`input.txt`, `input_ex*.txt`) shared at day level
- If Go solutions are added later, they can go in `2025/DD/go/` subdirectories

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

#### Swift Linting

The project uses SwiftLint with comprehensive rules:
```bash
cd 2025 && swiftlint         # Run from Swift project directory
```

The configuration at `2025/.swiftlint.yml` follows the same philosophy as Go linting: comprehensive but practical for Advent of Code. Key features:
- Comprehensive opt-in rules for code quality and best practices
- Relaxed rules for solution files (similar to Go's `\d+/main.go` pattern)
- Strict linting maintained for `AOCUtilities/` and `AOCTestSupport/`
- Excludes `.build/`, `.swiftpm/` directories and dependencies
- Allows TODOs in incomplete solutions
- Pragmatic limits on function/file length and complexity for puzzle solving

### Swift Solutions (2025+)

Running Swift solutions from the year directory:
```bash
cd 2025
swift run day00 < 00/input.txt
swift run day01 < 01/input.txt
```

Running Swift tests:
```bash
cd 2025
swift test                       # All tests
swift test --filter Day00Tests   # Specific day
```

Building for release:
```bash
cd 2025
swift build -c release
.build/release/day00 < 00/input.txt
```

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

### Go Solutions (2022-2024)

Each Go solution follows this pattern:

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

### Swift Solutions (2025+)

Each Swift solution follows this pattern:

```swift
import Foundation

func solvePart1(_ lines: [String]) -> Int {
    // Solution logic
    return result
}

func solvePart2(_ lines: [String]) -> Int {
    // Solution logic
    return result
}

@main
struct DayXX {
    static func main() {
        var lines: [String] = []
        while let line = readLine() {
            lines.append(line)
        }
        print("Part 1:", solvePart1(lines))
        print("Part 2:", solvePart2(lines))
    }
}
```

Swift tests use the shared `AOCTestSupport` library (similar to Go's `aoc.RunTests()`):

```swift
import Testing
import AOCTestSupport
@testable import DayXX

@Suite("Day XX Solutions")
struct DayXXTests {
    @Test("All test cases from tests.yaml")
    func testFromYAML() throws {
        try runAOCTests(bundle: .module, solvePart1: solvePart1, solvePart2: solvePart2)
    }
}
```

The `AOCTestSupport` library:
- Eliminates YAML parsing code duplication across all test files
- Provides `runAOCTests()` function that reads tests.yaml and runs all test cases
- Automatically runs all cases defined in the same `tests.yaml` format as Go
- Test files reduced from ~45 lines to ~10 lines

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
