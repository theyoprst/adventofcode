import Foundation
import AOCUtilities

private func maxJoltage(_ line: String, digits: Int) -> Int {
    // TODO: Try to fix type conversion mess
    let line = [UInt8](line.utf8)
    let zeroCode = Character("0").asciiValue! // TODO: Fix force unwrap
    var joltage = 0
    var maxIdx = -1
    for skipTail in (0..<digits).reversed() {
        var maxCh = UInt8(0)
        for i in maxIdx+1..<line.count-skipTail where line[i] - zeroCode > maxCh {
            maxCh = line[i] - zeroCode
            maxIdx = i
        }
        joltage = joltage * 10 + Int(maxCh)
    }
    return joltage
}

private func solvePart1(_ lines: [String]) -> Int {
    return lines.reduce(into: 0) { sum, line in
        sum += maxJoltage(line, digits: 2 )
    }
}

private func solvePart2(_ lines: [String]) -> Int {
    return lines.reduce(into: 0) { sum, line in
        sum += maxJoltage(line, digits: 12 )
    }
}

struct Day03: DaySolution {
    let dayNumber = 3

    let part1Solutions = [
        Solution(name: "Default", solve: solvePart1)
    ]

    let part2Solutions = [
        Solution(name: "Default", solve: solvePart2)
    ]
}
