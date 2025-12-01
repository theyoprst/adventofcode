/// Represents a named solution for a part of an Advent of Code puzzle.
/// Multiple solutions can be provided for the same part to test alternative implementations.
public struct Solution<T>: @unchecked Sendable {
    /// Human-readable name for this solution (e.g., "Default", "Optimized", "Linear")
    public let name: String

    /// The solving function that takes input lines and returns a result
    public let solve: @Sendable ([String]) -> T

    public init(name: String, solve: @escaping @Sendable ([String]) -> T) {
        self.name = name
        self.solve = solve
    }
}
