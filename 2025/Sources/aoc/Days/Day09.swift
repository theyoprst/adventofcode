import Foundation
import AOCUtilities

private func solvePart1(_ lines: [String]) -> Int {
    let points: [(x: Int, y: Int)] = lines.map { line in
        let coords = line.split(separator: ",").map(Int.mustParse)
        precondition(coords.count == 2)
        return (x: coords[0], y: coords[1])
    }

    var maxArea = 0
    for i in points.indices {
        for j in points.indices.dropFirst(i+1) {
            let dx = abs(points[i].x - points[j].x) + 1
            let dy = abs(points[i].y - points[j].y) + 1
            let area = abs(dx * dy)
            maxArea = max(maxArea, area)
        }
    }

    return maxArea
}

private typealias Point = (x: Int, y: Int)

private func compressPoints(_ points: [Point]) -> [Point] {
    let xs = Set(points.map(\.x)).sorted()
    let xCompress = Dictionary(uniqueKeysWithValues: xs.enumerated().map { (i, x) in
        (x, i + 1)
    })

    let ys = Set(points.map(\.y)).sorted()
    let yCompress = Dictionary(uniqueKeysWithValues: ys.enumerated().map { (i, y) in
        (y, i + 1)
    })

    return points.map {
        Point(xCompress[$0.x]!, yCompress[$0.y]!)
    }
}

private func solvePart2(_ lines: [String]) -> Int {
    let origPoints: [Point] = lines.map { line in
        let coords = line.split(separator: ",").map(Int.mustParse)
        precondition(coords.count == 2)
        return Point(x: coords[0], y: coords[1])
    }

    let points = compressPoints(origPoints)

    let maxX = points.reduce(0) { max($1.x, $0) }
    let maxY = points.reduce(0) { max($1.y, $0) }

    var field = Array(repeating: Array(repeating: Character("."), count: maxX + 2), count: maxY + 2)
    for i in 0..<points.count {
        let prev = points[i]
        let cur = points[(i + 1) % points.count]
        field[cur.y][cur.x] = "#"
        let step = Point(
            x: (cur.x - prev.x).signum(),
            y: (cur.y - prev.y).signum(),
        )
        var p = Point(
            x: prev.x + step.x,
            y: prev.y + step.y,
        )
        while p != cur {
            field[p.y][p.x] = "X"
            p.x += step.x
            p.y += step.y
        }
    }

    func fill(_ p: Point, with: Character) {
        var queue = [p]
        var head = 0

        while head < queue.count {
            let p = queue[head]
            head += 1

            if 0 > p.y || p.y >= field.count {
                continue
            }
            if 0 > p.x || p.x >= field[0].count {
                continue
            }
            if field[p.y][p.x] != "." {
                continue
            }
            field[p.y][p.x] = with

            queue.append(Point(x: p.x - 1, y: p.y))
            queue.append(Point(x: p.x + 1, y: p.y))
            queue.append(Point(x: p.x, y: p.y - 1))
            queue.append(Point(x: p.x, y: p.y + 1))
        }
    }

    // Fill outer area with spaces.
    fill(Point(0, 0), with: " ")

    var pref = Array(repeating: Array(repeating: 0, count: maxX + 1), count: maxY + 1)
    for y in 0...maxY {
        for x in 0...maxX {
            let topSum = y > 0 ? pref[y-1][x] : 0
            let leftSum = x > 0 ? pref[y][x-1] : 0
            let diagSum = (y > 0 && x > 0) ? pref[y-1][x-1] : 0
            pref[y][x] = topSum + leftSum - diagSum
            if field[y][x] != " " {
                pref[y][x] += 1
            }
        }
    }

    var maxArea = 0
    for i in points.indices {
        for j in points.indices.dropFirst(i+1) {
            let pointI = points[i]
            let pointJ = points[j]

            let maxX = max(pointI.x, pointJ.x)
            let minX = min(pointI.x, pointJ.x)
            let maxY = max(pointI.y, pointJ.y)
            let minY = min(pointI.y, pointJ.y)

            let rectArea = (maxX - minX + 1) * (maxY - minY + 1)
            let filledAreaTop = minY > 0 ? pref[minY-1][maxX] : 0
            let filledAreaLeft = minX > 0 ? pref[maxY][minX-1] : 0
            let filledAreaDiag = (minX > 0 && minY > 0) ? pref[minY-1][minX-1] : 0
            let filledArea = pref[maxY][maxX] - filledAreaTop - filledAreaLeft + filledAreaDiag

            if rectArea != filledArea {
                continue
            }

            let dx = abs(origPoints[i].x - origPoints[j].x) + 1
            let dy = abs(origPoints[i].y - origPoints[j].y) + 1
            let uncompressedArea = dx * dy
            maxArea = max(maxArea, uncompressedArea)
        }
    }

    return maxArea
}

struct Day09: DaySolution {
    let dayNumber = 9

    let part1Solutions = [
        Solution(name: "Default", solve: solvePart1)
    ]

    let part2Solutions = [
        Solution(name: "Default", solve: solvePart2)
    ]
}
