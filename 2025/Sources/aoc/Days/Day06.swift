import Foundation
import AOCUtilities

private func solvePart1(_ lines: [String]) -> Int {
    assert(lines.count >= 2)

    let table = lines.dropLast().map { line in
        line.split(separator: " ").map(Int.mustParse)
    }
    assert(table.allSatisfy { row in row.count == table[0].count })

    let ops: [Character] = lines.last!.split(separator: " ").map { $0.first! }
    assert(ops.count == table[0].count)

    var grandTotal = 0
    for (col, chOp) in ops.enumerated() {
        let (op, identity): ((Int, Int) -> Int, identity: Int) = switch chOp {
        case "+": ((+), 0)
        case "*": ((*), 1)
        default: fatalError("Unexpected character: \(chOp)")
        }
        grandTotal += table.reduce(identity) { op($0, $1[col]) }
    }

    return grandTotal
}

private func solvePart2(_ lines: [String]) -> Int {
    var grid = lines.map { Array($0) }
    let rows = grid.count
    let cols = grid.map { $0.count }.max()! + 1 // +1 for extra spaces column.

    // Make rows equal length
    grid = grid.map { row in
        row + Array(repeating: " ", count: cols - row.count)
    }

    func parseColumnNumber(_ col: Int) -> Int? {
        var result = 0
        var hasDigit = false
        for row in 0..<rows {
            if let digit = grid[row][col].wholeNumberValue {
                result = 10 * result + digit
                hasDigit = true
            }
        }
        return hasDigit ? result : nil
    }

    typealias OpPair = (op: (Int, Int) -> Int, identity: Int)

    func parseColumnOp(_ col: Int) -> OpPair? {
        switch grid[rows-1][col] {
        case "+": return ((+), 0)
        case "*": return ((*), 1)
        default: return nil
        }
    }

    var grandTotal = 0
    var currentNumbers: [Int] = []
    var currentOp: OpPair?
    for col in 0..<cols {
        if let colNumber = parseColumnNumber(col) {
            currentNumbers.append(colNumber)
            if let op = parseColumnOp(col) {
                assert(currentOp == nil)
                currentOp = op
            }
        } else {
            let (op, identity) = currentOp!
            grandTotal += currentNumbers.reduce(identity, op)
            (currentNumbers, currentOp) = ([], nil)
        }
    }
    assert(currentOp == nil)
    return grandTotal
}

struct Day06: DaySolution {
    let dayNumber = 6

    let part1Solutions = [
        Solution(name: "Default", solve: solvePart1)
    ]

    let part2Solutions = [
        Solution(name: "Default", solve: solvePart2)
    ]
}
