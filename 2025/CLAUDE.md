# CLAUDE.md - Swift Solutions (2025)

This file provides guidance for working with the Swift-based Advent of Code 2025 solutions.

## Project Overview

This is a Swift Package Manager project with a unified executable architecture:
- **Single executable** (`aoc`) with `--day=` flag for day selection
- Solutions in `Sources/aoc/Days/` directory (Day01.swift, Day02.swift, etc.)
- Resources in `Resources/DD/` directories (tests.yaml, input files)
- Tests parse the same `tests.yaml` format as Go solutions (language-agnostic)
- Single test target (`AOCTests`) testing all days
- Package.swift defines the executable and test targets

## Project Structure

```
2025/
├── Package.swift                      # Package manifest
├── Sources/aoc/                       # Main executable target
│   ├── main.swift                     # Entry point with CLI arg parsing
│   ├── DayRegistry.swift              # Central registry and DaySolution protocol
│   ├── Runner.swift                   # Automatic execution of all inputs
│   └── Days/
│       ├── Day01.swift                # Individual day solutions
│       ├── Day02.swift
│       └── ...
├── Tests/AOCTests/
│   ├── AllDaysTests.swift             # Consolidated test file
│   └── AOCTestRunner.swift            # YAML-based test runner
├── Resources/
│   ├── 01/                            # Day-specific resources
│   │   ├── tests.yaml                 # Test configuration
│   │   ├── input.txt                  # Actual input (gitignored)
│   │   └── input_ex*.txt              # Example inputs
│   ├── 02/
│   └── ...
└── AOCUtilities/                      # Shared utilities library
    ├── Solution.swift                 # Solution wrapper type
    ├── IntParsing.swift               # Integer parsing utilities
    └── StringExtensions.swift         # String helpers
```

## Common Commands

### Running Solutions

From the 2025 directory:
```bash
swift run aoc                    # Run latest day with all inputs
swift run aoc --day=07           # Run specific day with all inputs
swift run aoc --day=last         # Run latest day (explicit)
```

The program automatically runs all solution variants against all discovered input files in order (example inputs first, then main input):
```
Running Day 07

=== Processing: input_ex1.txt ===
Part 1 (Default): 3749
Part 2 (Default): 11387

=== Processing: input.txt ===
Part 1 (Default): 1298103531759
Part 2 (Default): 140575048428831
```

### Running Tests

```bash
swift test                       # All tests
swift test --filter testDay07    # Specific day
```

### Building

```bash
swift build                      # Debug build
swift build -c release           # Release build
.build/release/aoc --day=07      # Run release build
```

### Linting

SwiftLint with comprehensive rules:
```bash
swiftlint                        # Lint all files
swiftlint --fix                  # Auto-fix violations
```

