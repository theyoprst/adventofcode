import Foundation
import AOCUtilities

enum EventType: Int, Comparable {
    case intervalStart, test, intervalFinish
    static func < (lhs: Self, rhs: Self) -> Bool {
        lhs.rawValue < rhs.rawValue
    }
}

struct Event: Comparable {
    let id: Int
    let type: EventType
    static func < (lhs: Self, rhs: Self) -> Bool {
        (lhs.id, lhs.type) < (rhs.id, rhs.type)
    }
}

func buildEvents(_ lines: [String]) -> [Event] {
    let sections = lines.split(separator: "", omittingEmptySubsequences: false)
    precondition(sections.count == 2, "Want 2 blocks, got \(sections.count)")
    let (intervalsSection, testsSection) = (sections[0], sections[1])
    let intervalEvents = intervalsSection.flatMap { line -> [Event] in
        let parts = line.split(separator: "-").map(Int.mustParse)
        precondition(parts.count == 2, "Invalid interval \(line)")
        let (start, finish) = (parts[0], parts[1])
        return [Event(id: start, type: .intervalStart), Event(id: finish, type: .intervalFinish)]
    }
    let testEvents = testsSection.map { line in
        Event(id: Int.mustParse(line), type: .test)
    }
    return (intervalEvents + testEvents).sorted()
}

func solvePart1(_ lines: [String]) -> Int {
    var (result, depth) = (0, 0)
    for event in buildEvents(lines) {
        switch event.type {
        case .intervalStart:
            depth += 1
        case .intervalFinish:
            depth -= 1
        case .test:
            if depth > 0 { result += 1 }
        }
    }

    return result
}

func solvePart2(_ lines: [String]) -> Int {
    var (result, depth, firstStart) = (0, 0, 0)
    for event in buildEvents(lines) {
        switch event.type {
        case .intervalStart:
            if depth == 0 { firstStart = event.id }
            depth += 1
        case .intervalFinish:
            depth -= 1
            if depth == 0 { result += event.id - firstStart + 1 }
        case .test:
            continue // Ignore tests in part2
        }
    }

    return result
}

let part1Solutions = [
    Solution(name: "Default", solve: solvePart1)
]

let part2Solutions = [
    Solution(name: "Default", solve: solvePart2)
]

@main
struct Day05 {
    static func main() {
        runInteractively(part1Solutions: part1Solutions, part2Solutions: part2Solutions, bundle: .module)
    }
}
