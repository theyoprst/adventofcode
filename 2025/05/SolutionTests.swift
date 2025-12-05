import Testing
import AOCTestSupport
@testable import Day05

@Suite("Day 05 Solutions")
struct Day05Tests {
    @Test("All test cases from tests.yaml")
    func testFromYAML() throws {
        try runAOCTests(bundle: .module, part1Solutions: part1Solutions, part2Solutions: part2Solutions)
    }
}
