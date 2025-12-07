import Foundation
import AOCUtilities

let roll: Character = "@"
let empty: Character = "."

private func solvePart1(_ lines: [String]) -> Int {
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

private func solvePart2(_ lines: [String]) -> Int {
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

struct Day04: DaySolution {
    let dayNumber = 4

    let part1Solutions = [
        Solution(name: "Default", solve: solvePart1)
    ]

    let part2Solutions = [
        Solution(name: "Default", solve: solvePart2)
    ]
}
