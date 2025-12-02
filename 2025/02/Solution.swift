import Foundation
import AOCUtilities

func pow10(_ num: Int) -> Int {
    precondition(num >= 0)
    var pow = 1
    for _ in 0..<num {
        pow *= 10
    }
    return pow
}

func solvePart1(_ lines: [String]) -> Int {
    let text = lines.joined(separator: "")
    let intervals = text.split(separator: ",")
    var sum = 0
    for interval in intervals {
        let firstLast = interval.split(separator: "-")
        precondition(firstLast.count == 2, "Invalid interval: \(interval)")

        var firstStr = String(firstLast[0])
        var lastStr = String(firstLast[1])
        var first = Int.mustParse(firstStr)
        var last = Int.mustParse(lastStr)
        precondition(lastStr.count - firstStr.count <= 1, "Invalid interval: \(interval)")
        precondition(first <= last, "Invalid interval: \(interval)")

        // Normalize ends to have same digit number
        if firstStr.count != lastStr.count {
            if firstStr.count & 1 != 0 {
                precondition(lastStr.count & 1 == 0)
                first = pow10(lastStr.count - 1)
                firstStr = String(first)
            } else if lastStr.count & 1 != 0 {
                precondition(firstStr.count & 1 == 0)
                last = pow10(firstStr.count) - 1
                lastStr = String(last)
            } else {
                precondition(false, "Either first or last must be an odd number")
            }
        }

        precondition(firstStr.count == lastStr.count)
        if firstStr.count & 1 != 0 {
            continue // no bad ids possible
        }

        let root = pow10(firstStr.count / 2) + 1
        let firstK = (first + root - 1) / root // ceil(first / root)
        let lastK = last / root // floor(first / root)

        sum += root * (firstK + lastK) * (lastK - firstK + 1) / 2
    }

    return sum
}

func isDivisibleBy(_ num: Int, _ roots: [Int]) -> Bool {
    for root in roots where num % root == 0 {
        return true
    }
    return false
}

func digitsNumber(_ num: Int) -> Int {
    var count = 0
    var cur = num
    repeat {
        count += 1
        cur /= 10
    } while cur > 0
    return count
}

func getRoots(_ len: Int) -> [Int] {
    var roots: [Int] = []
    if len <= 1 {
        return roots
    }
    for k in 2...len {
        if len % k != 0 {
            continue
        }
        var root = 1
        for _ in 1..<k {
            root = root * pow10(len / k) + 1
        }
        roots.append(root)
    }
    return roots
}

func solvePart2(_ lines: [String]) -> Int {
    let text = lines.joined(separator: "")
    let intervals = text.split(separator: ",")
    var sum = 0
    let precalcRoots = (0...18).map(getRoots) // Precalculating roots speeds up the solution 10x.
    for interval in intervals {
        let ends = interval.split(separator: "-")
        precondition(ends.count == 2, "Invalid interval \(interval)")
        let first = Int.mustParse(ends[0])
        let last = Int.mustParse(ends[1])
        // TODO: optimize by splitting first-last interval into same-digit-number subintervals
        // and then using precalculated roots to get only candidates.
        // Current approach is good enough though: ~ 1s.
        for k in first...last where isDivisibleBy(k, precalcRoots[digitsNumber(k)]) {
            sum += k
        }
    }
    return sum
}

let part1Solutions = [
    Solution(name: "Default", solve: solvePart1)
]

let part2Solutions = [
    Solution(name: "Default", solve: solvePart2)
]

@main
struct Day02 {
    static func main() {
        var lines: [String] = []
        while let line = readLine() {
            lines.append(line)
        }

        for solution in part1Solutions {
            print("Part 1 (\(solution.name)):", solution.solve(lines))
        }
        for solution in part2Solutions {
            print("Part 2 (\(solution.name)):", solution.solve(lines))
        }
    }
}
