import Foundation
import AOCUtilities

func parseRotations(_ lines: [String]) -> [Int] {
    var rotations: [Int] = []
    for line in lines {
        guard let direction = line.first else {
            precondition(false, "Empty line found")
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

func solvePart1(_ lines: [String]) -> Int {
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

func solvePart2(_ lines: [String]) -> Int {
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

func solvePart2Linear(_ lines: [String]) -> Int {
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

let part1Solutions = [
    Solution(name: "Default", solve: solvePart1)
]

let part2Solutions = [
    Solution(name: "Iterative", solve: solvePart2),
    Solution(name: "Linear", solve: solvePart2Linear)
]

@main
struct Day01 {
    static func main() {
        var lines: [String] = []
        while let line = readLine() {
            lines.append(line)
        }

        for solution in part1Solutions {
            print("Part 1 (\(solution.name)):", solution.solve(lines))
        }
        for solution in part2Solutions {
            print("Part 2 (\(solution.name)):", solution.solve(lines))
        }
    }
}
