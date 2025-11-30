import Foundation
import Testing
import Yams
@testable import Day00

struct TestInput: Decodable {
    let path: String
    let wantPart1: String?
    let wantPart2: String?
}

struct TestsYAML: Decodable {
    let inputs: [TestInput]
}

@Suite("Day 00 Solutions")
struct Day00Tests {
    @Test("All test cases from tests.yaml")
    func testFromYAML() throws {
        guard let yamlURL = Bundle.module.url(forResource: "tests", withExtension: "yaml") else {
            Issue.record("tests.yaml not found in module bundle")
            return
        }
        let yamlString = try String(contentsOf: yamlURL)
        let tests = try YAMLDecoder().decode(TestsYAML.self, from: yamlString)

        for test in tests.inputs {
            guard let inputURL = Bundle.module.url(forResource: test.path, withExtension: nil) else {
                Issue.record("Input file \(test.path) not found in module bundle")
                continue
            }
            let input = try String(contentsOf: inputURL)
            let lines = input.split(separator: "\n").map(String.init)

            if let expected = test.wantPart1 {
                let result = String(solvePart1(lines))
                #expect(result == expected, "Part 1 failed for \(test.path)")
            }

            if let expected = test.wantPart2 {
                let result = String(solvePart2(lines))
                #expect(result == expected, "Part 2 failed for \(test.path)")
            }
        }
    }
}
