import Foundation

public extension Int {
    /// Parses an integer from a string, crashing with precondition if parsing fails.
    /// This is intended for Advent of Code where input is assumed to be well-formed.
    static func mustParse<S: StringProtocol>(_ string: S) -> Int {
        guard let value = Int(string) else {
            preconditionFailure("Failed to parse '\(string)' as Int")
        }
        return value
    }
}
