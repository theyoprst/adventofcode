import Foundation
import AOCUtilities

private func solvePart1(_ lines: [String]) -> Int {
    var reverseGraph: [String: [String]] = [:]
    for line in lines {
        let srcDst = line.split(separator: ": ", maxSplits: 2)
        let src = srcDst.first!
        for dst in srcDst.last!.split(separator: " ") {
            reverseGraph[String(dst), default: []].append(String(src))
        }
    }

    var routes: [String: Int] = ["you": 1]
    func findRoutesDfs(_ from: String) -> Int {
        if let count = routes[from] {
            return count
        }
        let sources = reverseGraph[from] ?? []
        let total = sources.reduce(0) { $0 + findRoutesDfs($1) }
        routes[from] = total
        return total
    }

    return findRoutesDfs("out")
}

private func solvePart2(_ lines: [String]) -> Int {
    struct Vertex: Hashable {
        let name: String
        let fft: Bool
        let dac: Bool
    }

    var reverseGraph: [Vertex: [Vertex]] = [:]
    for line in lines {
        let srcDst = line.split(separator: ": ", maxSplits: 2)
        let src = String(srcDst.first!)
        for dst in srcDst.last!.split(separator: " ") {
            let dst = String(dst)
            let setsFft = (dst == "fft")
            let setsDac = (dst == "dac")

            for fft in [true, false] {
                for dac in [true, false] {
                    let vSrc = Vertex(name: src, fft: fft, dac: dac)
                    let vDst = Vertex(name: dst, fft: fft || setsFft, dac: dac || setsDac )
                    reverseGraph[vDst, default: []].append(vSrc)
                }
            }
        }
    }

    let start = Vertex(name: "svr", fft: false, dac: false)
    var routes: [Vertex: Int] = [start: 1]
    func findRoutesDfs(_ from: Vertex) -> Int {
        if let count = routes[from] {
            return count
        }
        let sources = reverseGraph[from] ?? []
        let total = sources.reduce(0) { $0 + findRoutesDfs($1) }
        routes[from] = total
        return total
    }

    let target = Vertex(name: "out", fft: true, dac: true)
    return findRoutesDfs(target)
}

struct Day11: DaySolution {
    let dayNumber = 11

    let part1Solutions = [
        Solution(name: "Default", solve: solvePart1)
    ]

    let part2Solutions = [
        Solution(name: "Default", solve: solvePart2)
    ]
}
