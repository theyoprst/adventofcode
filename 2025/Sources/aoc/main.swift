import ArgumentParser
import Foundation
import AOCUtilities

// Helper to find the resource bundle for executables
private func findResourceBundle() -> Bundle {
    // For executables, SPM creates a separate bundle for resources
    // The bundle name format is: <PackageName>_<TargetName>.bundle
    let bundleName = "AdventOfCode2025_aoc"

    guard let executableURL = Bundle.main.executableURL else {
        fatalError("Could not determine executable location")
    }

    let bundleURL = executableURL.deletingLastPathComponent().appendingPathComponent("\(bundleName).bundle")

    guard let bundle = Bundle(url: bundleURL) else {
        fatalError("Could not find resource bundle at \(bundleURL.path)")
    }

    return bundle
}

struct AOC: ParsableCommand {
    static let configuration = CommandConfiguration(
        commandName: "aoc",
        abstract: "Advent of Code 2025 Solutions",
        discussion: """
            Run Advent of Code 2025 solutions for a specific day.
            By default, runs the latest available day.
            """,
        version: "1.0.0"
    )

    @Option(
        name: .long,
        help: ArgumentHelp(
            "Day to run (1-12) or 'last' for the latest day.",
            discussion: "Defaults to the last available day if not specified."
        )
    )
    var day: String?

    func run() throws {
        let selectedDay = try parseDayValue(day)

        print("Running Day \(String(format: "%02d", selectedDay))")
        print()

        guard let daySolution = DayRegistry.get(day: selectedDay) else {
            let availableDays = DayRegistry.availableDays.map(String.init).joined(separator: ", ")
            throw ValidationError("Day \(selectedDay) not found. Available days: \(availableDays)")
        }

        let bundle = findResourceBundle()
        runAllInputs(
            part1Solutions: daySolution.part1Solutions,
            part2Solutions: daySolution.part2Solutions,
            bundle: bundle,
            daySubdirectory: String(format: "%02d", daySolution.dayNumber)
        )
    }

    private func parseDayValue(_ value: String?) throws -> Int {
        guard let value else {
            return DayRegistry.lastDay
        }

        if value.lowercased() == "last" {
            return DayRegistry.lastDay
        }

        guard let dayNumber = Int(value) else {
            throw ValidationError("Day must be a number (1-12) or 'last', got: '\(value)'")
        }

        guard (1...12).contains(dayNumber) else {
            throw ValidationError("Day must be between 1 and 12, got: \(dayNumber)")
        }

        return dayNumber
    }
}

AOC.main()
