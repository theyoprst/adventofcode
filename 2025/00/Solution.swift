import Foundation
import AOCUtilities

func solvePart1(_ lines: [String]) -> Int {
    var leftList: [Int] = []
    var rightList: [Int] = []

    for line in lines {
        let numbers = line.extractInts()
        precondition(numbers.count == 2, "Expected 2 numbers per line, got \(numbers.count) in: \(line)")
        leftList.append(numbers[0])
        rightList.append(numbers[1])
    }

    leftList.sort()
    rightList.sort()

    return zip(leftList, rightList)
        .map { abs($0 - $1) }
        .reduce(0, +)
}

func solvePart2(_ lines: [String]) -> Int {
    var leftList: [Int] = []
    var rightCount: [Int: Int] = [:]

    for line in lines {
        let numbers = line.extractInts()
        precondition(numbers.count == 2, "Expected 2 numbers per line, got \(numbers.count) in: \(line)")
        leftList.append(numbers[0])
        rightCount[numbers[1], default: 0] += 1
    }

    return leftList
        .map { $0 * (rightCount[$0] ?? 0) }
        .reduce(0, +)
}

@main
struct Day00 {
    static func main() {
        var lines: [String] = []
        while let line = readLine() {
            lines.append(line)
        }

        print("Part 1:", solvePart1(lines))
        print("Part 2:", solvePart2(lines))
    }
}
