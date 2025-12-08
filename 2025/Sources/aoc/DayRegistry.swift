import Foundation
import AOCUtilities

/// Protocol for days that provide solutions for both parts.
protocol DaySolution: Sendable {
    /// Solutions for Part 1.
    var part1Solutions: [Solution] { get }

    /// Solutions for Part 2.
    var part2Solutions: [Solution] { get }

    /// The day number (e.g., 1, 2, 7).
    var dayNumber: Int { get }
}

enum DayRegistry {
    // Array of all registered day instances
    private static let allDays: [any DaySolution] = [
        Day01(),
        Day02(),
        Day03(),
        Day04(),
        Day05(),
        Day06(),
        Day07(),
        Day08(),
    ]

    // Dictionary for O(1) lookup by day number
    private static let dayMap: [Int: any DaySolution] = {
        Dictionary(uniqueKeysWithValues: allDays.map { ($0.dayNumber, $0) })
    }()

    // Dynamically computed last day
    static var lastDay: Int {
        allDays.map(\.dayNumber).max() ?? 1
    }

    // Available days (sorted)
    static var availableDays: [Int] {
        allDays.map(\.dayNumber).sorted()
    }

    // Get a day solution by day number
    static func get(day: Int) -> (any DaySolution)? {
        dayMap[day]
    }
}
