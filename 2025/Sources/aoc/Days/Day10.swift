import Foundation
import AOCUtilities

private struct Button {
    let indexes: [Int]
    let mask: Int

    init(s: Substring) {
        let s = s.dropFirst().dropLast()
        self.indexes = String(s).extractInts()
        self.mask = Machine.makeNumber(self.indexes)
    }
}

private struct Machine {
    let targetMask: Int
    let buttons: [Button]
    let targetCounters: [Int]

    init(s: String) {
        let parts = s.split(separator: " ")
        self.targetMask = Self.parseTarget(parts.first!)
        let rawButtons = parts[parts.indices.dropFirst().dropLast()]
        self.buttons = Self.parseButtons(rawButtons)
        self.targetCounters = Self.parseTargetPresses(parts.last!)
    }

    static func parseTarget(_ s: Substring) -> Int {
        let s = Self.dropBrackets(s)
        let bits = s.enumerated().compactMap { $1 == "#" ? $0 : nil }
        return Self.makeNumber(bits)
    }

    static func parseButtons(_ arr: ArraySlice<Substring>) -> [Button] {
        return arr.map(Button.init)
    }

    static func parseTargetPresses(_ s: Substring) -> [Int] {
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

private func getMinPressesPart1(_ machine: Machine) -> Int {
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

        // Option 1: Press the button
        recursive(buttonIdx: buttonIdx+1, pressesDone: pressesDone+1, mask: mask ^ button.mask)
        // Option 2: Do not press the button
        recursive(buttonIdx: buttonIdx+1, pressesDone: pressesDone, mask: mask)
    }

    recursive(buttonIdx: 0, pressesDone: 0, mask: machine.targetMask)

    return minPresses
}

private func solvePart1(_ lines: [String]) -> Int {
    let machines = lines.map(Machine.init)
    let minPresses = machines.map{ machine in
        return getMinPressesPart1(machine)
    }
    return minPresses.reduce(0, +)
}

private func solvePart2(_ lines: [String]) -> Int {
    // equation system
    // button  (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
    // press#   x0   x1   x2   x3    x4   x5

    // equasions:
    //          x0  x1  x2  x3  x4  x5  | b
    //          0   0   0   0   1   1   | 3
    //          0   1   0   0   0   1   | 5
    //          0   0   1   1   1   0   | 4
    //          1   1   0   1   0   0   | 7

    // REF (Row Echelon Form), pivot-columns: 0, 1, 2, 4
    //          x0  x1  x2  x3  x4  x5  | b
    //          1   1   0   1   0   0   | 7
    //          0   1   0   0   0   1   | 5
    //          0   0   1   1   1   0   | 4
    //          0   0   0   0   1   1   | 3
    //          *   *   *       *      <- pivots

    // A[0] -= A[1]:
    //          x0  x1  x2  x3  x4  x5  | b
    //          1   0   0   1   0   -1  | 2
    //          0   1   0   0   0    1  | 5
    //          0   0   1   1   1    0  | 4
    //          0   0   0   0   1    1  | 3
    //          *   *   *       *      <- pivots

    // A[0] -= A[1]:
    //          x0  x1  x2  x3  x4  x5  | b
    //          1   0   0   1   0   -1  | 2
    //          0   1   0   0   0    1  | 5
    //          0   0   1   1   1    0  | 4
    //          0   0   0   0   1    1  | 3
    //          *   *   *       *      <- pivots

    // A[2] -= A[3]:
    //          x0  x1  x2  x3  x4  x5  | b
    //          1   0   0   1   0   -1  | 2
    //          0   1   0   0   0    1  | 5
    //          0   0   1   1   0   -1  | 1
    //          0   0   0   0   1    1  | 3
    //          *   *   *       *      <- pivots

    // got RREF (Reduced Ref). Free vars: x3, x5, basis vars: x0, x1, x2, x4.

    // x3, x5 constraints:
    //          x0  x1  x2  x3  x4  x5  | b
    //          1   0   0   1   0   -1  | 2   x0 >= 0  ->  x3 - x5 <= 2   ->  x3 <= x5 + 2
    //          0   1   0   0   0    1  | 5   x1 >= 0  ->  x5 <= 5        ->  x5 <= 5
    //          0   0   1   1   0   -1  | 1   x2 >= 0  ->  x3 - x5 <= 1   ->  x3 <= x5 + 1   ->  x3 = 0...x5+1
    //          0   0   0   0   1    1  | 3   x4 >= 0  ->  x5 <= 3        ->  x5 <= 3        ->  x5 = 0...2
    //          *   *   *       *      <- pivots

    // x5 = 0, x3 = 0, x = [2, 5, 1, 0, 3, 0]  // 11 presses
    // x5 = 0, x3 = 1, x = [1, 5, 0, 1, 3, 0]  // 10 presses

    // x5 = 1, x3 = 0, x = [3, 4, 2, 0, 2, 1]  // 12 presses
    // x5 = 1, x3 = 1, x = [2, 4, 1, 1, 2, 1]  // 11 presses
    // x5 = 1, x3 = 2, x = [1, 4, 0, 2, 2, 1]  // 10 presses

    // x5 = 2, x3 = 0, x = [4, 3, 3, 0, 1, 2]  // 13 presses
    // x5 = 2, x3 = 1, x = [3, 3, 2, 1, 1, 2]  // 12 presses
    // x5 = 2, x3 = 2, x = [2, 3, 1, 2, 1, 2]  // 11 presses
    // x5 = 2, x3 = 3, x = [1, 3, 0, 3, 1, 2]  // 10 presses


    // buttons: (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
    // presses:       x0    x1    x2      x3        x4           b

    // Equasions:
    // x0  x1  x2  x3  x4  |  b
    //  1   0   1   1   0  |  7
    //  0   0   0   1   1  |  5
    //  1   1   0   1   1  | 12
    //  1   1   0   0   1  |  7
    //  1   0   1   0   1  |  2

    // A[2+] -= A[0]:
    // x0  x1  x2  x3  x4  |  b
    //  1   0   1   1   0  |  7
    //  0   0   0   1   1  |  5
    //  0   1  -1   0   1  |  5
    //  0   1  -1  -1   1  |  0
    //  0   0   0  -1   1  | -5

    // Swap A[1] and A[2]
    // x0  x1  x2  x3  x4  |  b
    //  1   0   1   1   0  |  7
    //  0   1  -1   0   1  |  5
    //  0   0   0   1   1  |  5
    //  0   1  -1  -1   1  |  0
    //  0   0   0  -1   1  | -5

    // A[3] -= A[1]
    // x0  x1  x2  x3  x4  |  b
    //  1   0   1   1   0  |  7
    //  0   1  -1   0   1  |  5
    //  0   0   0   1   1  |  5
    //  0   0   0  -1   0  | -5
    //  0   0   0  -1   1  | -5

    // A[3+] += A[2]
    // x0  x1  x2  x3  x4  |  b
    //  1   0   1   1   0  |  7
    //  0   1  -1   0   1  |  5
    //  0   0   0   1   1  |  5
    //  0   0   0   0   1  |  0
    //  0   0   0   0   2  |  0

    // A[4] -= 2 * A[3]
    // x0  x1  x2  x3  x4  |  b
    //  1   0   1   1   0  |  7
    //  0   1  -1   0   1  |  5
    //  0   0   0   1   1  |  5
    //  0   0   0   0   1  |  0
    //  0   0   0   0   0  |  0 <- eliminate

    //    Eliminate A[4]:
    //     x0  x1  x2  x3  x4  |  b
    // A0:  1   0   1   1   0  |  7
    // A1:  0   1  -1   0   1  |  5
    // A2:  0   0   0   1   1  |  5
    // A3:  0   0   0   0   1  |  0
    //      *   *       *   *  <- pivots

    // Got REF, now getting RREF:

    // A0 -= A2
    //     x0  x1  x2  x3  x4  |  b
    // A0:  1   0   1   0  -1  |  2
    // A1:  0   1  -1   0   1  |  5
    // A2:  0   0   0   1   1  |  5
    // A3:  0   0   0   0   1  |  0
    //      *   *       *   *  <- pivots

    // A0 += A3, A1 -= A3, A2 -= A3:
    //     x0  x1  x2  x3  x4  |  b
    // A0:  1   0   1   0   0  |  2
    // A1:  0   1  -1   0   0  |  5
    // A2:  0   0   0   1   0  |  5
    // A3:  0   0   0   0   1  |  0
    //      *   *       *   *  <- pivots


    // Got RREF, now finding bounds for free var x2:

    //     x0  x1  x2  x3  x4  |  b
    // A0:  1   0   1   0   0  |  2       x2 <= 2          ->   x2 = 0...2
    // A1:  0   1  -1   0   0  |  5       -x2 <= 5
    // A2:  0   0   0   1   0  |  5       0 * x2 <= 5
    // A3:  0   0   0   0   1  |  0       0 * x2 <= 2
    //      *   *       *   *  <- pivots

    // x2 = 0, x=[2, 5, 0, 5, 0] <- minimum (12 presses)
    // x2 = 1, x=[1, 6, 1, 5, 0]
    // x2 = 2, x=[0, 7, 2, 5, 0]


    // Example3:
    // buttons: (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
    // presses:         x0      x1          x2    x3

    //      x0  x1  x2  x3  |  b
    // A0:   1   1   1   0  | 10
    // A1:   1   0   1   1  | 11
    // A2:   1   0   1   1  | 11
    // A3:   1   1   0   0  |  5
    // A4:   1   1   1   0  | 10
    // A5:   0   0   1   0  |  5

    // A[1...5] -= A0
    //      x0  x1  x2  x3  |  b
    // A0:   1   1   1   0  | 10
    // A1:   0  -1   0   1  |  1
    // A2:   0  -1   0   1  |  1
    // A3:   0   0  -1   0  | -5
    // A4:   0   0   0   0  |  0
    // A5:   0   0   1   0  |  5

    // A2 -= A1, A1 *= -1
    //      x0  x1  x2  x3  |  b
    // A0:   1   1   1   0  | 10
    // A1:   0   1   0  -1  | -1
    // A2:   0   0   0   0  |  0
    // A3:   0   0  -1   0  | -5
    // A4:   0   0   0   0  |  0
    // A5:   0   0   1   0  |  5

    // swap A2, A3, then A2 (former A3) *= -1
    //      x0  x1  x2  x3  |  b
    // A0:   1   1   1   0  | 10
    // A1:   0   1   0  -1  | -1
    // A3:   0   0   1   0  |  5
    // A2:   0   0   0   0  |  0
    // A4:   0   0   0   0  |  0
    // A5:   0   0   1   0  |  5

    // A5 -= A3:
    //      x0  x1  x2  x3  |  b
    // A0:   1   1   1   0  | 10
    // A1:   0   1   0  -1  | -1
    // A3:   0   0   1   0  |  5
    // A2:   0   0   0   0  |  0
    // A4:   0   0   0   0  |  0
    // A5:   0   0   1   0  |  5

    //      x0  x1  x2  x3  |  b
    // A0:   1   1   1   0  | 10
    // A1:   0   1   0  -1  | -1
    // A3:   0   0   1   0  |  5
    // A2:   0   0   0   0  |  0
    // A4:   0   0   0   0  |  0
    // A5:   0   0   0   0  |  0

    // Found REF, pivot columns: x0, x1, x2, free column: x3.

    // A0 -= A1:
    //      x0  x1  x2  x3  |  b
    // A0:   1   0   1   1  | 11
    // A1:   0   1   0  -1  | -1
    // A3:   0   0   1   0  |  5
    // A2:   0   0   0   0  |  0
    // A4:   0   0   0   0  |  0
    // A5:   0   0   0   0  |  0

    // A0 -= A3:
    //      x0  x1  x2  x3  |  b
    // A0:   1   0   0   1  |  6
    // A1:   0   1   0  -1  | -1
    // A3:   0   0   1   0  |  5
    // A2:   0   0   0   0  |  0
    // A4:   0   0   0   0  |  0
    // A5:   0   0   0   0  |  0

    // x0 = 6 - x3 >= 0  ->  x3 <= 6
    // x1 = x3 - 1 >= 0  ->  x3 >= 1
    // x2 = 5

    // x3 = 6: [0, 5, 5, 6], sum = 16
    // x3 = 1: [5, 0, 5, 1], sum = 11


    // More complex RREF
    //                                            x   y   z
    //    1   0   0   0   0   0   0   0   0   0   0   0   1  |  32  ->  z       <= 32
    //    0   1   0   0   0   0   0   0   0   0   1   0   1  |  26  ->  x + z   <= 26  ->  x <= 26 - z
    //    0   0   1   0   0   0   0   0   0   0   0  -1   0  |   1  ->  -y      <=  1
    //    0   0   0   1   0   0   0   0   0   0   1   0   1  |  34  ->  x + z   <= 34  ->  x <= 34 - z
    //    0   0   0   0   1   0   0   0   0   0   0   0   0  |  15
    //    0   0   0   0   0   1   0   0   0   0   0   1  -1  |  10  ->  y - z   <= 10  ->  z >= 10 + y
    //    0   0   0   0   0   0   1   0   0   0   0   0   0  |  11
    //    0   0   0   0   0   0   0   1   0   0   0   1   0  |  20  ->  y       <= 20
    //    0   0   0   0   0   0   0   0   1   0  -2   0  -1  | -25  ->  -2x - z <= -25  ->  x >= (25 - z) / 2
    //    0   0   0   0   0   0   0   0   0   1   0   0  -1  |  -5  ->  -z      <= -5

    // y = 1..20
    // z = max(5, 10+y)...32
    // x = (25 - z) / 2 ... 26 - z

    //    1   0   1   1   0   0   0   1  |  52
    //    1   1   0   0   1   0   0   0  |  30
    //    0   1   1   0   1   1   1   0  |  67
    //    0   1   1   0   0   0   0   1  |  53
    //    1   0   1   1   0   0   1   1  |  55
    //    0   1   1   1   0   1   0   0  |  67
    //    1   1   1   1   0   0   1   1  |  72
    //    0   0   0   1   0   0   1   0  |  17
    //    1   1   1   0   0   0   0   0  |  36
    //    1   1   0   0   1   1   0   1  |  68

    // A0:    1   0   1   1   0   0   0   1  |  52
    // A1:    0   1  -1  -1   1   0   0  -1  | -22
    // A2:    0   0   2   1   0   1   1   1  |  89
    // A3:    0   0   0   1  -1   0  -1   0  |   0
    // A4:    0   0   0   0   1   1   1  -1  |  14
    // A5:    0   0   0   0   0   1  -1  -1  |  -3
    // A6:    0   0   0   0   0   0   1   0  |   3
    // A7:    0   0   0   0   0   0   0   1  |  19
    // A8:    0   0   0   0   0   0   0   0  |   0
    // A9:    0   0   0   0   0   0   0   0  |   0

    // eliminate(a2)
    // A0:    1   0   1   1   0   0   0   1  |  52
    // A1:    0   1  -1  -1   1   0   0  -1  | -22
    // A2:    0   0   2   1   0   1   1   1  |  89
    // A3:    0   0   0   1  -1   0  -1   0  |   0
    // A4:    0   0   0   0   1   1   1  -1  |  14
    // A5:    0   0   0   0   0   1  -1  -1  |  -3
    // A6:    0   0   0   0   0   0   1   0  |   3
    // A7:    0   0   0   0   0   0   0   1  |  19
    // A8:    0   0   0   0   0   0   0   0  |   0
    // A9:    0   0   0   0   0   0   0   0  |   0


    let machines = lines.map(Machine.init)
    let minPresses = machines.map{ machine in
        return getMinPressesPart2(machine)
    }
    return minPresses.reduce(0, +)
}

private func getMinPressesPart2(_ machine: Machine) -> Int {
    // Create a matrix A and vector b to solve Ax = b
    let cols = machine.buttons.count
    var rows = machine.targetCounters.count
    var a: [[Int]] = Array(repeating: Array(repeating: 0, count: cols), count: rows)
    for (col, button) in machine.buttons.enumerated() {
        for row in button.indexes {
            a[row][col] = 1
        }
    }
    var b: [Int] = machine.targetCounters

    // upperBounds[col] is an upper bound for x_col variable. Lower bound is 0 (start condition).
    var upperBounds: [Int] = Array(repeating: Int.max, count: cols)

    func swapRows(_ first: Int, _ second: Int) {
        a.swapAt(first, second)
        b.swapAt(first, second)
    }

    // normalizeRow:
    // 1. makes first nonzero element positive
    // 2. divides everyone by gcd of nonzero elements (absolutes of them)
    func normalizeRow(_ row: Int) {
        let nonZeroElements = a[row].filter { $0 != 0 } + (b[row] != 0 ? [b[row]] : [])
        if nonZeroElements.count == 0 {
            return
        }
        if nonZeroElements[0] < 0 {
            a[row] = a[row].map { -$0 }
            b[row] *= -1
        }
        let x = gcdAbs(nonZeroElements)
        a[row] = a[row].map { $0 / x }
        b[row] /= x
    }

    func normalizeAllRows() {
        for row in 0..<rows {
            normalizeRow(row)
        }
    }

    func updateUpperBounds(_ row: Int) {
        if a[row].allSatisfy({ $0 >= 0 }) {
            for col in 0..<cols where a[row][col] > 0 {
                upperBounds[col] = min(upperBounds[col], b[row] / a[row][col])
            }
        }
    }

    func updateUpperBounds() {
        for row in 0..<rows {
            updateUpperBounds(row)
        }
    }

    func eliminateRow(row: Int, pivotRow: Int, pivotCol: Int) {
        if a[row][pivotCol] == 0 {
            return
        }
        let signedLcm = a[row][pivotCol] * a[pivotRow][pivotCol] / gcd(abs(a[row][pivotCol]), abs(a[pivotRow][pivotCol]))
        let pivotScalar = signedLcm / a[pivotRow][pivotCol]
        let curScalar = signedLcm / a[row][pivotCol]
        assert(signedLcm % a[row][pivotCol] == 0)
        assert(signedLcm % a[pivotRow][pivotCol] == 0)

        for col in 0..<cols {
            a[row][col] = a[row][col] * curScalar - a[pivotRow][col] * pivotScalar
        }
        assert(a[row][pivotCol] == 0)

        b[row] = b[row] * curScalar - b[pivotRow] * pivotScalar

        normalizeRow(row)
        updateUpperBounds(row)
    }

    normalizeAllRows()
    updateUpperBounds()

    // Build REF
    var (pivotRow, pivotCol) = (0, 0)
    var isPivotCol: [Bool] = Array(repeating: false, count: cols)
    var pivotCols: Array<Int> = []
    while pivotRow < rows && pivotCol < cols {
        var nonZeroRow: Int? = nil
        for row in pivotRow..<rows {
            if a[row][pivotCol] != 0 {
                nonZeroRow = row
                break
            }
        }
        guard let nonZeroRowIndex = nonZeroRow else {
            pivotCol += 1
            continue
        }
        isPivotCol[pivotCol] = true
        pivotCols.append(pivotCol)
        swapRows(nonZeroRowIndex, pivotRow)
        assert(a[pivotRow][pivotCol] != 0)

        // Eliminate non-zero values under the pivot element:
        for row in pivotRow+1..<rows {
            eliminateRow(row: row, pivotRow: pivotRow, pivotCol: pivotCol)
        }
        pivotCol += 1
        pivotRow += 1
    }

    // Build RREF: zero elements above the pivots
    for (pivotRow, pivotCol) in pivotCols.enumerated() {
        for row in 0..<pivotRow {
            eliminateRow(row: row, pivotRow: pivotRow, pivotCol: pivotCol)
        }
    }

    // Remove zero lines
    rows = pivotCols.count
    a.removeLast(a.count - rows)
    b.removeLast(a.count - rows)

    // Move non-pivot columns to the right
    for row in 0..<rows {
        var newRow : [Int] = []
        for (col, isPivot) in isPivotCol.enumerated() where isPivot{
            newRow.append(a[row][col])
        }
        for (col, isPivot) in isPivotCol.enumerated() where !isPivot {
            newRow.append(a[row][col])
        }
        a[row] = newRow
    }

    var upperBounds2: [Int] = []
    for (col, isPivot) in isPivotCol.enumerated() where isPivot {
        upperBounds2.append(upperBounds[col])
    }
    for (col, isPivot) in isPivotCol.enumerated() where !isPivot {
        upperBounds2.append(upperBounds[col])
    }
    upperBounds = upperBounds2


    var solution: [Int] = Array(repeating: 0, count: cols)

    func getSolution() -> [Int]? {
        for row in 0..<rows {
            var y = b[row]
            for freeVar in rows..<cols {
                y -= solution[freeVar] * a[row][freeVar]
            }
            if y % a[row][row] != 0 {
                return nil
            }
            solution[row] = y / a[row][row]
            if solution[row] < 0 {
                return nil
            }
        }
        return solution
    }

    // Now real fun starts. Iterate over possible values of free variables.
    var minSum = 1000000000
    func iter(_ col: Int) {
        if col == cols {
            guard let sol = getSolution() else {
                return
            }
            let sum = sol.reduce(0, +)
            minSum = min(minSum, sum)
            return
        }

        for value in 0...upperBounds[col] {
            solution[col] = value
            iter(col+1)
        }
        solution[col] = 0
    }

    iter(rows)

    return minSum
}

func gcd(_ x: Int, _ y: Int) -> Int {
    var (x, y) = (x, y)
    while x != 0 {
        (x, y) = (y % x, x)
    }
    return y
}

func gcdAbs(_ numbers: [Int]) -> Int {
    assert(numbers.count > 0)
    var result = abs(numbers[0])
    for number in numbers.dropFirst() {
        result = gcd(result, abs(number))
    }
    return result
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
