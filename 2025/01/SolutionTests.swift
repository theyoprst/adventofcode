import Testing
import AOCTestSupport
@testable import Day01

@Suite("Day 01 Solutions")
struct Day01Tests {
    @Test("All test cases from tests.yaml")
    func testFromYAML() throws {
        try runAOCTests(bundle: .module, solvePart1: solvePart1, solvePart2: solvePart2)
    }
}
