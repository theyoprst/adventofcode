import Foundation
import Testing
import Yams

public struct TestInput: Decodable {
    public let path: String
    public let wantPart1: String?
    public let wantPart2: String?
}

public struct TestsYAML: Decodable {
    public let inputs: [TestInput]
}

public func runAOCTests<T>(
    bundle: Bundle,
    solvePart1: ([String]) -> T,
    solvePart2: ([String]) -> T
) throws {
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

        if let expected = test.wantPart1 {
            let result = String(describing: solvePart1(lines))
            #expect(result == expected, "Part 1 failed for \(String(test.path))")
        }

        if let expected = test.wantPart2 {
            let result = String(describing: solvePart2(lines))
            #expect(result == expected, "Part 2 failed for \(String(test.path))")
        }
    }
}
