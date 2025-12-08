import Foundation
import AOCUtilities

private func parseRotations(_ lines: [String]) -> [Int] {
    var rotations: [Int] = []
    for line in lines {
        guard let direction = line.first else {
            preconditionFailure("Empty line found")
        }
        let multiplier: Int
        switch direction {
        case "R":
            multiplier = 1
        case "L":
            multiplier = -1
        default:
            multiplier = 0
            precondition(false, "Unexpected direction: \(direction)")
        }
        let steps = Int.mustParse(line.dropFirst(1))
        precondition(steps != 0, "Rotation of 0 found")
        rotations.append(multiplier * steps)
    }
    return rotations
}

private func solvePart1(_ lines: [String]) -> Int {
    var idx = 50
    var password = 0
    for rotation in parseRotations(lines) {
        idx += rotation
        if idx % 100 == 0 {
            password += 1
        }
    }
    return password
}

private func solvePart2(_ lines: [String]) -> Int {
    let rotations = parseRotations(lines)
    var idx = 50
    var password = 0
    for rotation in rotations {
        let count = abs(rotation)
        let direction = rotation > 0 ? 1 : -1
        for _ in 0..<count {
            idx += direction
            if idx % 100 == 0 {
                password += 1
            }
        }
    }
    return password
}

let dialSize = 100

private func solvePart2Linear(_ lines: [String]) -> Int {
    let rotations = parseRotations(lines)
    var idx = 50
    var password = 0
    for rotation in rotations {
        password += abs(rotation / dialSize)
        let oldIdx = idx
        idx += rotation % dialSize
        if idx < 0 {
            if oldIdx > 0 {
                password += 1
            }
            idx += dialSize
        } else if idx == 0 {
            password += 1
        } else if idx >= dialSize {
            password += 1
            idx -= dialSize
        }
    }
    return password
}

struct Day01: DaySolution {
    let dayNumber = 1

    let part1Solutions = [
        Solution(name: "Default", solve: solvePart1)
    ]

    let part2Solutions = [
        Solution(name: "Iterative", solve: solvePart2),
        Solution(name: "Linear", solve: solvePart2Linear)
    ]
}