Configuration at `.swiftlint.yml`:
- Comprehensive opt-in rules for code quality and best practices
- Relaxed rules for solution files (Days/*.swift)
- Strict linting maintained for `AOCUtilities/` and test files
- Excludes `.build/`, `.swiftpm/` directories
- Allows TODOs in incomplete solutions
- Pragmatic limits on function/file length and complexity for puzzle solving

## Solution Structure

Each Swift solution follows this pattern in `Sources/aoc/Days/DayXX.swift`:

```swift
import Foundation
import AOCUtilities

private func solvePart1(_ lines: [String]) -> Int {
    // Solution logic
    return result
}

private func solvePart2(_ lines: [String]) -> Int {
    // Solution logic
    return result
}

struct DayXX: DaySolution {
    let dayNumber = XX

    let part1Solutions = [
        Solution(name: "Default", solve: solvePart1)
    ]

    let part2Solutions = [
        Solution(name: "Default", solve: solvePart2)
    ]
}
```

### Design Principles

1. **DaySolution protocol**: Each day implements this protocol with instance properties
   - `part1Solutions: [Solution]` - Array of Part 1 solution variants
   - `part2Solutions: [Solution]` - Array of Part 2 solution variants
   - `dayNumber: Int` - Day number (e.g., 1, 7)

2. **Struct instances**: Days are structs that get instantiated
   - DayRegistry creates instances in the `allDays` array
   - Tests create instances to pass to `runAOCTests()`

3. **Private helpers**: All helper functions are `private` to avoid name conflicts
   - Multiple days can have `solvePart1`, `solvePart2`, etc. without conflicts
   - Only the solution arrays and protocol conformance are public/internal

4. **Instance solutions**: Solution arrays are instance properties
   - Created when the day struct is instantiated
   - Cleaner namespace, better encapsulation

5. **DayRegistry dispatch**: Central registry manages all day instances
   - Maintains array of all days and dictionary for O(1) lookup
   - Dynamically computes last day from available instances

6. **Multiple solutions**: Can have multiple solution variants with different approaches
   ```swift
   let part2Solutions = [
       Solution(name: "Iterative", solve: solvePart2),
       Solution(name: "Linear", solve: solvePart2Linear)
   ]
   ```

## Test Structure

Tests are consolidated in `Tests/AOCTests/AllDaysTests.swift`:

```swift
import Testing
@testable import aoc

@Suite("Advent of Code 2025")
struct AllDaysTests {
    @Test("Day 07")
    func testDay07() throws {
        try runAOCTests(bundle: .module, day: Day07())
    }
}
```

### Test Configuration (tests.yaml)

Each day has a `Resources/DD/tests.yaml` file (same format as Go solutions):

```yaml
inputs:
- path: input_ex1.txt
  wantPart1: "42"
  wantPart2: "123"
- path: input.txt
  wantPart1: "2904518"
  wantPart2: "18650129"
```

The `runAOCTests()` function in `Tests/AOCTests/AOCTestRunner.swift`:
- Eliminates YAML parsing code duplication across all test files
- Accepts a day instance conforming to `DaySolution` protocol
- Automatically extracts part1/part2 solutions and day number from the day
- Reads tests.yaml and runs all test cases automatically
- Runs all cases defined in the same `tests.yaml` format as Go solutions
- Internal to the test module (not a separate library)

## Key Utilities

### AOCUtilities

- `Solution<T>` - Wrapper type for solution functions with names
  ```swift
  Solution(name: "Default", solve: solvePart1)
  ```

- `runAllInputs()` - Automatic execution with all discovered inputs
  - Discovers input*.txt files in resource bundle (example inputs first, then main input)
  - Automatically runs all solution variants for both parts against each input
  - Clear output separating results by input file

- `Int.mustParse(_:)` - Parse integer from String/Substring, crashes on failure
  ```swift
  let num = Int.mustParse("123")  // 123
  let num = Int.mustParse("abc")  // Fatal error
  ```

### Test Utilities (AOCTests module)

- `runAOCTests()` - YAML-based test runner (in `AOCTestRunner.swift`)
  - Reads tests.yaml from bundle
  - Runs all test cases against all solution variants
  - Reports failures with clear error messages
  - Internal function available within the test module

## Architecture Details

### DaySolution Protocol

```swift
protocol DaySolution: Sendable {
    var part1Solutions: [Solution] { get }
    var part2Solutions: [Solution] { get }
    var dayNumber: Int { get }
}
```

### DayRegistry

Central registry with array of day instances and O(1) lookup:

```swift
enum DayRegistry {
    @MainActor
    private static let allDays: [any DaySolution] = [
        Day01(),
        Day02(),
        // ...
    ]

    @MainActor
    private static let dayMap: [Int: any DaySolution] = {
        Dictionary(uniqueKeysWithValues: allDays.map { ($0.dayNumber, $0) })
    }()

    @MainActor
    static var lastDay: Int {
        allDays.map(\.dayNumber).max() ?? 1
    }

    @MainActor
    static func get(day: Int) -> (any DaySolution)? {
        dayMap[day]
    }
}
```

### Main Entry Point

```swift
@MainActor
func main() {
    // Parse --day argument (supports --day=<number> or --day=last)
    let day = /* parse from CommandLine.arguments */ ?? DayRegistry.lastDay

    // Get the day solution from registry
    guard let daySolution = DayRegistry.get(day: day) else {
        print("Error: Day \(day) not found")
        exit(1)
    }

    // Run the day's solution with all inputs
    let bundle = findResourceBundle()
    runAllInputs(
        part1Solutions: daySolution.part1Solutions,
        part2Solutions: daySolution.part2Solutions,
        bundle: bundle,
        daySubdirectory: String(format: "%02d", daySolution.dayNumber)
    )
}

main()
```

## Adding a New Day

1. **Create solution file**: `Sources/aoc/Days/DayXX.swift`
   ```swift
   import Foundation
   import AOCUtilities

   private func solvePart1(_ lines: [String]) -> Int {
       // TODO: Implement
       return 0
   }

   private func solvePart2(_ lines: [String]) -> Int {
       // TODO: Implement
       return 0
   }

   struct DayXX: DaySolution {
       let dayNumber = XX

       let part1Solutions = [
           Solution(name: "Default", solve: solvePart1)
       ]

       let part2Solutions = [
           Solution(name: "Default", solve: solvePart2)
       ]
   }
   ```

2. **Add to DayRegistry**: Update `DayRegistry.swift`
   ```swift
   @MainActor
   private static let allDays: [any DaySolution] = [
       Day01(),
       Day02(),
       // ... existing days
       DayXX(),
   ]
   ```

3. **Add test**: Update `Tests/AOCTests/AllDaysTests.swift`
   ```swift
   @Test("Day XX")
   func testDayXX() throws {
       try runAOCTests(bundle: .module, day: DayXX())
   }
   ```

4. **Update Package.swift**: Add resources
   ```swift
   resources: [
       // ... existing resources
       .copy("../../Resources/XX"),
   ]
   ```

5. **Create resource directory**: `Resources/XX/`
   - Add `tests.yaml` with test cases
   - Add `input_ex1.txt` with example input
   - Add `input.txt` with actual input (gitignored)
