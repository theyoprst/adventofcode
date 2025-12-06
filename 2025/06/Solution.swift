import Foundation
import AOCUtilities

let chAdd: Character = "+"
let chMul: Character = "*"

func solvePart1(_ lines: [String]) -> Int {
    precondition(lines.count >= 2)
    
    var numbers: [[Int]] = []
    for i in 0..<lines.count-1 {
        numbers.append(lines[i].split(separator:" ").map(Int.mustParse))
        precondition(numbers[i].count == numbers[0].count)
    }
    
    let ops: [Character] = lines[lines.count-1].split(separator:" ").map { $0.first! }
    precondition(ops.count == numbers[0].count)
    
    var grandTotal = 0
    for (col, chOp) in ops.enumerated() {
        let op: (Int, Int) -> Int
        var sum: Int
        switch chOp {
        case chAdd:
            op = (+)
            sum = 0
        case chMul:
            op = (*)
            sum = 1
        default:
            fatalError("Unexpected character: \(chOp)")
        }
        for row in 0..<numbers.count {
            sum = op(sum, numbers[row][col])
        }
        grandTotal += sum
    }
    return grandTotal
}

func solvePart2(_ lines: [String]) -> Int {
    var lines = lines.map { Array($0) }
    var grandTotal = 0
    var localTotal = 0
    var localOp: (Int, Int) -> Int = (+) // TODO: had to put value because of a weird compiler error about using before initialize
    let rows = lines.count
    let cols = lines.reduce(0) { max($0, $1.count) }

    // Make rows equal length
    for row in 0..<rows {
        while lines[row].count < cols {
            lines[row].append(" ")
        }
    }

    for col in 0..<cols {
        var colNumber = 0
        for row in 0..<lines.count-1 {
            let ch: Character = lines[row][col]
            switch ch {
            case " ":
                continue
            case "0"..."9":
                let digit = ch.asciiValue! - Character("0").asciiValue!
                colNumber = 10 * colNumber + Int(digit)
            default:
                fatalError("Unexpected character: \(ch)")
            }
        }
        let colOpCh = lines[rows-1][col]
        switch colOpCh {
        case "+":
            localOp = (+)
            grandTotal += localTotal
            localTotal = colNumber
        case "*":
            localOp = (*)
            grandTotal += localTotal
            localTotal = colNumber
        case " ":
            if colNumber != 0 {
                localTotal = localOp(localTotal, colNumber)
            }
        default:
            fatalError("Unexpected operation: \(colOpCh)")
        }
    }
    grandTotal += localTotal // last column
    return grandTotal
}

let part1Solutions = [
    Solution(name: "Default", solve: solvePart1)
]

let part2Solutions = [
    Solution(name: "Default", solve: solvePart2)
]

@main
struct Day06 {
    static func main() {
        runInteractively(part1Solutions: part1Solutions, part2Solutions: part2Solutions, bundle: .module)
    }
}
