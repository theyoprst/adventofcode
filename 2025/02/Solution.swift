import Foundation
import AOCUtilities

func pow10(_ n: Int) -> Int {
    precondition(n >= 0)
    var p = 1
    for _ in 0..<n {
        p *= 10
    }
    return p
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

func hasRepetitions(_ t: String) -> Bool {
    let bytes = Array(t.utf8)
    if bytes.count == 1 {
        return false // Prevent invalid iteration
    }
    for k in 1...(bytes.count/2) {
        if bytes.count % k != 0 {
            continue
        }
        var matched = true
        for i in 1..<(bytes.count/k) {
            if bytes[0..<k] != bytes[i*k..<(i+1)*k] {
                matched = false
                break
            }
        }
        if matched {
            return true
        }
    }
    return false
}

func solvePart2(_ lines: [String]) -> Int {
    let text = lines.joined(separator: "")
    let intervals = text.split(separator: ",")
    var sum = 0
    for interval in intervals {
        let ends = interval.split(separator: "-")
        precondition(ends.count == 2, "Invalid interval \(interval)")
        let first = Int.mustParse(ends[0])
        let last = Int.mustParse(ends[1])
        for k in first...last {
            if hasRepetitions(String(k)) {
               sum += k
            }
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
