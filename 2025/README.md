# Advent of Code 2025 - Swift Solutions

This directory contains Swift solutions for Advent of Code 2025, organized as a Swift Package.

## Directory Structure

Each day has its own directory with the following structure:
```
DD/
├── Solution.swift         # Solution code with solvePart1 and solvePart2 functions
├── SolutionTests.swift    # Tests using Swift Testing framework
├── tests.yaml            # Test configuration (language-agnostic)
├── input.txt             # Actual puzzle input (gitignored)
└── input_ex*.txt         # Example inputs from problem descriptions
```

## Running Solutions

From the 2025 directory:
```bash
swift run day00 < 00/input.txt    # Run day 00 (port of 2024/01)
swift run day01 < 01/input.txt    # Run day 01
```

## Running Tests

Tests use Swift Testing framework and parse tests.yaml for test cases:

```bash
swift test                        # Run all tests
swift test --filter Day00Tests    # Run tests for a specific day
```

## Test Configuration

Each day's `tests.yaml` defines test cases:

```yaml
inputs:
- path: input_ex1.txt
  wantPart1: "expected_part1_result"
  wantPart2: "expected_part2_result"
- path: input.txt
  wantPart1: "expected_part1_result"
  wantPart2: "expected_part2_result"
```

The Swift tests automatically parse this YAML file and run all defined test cases.

## Building for Release

For faster execution:
```bash
swift build -c release
.build/release/day00 < 00/input.txt
```

## Adding a New Day

1. Create a new directory: `mkdir DD`
2. Copy template files (Solution.swift, SolutionTests.swift, tests.yaml)
3. Update `Package.swift` to add new targets:
   ```swift
   .executableTarget(name: "dayDD", path: "DD", sources: ["Solution.swift"]),
   .testTarget(
       name: "DayDDTests",
       dependencies: ["DayDD", "Yams"],
       path: "DD",
       sources: ["SolutionTests.swift"],
       resources: [.copy("tests.yaml"), .copy("input.txt"), .copy("input_ex1.txt")]
   ),
   ```

## Downloading Inputs

Use the existing Go tool from the repository root:
```bash
cd ..
go run ./cmd/aoc-input/main.go
```

This will download missing `input.txt` files to each day directory.

## Requirements

- Swift 6.0 or later
- macOS 14.0 or later

## Day 00

Day 00 is a port of the 2024 Day 01 problem, used to validate the Swift implementation against known correct results.
