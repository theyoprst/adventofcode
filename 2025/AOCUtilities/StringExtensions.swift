import Foundation

public extension String {
    /// Extracts all integers from the string, regardless of delimiters.
    /// Supports negative numbers (with `-`) and optional `+` sign.
    /// This mirrors Go's `aoc.Ints` function.
    ///
    /// Examples:
    /// - "1 2 3" → [1, 2, 3]
    /// - "1,2,3" → [1, 2, 3]
    /// - "-5 +10 20" → [-5, 10, 20]
    /// - "x=1, y=-2" → [1, -2]
    func extractInts() -> [Int] {
        let regex = /[-+]?\d+/
        return self.matches(of: regex).map { match in
            Int.mustParse(String(match.output))
        }
    }
}
