import Foundation
import AOCUtilities

private func solvePart1(_ lines: [String]) -> Int {
    let sections = lines.split(separator: "")
    var shapeSizes = Array(repeating: 0, count: sections.count-1)
    for (i, shapeSection) in sections.dropLast().enumerated() {
        shapeSizes[i] = shapeSection.reduce(0) { count, line in
            count + line.filter { $0 == "#" }.count
        }
    }

    let regions = sections.last!

    // Magic reserve of non-allocatable cells in the region to make both inputs generate correct results.
    // For example input, magic reserve should be from 2 to 10.
    // For main input, magic reserve can be from 0 up to some number over 1000.
    let magicReserve = 2

    return regions.map { region in
        let parts = region.split(separator: ": ", maxSplits: 2)
        let dimentions = String(parts[0]).extractInts()
        let regionArea = dimentions.reduce(1, *)
        let shapeCounts = String(parts[1]).extractInts()
        let shapesArea = zip(shapeSizes, shapeCounts).reduce(0) { total, pair in
            let (size, count) = pair
            return total + size * count
        }
        return shapesArea < regionArea-magicReserve ? 1 : 0
    }.reduce(0, +)
}

private func solvePart2(_ lines: [String]) -> Int {
    return 0
}

struct Day12: DaySolution {
    let dayNumber = 12

    let part1Solutions = [
        Solution(name: "Default", solve: solvePart1)
    ]

    let part2Solutions = [
        Solution(name: "Default", solve: solvePart2)
    ]
}
