import Foundation
import Testing
import Yams
import AOCUtilities

public struct TestInput: Decodable {
    public let path: String
    public let wantPart1: String?
    public let wantPart2: String?
}

public struct TestsYAML: Decodable {
    public let inputs: [TestInput]
}

/// Runs all test cases from tests.yaml against all provided solutions.
public func runAOCTests<T1: Equatable & CustomStringConvertible, T2: Equatable & CustomStringConvertible>(
    bundle: Bundle,
    part1Solutions: [Solution<T1>],
    part2Solutions: [Solution<T2>]
) throws {
    // Validate at least one solution per part
    precondition(!part1Solutions.isEmpty, "At least one Part 1 solution must be provided")
    precondition(!part2Solutions.isEmpty, "At least one Part 2 solution must be provided")

    guard let yamlURL = bundle.url(forResource: "tests", withExtension: "yaml") else {
        Issue.record("tests.yaml not found in module bundle")
        return
    }

    let yamlString = try String(contentsOf: yamlURL)
    let tests = try YAMLDecoder().decode(TestsYAML.self, from: yamlString)

    for test in tests.inputs {
        guard let inputURL = bundle.url(forResource: test.path, withExtension: nil) else {
            Issue.record("Input file \(String(test.path)) not found in module bundle")
            continue
        }

        let input = try String(contentsOf: inputURL)
        let lines = input.split(separator: "\n").map(String.init)

        // Test Part 1 solutions
        if let expected = test.wantPart1 {
            for solution in part1Solutions {
                let result = solution.solve(lines)
                let resultStr = String(describing: result)
                let message = "Part 1 (\(solution.name)) failed for \(test.path): " +
                              "expected \(expected), got \(resultStr)"
                #expect(resultStr == expected, Comment(rawValue: message))
            }
        }

        // Test Part 2 solutions
        if let expected = test.wantPart2 {
            for solution in part2Solutions {
                let result = solution.solve(lines)
                let resultStr = String(describing: result)
                let message = "Part 2 (\(solution.name)) failed for \(test.path): " +
                              "expected \(expected), got \(resultStr)"
                #expect(resultStr == expected, Comment(rawValue: message))
            }
        }
    }
}

/// Legacy function for backward compatibility with existing tests.
/// Wraps single solvers in Solution structs with "Default" name.
public func runAOCTests<T: Equatable & CustomStringConvertible>(
    bundle: Bundle,
    solvePart1: @escaping @Sendable ([String]) -> T,
    solvePart2: @escaping @Sendable ([String]) -> T
) throws {
    try runAOCTests(
        bundle: bundle,
        part1Solutions: [Solution(name: "Default", solve: solvePart1)],
        part2Solutions: [Solution(name: "Default", solve: solvePart2)]
    )
}
