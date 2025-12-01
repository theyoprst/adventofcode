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
        let n = Int.mustParse(line.dropFirst(1))
        precondition(n != 0, "Rotation of 0 found")
        rotations.append(multiplier * n)
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

@main
struct Day01 {
    static func main() {
        var lines: [String] = []
        while let line = readLine() {
            lines.append(line)
        }
        print("Part 1:", solvePart1(lines))
        print("Part 2:", solvePart2(lines))
    }
}
