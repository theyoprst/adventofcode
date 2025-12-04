import Testing
import AOCTestSupport
@testable import Day04

@Suite("Day 04 Solutions")
struct Day04Tests {
    @Test("All test cases from tests.yaml")
    func testFromYAML() throws {
        try runAOCTests(bundle: .module, part1Solutions: part1Solutions, part2Solutions: part2Solutions)
    }
}
