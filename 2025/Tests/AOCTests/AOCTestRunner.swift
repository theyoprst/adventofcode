import Foundation
import Testing
import Yams
import AOCUtilities
@testable import aoc

struct TestInput: Decodable {
    let path: String
    let wantPart1: String?
    let wantPart2: String?
}

struct TestsYAML: Decodable {
    let inputs: [TestInput]
}

/// Runs all test cases from tests.yaml against the provided day solution.
func runAOCTests(bundle: Bundle, day: any DaySolution) throws {
    // Validate at least one solution per part
    precondition(!day.part1Solutions.isEmpty, "At least one Part 1 solution must be provided")
    precondition(!day.part2Solutions.isEmpty, "At least one Part 2 solution must be provided")

    let daySubdirectory = String(format: "%02d", day.dayNumber)
    guard let yamlURL = bundle.url(forResource: "\(daySubdirectory)/tests", withExtension: "yaml") else {
        Issue.record("tests.yaml not found in module bundle")
        return
    }

    let yamlString = try String(contentsOf: yamlURL)
    let tests = try YAMLDecoder().decode(TestsYAML.self, from: yamlString)

    for test in tests.inputs {
        let resourcePath = "\(daySubdirectory)/\(test.path)"
        guard let inputURL = bundle.url(forResource: resourcePath, withExtension: nil) else {
            Issue.record("Input file \(resourcePath) not found in module bundle")
            continue
        }

        let input = try String(contentsOf: inputURL)
        var lines = input.split(separator: "\n", omittingEmptySubsequences: false).map(String.init)
        while lines.last?.isEmpty == true {
            lines.removeLast()
        }

        // TODO: Subtests for parts and for solutions.

        // Test Part 1 solutions
        if let expected = test.wantPart1 {
            for solution in day.part1Solutions {
                let result = solution.solve(lines)
                let resultStr = String(describing: result)
                let message = "Part 1 (\(solution.name)) failed for \(test.path): " +
                              "expected \(expected), got \(resultStr)"
                #expect(resultStr == expected, Comment(rawValue: message))
            }
        }

        // Test Part 2 solutions
        if let expected = test.wantPart2 {
            for solution in day.part2Solutions {
                let result = solution.solve(lines)
                let resultStr = String(describing: result)
                let message = "Part 2 (\(solution.name)) failed for \(test.path): " +
                              "expected \(expected), got \(resultStr)"
                #expect(resultStr == expected, Comment(rawValue: message))
            }
        }
    }
}
