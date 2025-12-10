import Foundation
import AOCUtilities

private struct Button {
    let bits: [Int]

    init(s: String) {
        let s = s.dropFirst().dropLast()
        self.bits = s.enumerated().compactMap { $1 == "#" ? $0 : nil }
    }
}

private struct Machine {
    let target: Int
    let buttons: [Int]
    let joltage: [Int]

    init(s: String) {
        let parts = s.split(separator: " ")
        self.target = Self.parseTarget(parts.first!)
        let rawButtons = parts[parts.indices.dropFirst().dropLast()]
        self.buttons = Self.parseButtons(rawButtons)
        self.joltage = Self.parseJoltage(parts.last!)
    }

    static func parseTarget(_ s: Substring) -> Int {
        let s = Self.dropBrackets(s)
        let bits = s.enumerated().compactMap { $1 == "#" ? $0 : nil }
        return Self.makeNumber(bits)
    }

    static func parseButtons(_ arr: ArraySlice<Substring>) -> [Int] {
        return arr.map { s in
            let s = Self.dropBrackets(s)
            let bits = String(s).extractInts()
            return Self.makeNumber(bits)
        }
    }

    static func parseJoltage(_ s: Substring) -> [Int] {
        let s = Self.dropBrackets(s)
        return String(s).extractInts() // TODO: write Substring extension too?
    }

    static func dropBrackets(_ s: Substring) -> Substring {
        return s.dropFirst().dropLast()
    }

    static func makeNumber(_ bits: [Int]) -> Int {
        var number = 0
        for bit in bits {
            number += 1 << bit
        }
        return number
    }
}

private func getMinPresses(_ machine: Machine) -> Int {
    var minPresses = Int.max

    func recursive(buttonIdx: Int, pressesDone: Int, mask: Int) {
        if mask == 0 {
            minPresses = min(minPresses, pressesDone)
            return
        }
        guard buttonIdx < machine.buttons.count else {
            return
        }
        let button = machine.buttons[buttonIdx]

        // Opt 1: Press the button
        recursive(buttonIdx: buttonIdx+1, pressesDone: pressesDone+1, mask: mask ^ button)
        // Opt 2: Do not press the button
        recursive(buttonIdx: buttonIdx+1, pressesDone: pressesDone, mask: mask)
    }

    recursive(buttonIdx: 0, pressesDone: 0, mask: machine.target)

    return minPresses
}

private func solvePart1(_ lines: [String]) -> Int {
    let machines = lines.map(Machine.init)
    let minPresses = machines.map{ machine in
        return getMinPresses(machine)
    }
    return minPresses.reduce(0, +)
}

private func solvePart2(_ lines: [String]) -> Int {
    // equation system
    // button  (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
    // press#   a0   a1   a2   a3    a4   a5
    // equasions:
    //          a0  a1  a2  a3  a4  a5 |
    //          0   0   0   0   1   1  | 3
    //          0   1   0   0   0   1  | 5
    //          0   0   1   1   1   0  | 4
    //          1   1   0   1   0   0  | 7

    //          a0  a1  a2  a3  a4  a5 |
    //          1   1   0   1   0   0  | 7
    //          0   1   0   0   0   1  | 5
    //          0   0   1   1   1   0  | 4
    //          0   0   0   0   1   1  | 3

//    let machines = lines.map(Machine.init)
//    for machine in machines {
//        print("equations count: \(machine.joltage.count), variables count: \(machine.buttons.count), diff is: \(machine.joltage.count - machine.buttons.count), total buttons presses: \(machine.joltage.reduce(0, +))")
//    }
    return 0
}

struct Day10: DaySolution {
    let dayNumber = 10

    let part1Solutions = [
        Solution(name: "Default", solve: solvePart1)
    ]

    let part2Solutions = [
        Solution(name: "Default", solve: solvePart2)
    ]
}
