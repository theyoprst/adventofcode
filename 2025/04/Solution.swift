import Foundation
import AOCUtilities

let roll: Character = "@"
let empty: Character = "."

func solvePart1(_ lines: [String]) -> Int {
    let rows = lines.count
    let cols = lines[0].count
    var field = Array(repeating: Array(repeating: empty, count: cols+2), count: rows+2)
    for (row, line) in lines.enumerated() {
        field[row+1][1...cols] = ArraySlice(Array(line))
    }
    var ans = 0
    for row in 1...rows {
        for col in 1...cols {
            if field[row][col] != roll {
                continue
            }
            var neighbours = 0
            for dRows in -1...1 {
                for dCols in -1...1 {
                    if dRows == 0 && dCols == 0 {
                        continue
                    }
                    if field[row+dRows][col+dCols] != empty {
                        neighbours += 1
                    }
                }
            }
            if neighbours <= 3 {
                ans += 1
            }
        }
    }
    return ans
}

func solvePart2(_ lines: [String]) -> Int {
    let rows = lines.count
    let cols = lines[0].count
    var field = Array(repeating: Array(repeating: empty, count: cols+2), count: rows+2)
    for (row, line) in lines.enumerated() {
        field[row+1][1...cols] = ArraySlice(Array(line))
    }

    var prevRemoved = 0
    while true {
        var removed = prevRemoved
        for row in 1...rows {
            for col in 1...cols {
                if field[row][col] != roll {
                    continue
                }
                var neighbours = 0
                for dRows in -1...1 {
                    for dCols in -1...1 {
                        if dRows == 0 && dCols == 0 {
                            continue
                        }
                        if field[row+dRows][col+dCols] != empty {
                            neighbours += 1
                        }
                    }
                }
                if neighbours <= 3 {
                    field[row][col] = empty
                    removed += 1
                }
            }
        }
        if removed == prevRemoved {
            break
        }
        prevRemoved = removed
    }

    return prevRemoved
}

let part1Solutions = [
    Solution(name: "Default", solve: solvePart1)
]

let part2Solutions = [
    Solution(name: "Default", solve: solvePart2)
]

@main
struct Day04 {
    static func main() {
        runInteractively(part1Solutions: part1Solutions, part2Solutions: part2Solutions, bundle: .module)
    }
}
