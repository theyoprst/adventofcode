import Foundation
import Collections
import AOCUtilities

private struct Pair: Comparable {
    var i, j: Int
    var distSquared: Int

    static func < (lhs: Self, rhs: Self) -> Bool {
        return lhs.distSquared < rhs.distSquared
    }
}

private struct DSU {
    struct Item {
        var parent: Int
        var size: Int
    }
    var items: [Item]

    init(count: Int) {
        self.items = (0..<count).map { Item(parent: $0, size: 1) }
    }

    mutating func getRoot(_ k: Int) -> Int {
        if items[k].parent == k {
            return k
        }
        let root = getRoot(items[k].parent)
        items[k].parent = root
        return root
    }

    mutating func join(_ a: Int, _ b: Int) -> Bool {
        let rootA = getRoot(a)
        let rootB = getRoot(b)
        if rootA == rootB {
            return false
        }

        if items[rootA].size < items[rootB].size {
            items[rootA].parent = rootB
            items[rootA].size += items[rootB].size
        } else {
            items[rootB].parent = rootA
            items[rootB].size += items[rootA].size
        }
        return true
    }
}

private func makePairsHeap(_ points: [[Int]]) -> Heap<Pair> {
    var pairs: [Pair] = []
    for (i, point1) in points.enumerated() {
        for (j, point2) in points.enumerated().dropFirst(i+1) {
            let distSquared = zip(point1, point2)
                .map { ($0 - $1) * ($0 - $1) }
                .reduce(0, +)
            pairs.append(Pair(i: i, j: j, distSquared: distSquared))
        }
    }
    return Heap<Pair>(pairs)
}

private func solvePart1(_ lines: [String]) -> Int {
    let points = lines.map { line in
        line.split(separator: ",").map(Int.mustParse)
    }

    var pairs = makePairsHeap(points)
    var dsu = DSU(count: points.count)
    let maxPairs = (points.count < 100) ? 10 : 1000 // Hack to distinguish example and real input.
    for _ in 0..<maxPairs {
        guard let pair = pairs.popMin() else { break }
        _ = dsu.join(pair.i, pair.j)
    }

    var componentSizes: [Int: Int] = [:]
    for i in points.indices {
        componentSizes[dsu.getRoot(i), default: 0] += 1
    }

    return componentSizes.map(\.value).sorted(by: >).prefix(3).reduce(1, *)
}

private func solvePart2(_ lines: [String]) -> Int {
    let points = lines.map { line in
        line.split(separator: ",").map(Int.mustParse)
    }

    var pairs = makePairsHeap(points)

    var dsu = DSU(count: points.count)
    var components = points.count
    while !pairs.isEmpty {
        let pair = pairs.popMin()!
        let (idx1, idx2) = (pair.i, pair.j)
        if dsu.join(idx1, idx2) {
            components -= 1
            if components == 1 {
                return points[idx1][0] * points[idx2][0]
            }
        }
    }
    assert(false)
    return -1
}

struct Day08: DaySolution {
    let dayNumber = 8

    let part1Solutions = [
        Solution(name: "Default", solve: solvePart1)
    ]

    let part2Solutions = [
        Solution(name: "Default", solve: solvePart2)
    ]
}
