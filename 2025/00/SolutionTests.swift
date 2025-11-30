import Testing
import AOCTestSupport
@testable import Day00

@Suite("Day 00 Solutions")
struct Day00Tests {
    @Test("All test cases from tests.yaml")
    func testFromYAML() throws {
        try runAOCTests(bundle: .module, solvePart1: solvePart1, solvePart2: solvePart2)
    }
}
