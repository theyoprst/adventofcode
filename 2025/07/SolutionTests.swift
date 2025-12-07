import Testing
import AOCTestSupport
@testable import Day07

@Suite("Day 07 Solutions")
struct Day07Tests {
    @Test("All test cases from tests.yaml")
    func testFromYAML() throws {
        try runAOCTests(bundle: .module, part1Solutions: part1Solutions, part2Solutions: part2Solutions)
    }
}
