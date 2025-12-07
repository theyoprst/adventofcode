import Foundation
import AOCUtilities

func solvePart1(_ lines: [String]) -> Int {
    var grid = lines.map { Array($0) }
    var splitsCount = 0
    for row in grid.indices.dropLast() {
        for col in grid[row].indices {
            switch grid[row][col] {
            case ".":
                continue
            case "S", "|":
                if grid[row+1][col] != "^" {
                    grid[row+1][col] = "|"
                }
            case "^":
                assert(row > 0)
                if grid[row-1][col] == "|" {
                    grid[row+1][col-1] = "|"
                    grid[row+1][col+1] = "|"
                    splitsCount += 1
                }
            default:
                assert(false, "Unexpected character in grid: \(grid[row][col])")
            }
        }
    }

    return splitsCount
}

func solvePart2(_ lines: [String]) -> Int {
    let grid = lines.map { Array($0) }
    let cols = grid[0].count
    var curRoutes = Array(repeating: 0, count: cols)
    for row in grid.indices.dropLast() {
        let prevRoutes = curRoutes
        curRoutes = Array(repeating: 0, count: cols)
        for col in grid[row].indices {
            switch grid[row][col] {
            case ".":
                curRoutes[col] += prevRoutes[col]
            case "S":
                curRoutes[col] = 1
            case "^":
                curRoutes[col-1] += prevRoutes[col]
                curRoutes[col+1] += prevRoutes[col]
            default:
                assert(false, "Unexpected character in grid: \(grid[row][col])")
            }
        }
    }

    return curRoutes.reduce(0, +)
}

let part1Solutions = [
    Solution(name: "Default", solve: solvePart1)
]

let part2Solutions = [
    Solution(name: "Default", solve: solvePart2)
]

@main
struct Day07 {
    static func main() {
        runInteractively(part1Solutions: part1Solutions, part2Solutions: part2Solutions, bundle: .module)
    }
}
